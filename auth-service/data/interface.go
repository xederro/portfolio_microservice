package data

import (
	"github.com/nedpals/supabase-go"
)

type Model struct {
	Auth        supabase.AuthenticatedDetails `json:"auth,omitempty"`
	Credentials supabase.UserCredentials      `json:"credentials,omitempty"`
	UpdateData  map[string]interface{}        `json:"updateData,omitempty"`
}

type IModel interface {
	SignIn() (string, error)
	SignUp() error
	SignOut(token string) error
	// GetAll returns a slice of all users, sorted by last name
	GetAll() ([]*supabase.User, error)

	// GetByEmail returns one user by email
	GetByEmail(email string) (*supabase.User, error)

	// GetOne returns one user by token
	GetOne(token string) (*supabase.User, error)

	// Update updates one user in the database, using the information
	// stored in the receiver u
	Update(token string, updateData map[string]interface{}) error

	// Delete deletes one user from the database, by Model.ID
	Delete() error

	// DeleteByID deletes one user from the database, by ID
	DeleteByID(id string) error

	// ResetPassword is the method we will use to reset password.
	ResetPassword(email string) error
}
