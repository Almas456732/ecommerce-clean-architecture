package domain

// UserRepository defines the contract for user data access
type UserRepository interface {
	Save(user *User) error
	FindByID(id string) (*User, error)
	FindByUsername(username string) (*User, error)
	FindByEmail(email string) (*User, error)
	Update(user *User) error
	Delete(id string) error
	List() ([]*User, error)
}
