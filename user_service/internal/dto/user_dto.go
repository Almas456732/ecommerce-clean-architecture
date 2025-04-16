package dto

// UserRegistrationRequest represents the data needed to register a new user
type UserRegistrationRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	FullName string `json:"full_name" binding:"required"`
}

// UserResponse is the standard user representation for API responses
type UserResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FullName  string `json:"full_name"`
	CreatedAt string `json:"created_at"`
}

// UserProfileResponse contains detailed user information
type UserProfileResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FullName  string `json:"full_name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// UserUpdateRequest represents the data needed to update a user
type UserUpdateRequest struct {
	Email    string `json:"email" binding:"omitempty,email"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
}
