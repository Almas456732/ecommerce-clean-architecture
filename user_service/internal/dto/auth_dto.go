package dto

// LoginRequest represents credentials needed for login
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse represents the successful login response with JWT token
type LoginResponse struct {
	Token    string       `json:"token"`
	User     UserResponse `json:"user"`
	ExpireAt int64        `json:"expire_at"`
}
