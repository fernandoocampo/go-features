package newerrors

import "errors"

var (
	ErrInvalidEmail     = errors.New("invalid email")
	ErrInvalidLastName  = errors.New("invalid last name")
	ErrInvalidUserName  = errors.New("invalid user name")
	ErrInvalidFirstName = errors.New("invalid first name")
)

type User struct {
	FirstName string
	LastName  string
	Username  string
	Email     string
}

func (u User) Validate() error {
	var invalidUser error

	if u.FirstName == "" {
		invalidUser = errors.Join(invalidUser, ErrInvalidFirstName)
	}

	if u.LastName == "" {
		invalidUser = errors.Join(invalidUser, ErrInvalidLastName)
	}

	if u.Username == "" {
		invalidUser = errors.Join(invalidUser, ErrInvalidUserName)
	}

	if u.Email == "" {
		invalidUser = errors.Join(invalidUser, ErrInvalidEmail)
	}

	if invalidUser != nil {
		return invalidUser
	}
	return nil
}
