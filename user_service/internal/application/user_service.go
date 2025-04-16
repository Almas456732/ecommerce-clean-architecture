package application

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"
	"user_service/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

// UserService defines the interface for user operations
type UserService interface {
	Register(username, password, email, fullName string) (*domain.User, error)
	GetByID(id string) (*domain.User, error)
	GetByUsername(username string) (*domain.User, error)
	Authenticate(username, password string) (*domain.User, error)
	UpdateUser(user *domain.User, email, fullName, password string) (*domain.User, error)
	DeleteUser(id string) error
	ListUsers() ([]*domain.User, error)
}

// userService implements the UserService interface
type userService struct {
	userRepo domain.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo domain.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Register creates a new user
func (s *userService) Register(username, password, email, fullName string) (*domain.User, error) {
	// Check if user with the same username already exists
	existingUser, err := s.userRepo.FindByUsername(username)
	if err == nil && existingUser != nil {
		return nil, errors.New("username already taken")
	}

	// Check if user with the same email already exists
	existingUser, err = s.userRepo.FindByEmail(email)
	if err == nil && existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Generate a unique ID
	id, err := generateID()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	user := &domain.User{
		ID:        id,
		Username:  username,
		Password:  string(hashedPassword),
		Email:     email,
		FullName:  fullName,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = s.userRepo.Save(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetByID retrieves a user by ID
func (s *userService) GetByID(id string) (*domain.User, error) {
	return s.userRepo.FindByID(id)
}

// GetByUsername retrieves a user by username
func (s *userService) GetByUsername(username string) (*domain.User, error) {
	return s.userRepo.FindByUsername(username)
}

// Authenticate verifies user credentials and returns the user if valid
func (s *userService) Authenticate(username, password string) (*domain.User, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	return user, nil
}

// UpdateUser updates user information
func (s *userService) UpdateUser(user *domain.User, email, fullName, password string) (*domain.User, error) {
	if email != "" && email != user.Email {
		// Check if the new email is already taken
		existingUser, err := s.userRepo.FindByEmail(email)
		if err == nil && existingUser != nil && existingUser.ID != user.ID {
			return nil, errors.New("email already registered")
		}
		user.Email = email
	}

	if fullName != "" {
		user.FullName = fullName
	}

	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}

	user.UpdatedAt = time.Now()

	err := s.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser removes a user by ID
func (s *userService) DeleteUser(id string) error {
	return s.userRepo.Delete(id)
}

// ListUsers returns all users
func (s *userService) ListUsers() ([]*domain.User, error) {
	return s.userRepo.List()
}

// Helper function to generate a unique ID
func generateID() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
