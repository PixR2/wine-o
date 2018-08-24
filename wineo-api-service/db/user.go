package db

import "errors"

type User struct {
	Email string `json:"email" gorethink:"id"`

	FirstName *string `json:"firstName" gorethink:"firstName"`
	LastName  *string `json:"lastName" gorethink:"lastName"`

	EmailConfirmed bool `json:"emailConfirmed" gorethink:"emailConfirmed"`
	HasPassword    bool `json:"hasPassword" gorethink:"hasPassword"`

	PasswordHash []byte `json:"-" gorethink:"passwordHash"`
}

func (s *Session) ConfirmEmail(email string) error {
	return errors.New("not implemented")
}

func (s *Session) CreateUser(email string) error {
	return errors.New("not implemented")
}

func (s *Session) UpdateUser(email string, firstName *string, lastName *string, password *string) error {
	return errors.New("not implemented")
}

func (s *Session) GetUser(email string) (*User, error) {
	return nil, errors.New("not implemented")
}
