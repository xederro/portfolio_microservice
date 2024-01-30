package data

import (
	"firebase.google.com/go/v4/auth"
)

type UserCredentials struct {
	Email    string
	Password string
}

type Model struct {
	Auth        auth.UserInfo   `json:"auth,omitempty"`
	Credentials UserCredentials `json:"credentials,omitempty"`
}

type IModel interface {
	SignUp() error

	// Update updates one user in the database, using the information
	// stored in the receiver u
	Update(token string) error

	// Delete deletes one user from the database, by Model.ID
	Delete(token string) error

	GetUUID(token string) (string, error)
}
