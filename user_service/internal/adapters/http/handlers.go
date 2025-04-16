package http

import (
	"net/http"
	"time"

	"user_service/internal/adapters/auth"
	"user_service/internal/application"
	"user_service/internal/domain"
	"user_service/internal/dto"

	"github.com/gin-gonic/gin"
)

// UserHandler handles HTTP requests for user operations
type UserHandler struct {
	userService  application.UserService
	tokenService *auth.TokenService
}

// NewUserHandler creates a new instance of UserHandler
func NewUserHandler(userService application.UserService, tokenService *auth.TokenService) *UserHandler {
	return &UserHandler{
		userService:  userService,
		tokenService: tokenService,
	}
}

// SetupRoutes configures the routes for the user service
func SetupRoutes(router *gin.Engine, handler *UserHandler) {
	// Root health check
	router.GET("/health", handler.HealthCheck)

	// User routes group
	users := router.Group("/users")
	{
		// Public routes
		users.POST("/register", handler.Register)
		users.POST("/login", handler.Login)
		users.GET("/health", handler.HealthCheck)

		// These routes will be protected by API gateway's auth middleware
		users.GET("/profile", handler.GetProfile)
		users.PUT("/profile", handler.UpdateProfile)
		users.DELETE("/profile", handler.DeleteProfile)

		// Admin routes (protected by API gateway)
		users.GET("", handler.ListUsers)
		users.GET("/:id", handler.GetUser)
	}
}

// Register handles user registration
func (h *UserHandler) Register(c *gin.Context) {
	var request dto.UserRegistrationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.Register(
		request.Username,
		request.Password,
		request.Email,
		request.FullName,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, toUserResponse(user))
}

// Login handles user authentication and returns a JWT token
func (h *UserHandler) Login(c *gin.Context) {
	var request dto.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.Authenticate(request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, expireAt, err := h.tokenService.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{
		Token:    token,
		User:     toUserResponse(user),
		ExpireAt: expireAt,
	})
}

// GetProfile returns the authenticated user's profile
func (h *UserHandler) GetProfile(c *gin.Context) {
	// The API gateway will validate the token and pass along the userID
	userID := c.GetHeader("X-User-ID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not provided"})
		return
	}

	user, err := h.userService.GetByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, toUserProfileResponse(user))
}

// UpdateProfile updates the authenticated user's profile
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	var request dto.UserUpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetHeader("X-User-ID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not provided"})
		return
	}

	user, err := h.userService.GetByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	updatedUser, err := h.userService.UpdateUser(user, request.Email, request.FullName, request.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, toUserProfileResponse(updatedUser))
}

// DeleteProfile deletes the authenticated user's profile
func (h *UserHandler) DeleteProfile(c *gin.Context) {
	userID := c.GetHeader("X-User-ID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not provided"})
		return
	}

	err := h.userService.DeleteUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// GetUser gets a specific user by ID (admin function)
func (h *UserHandler) GetUser(c *gin.Context) {
	userID := c.Param("id")
	user, err := h.userService.GetByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, toUserResponse(user))
}

// ListUsers returns a list of all users (admin function)
func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.userService.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	userResponses := make([]dto.UserResponse, 0, len(users))
	for _, user := range users {
		userResponses = append(userResponses, toUserResponse(user))
	}

	c.JSON(http.StatusOK, userResponses)
}

// HealthCheck returns a 200 OK response if the service is running
func (h *UserHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"time":   time.Now().Format(time.RFC3339),
	})
}

// Helper function to convert a domain.User to dto.UserResponse
func toUserResponse(user *domain.User) dto.UserResponse {
	return dto.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		FullName:  user.FullName,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}
}

// Helper function to convert a domain.User to dto.UserProfileResponse
func toUserProfileResponse(user *domain.User) dto.UserProfileResponse {
	return dto.UserProfileResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		FullName:  user.FullName,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}
}
