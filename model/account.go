package model

import (
	"github.com/pkg/errors"

	r "gopkg.in/gorethink/gorethink.v4"
)

type Account struct {
	Email        string `gorethink:"id"`
	PasswordHash []byte `gorethink:"pwd_hash" json:"-"`
	FirstName    string `gorethink:"first_name"`
	LastName     string `gorethink:"last_name"`
	Verified     bool   `gorethink:"verified"`
}

func CreateAccount(acc *Account, session *r.Session) error {
	_, err := r.DB("wineo").Table("accounts").Insert(acc).RunWrite(session)
	if err != nil {
		return errors.Wrap(err, "failed to insert new account into rethinkdb")
	}

	return nil
}

func UpdateAccount(acc *Account) error {
	return errors.New("not implemented")
}

func DeleteAccount(email string) error {
	return errors.New("not implemented")
}

func GetAccount(email string) error {
	return errors.New("not implemented")
}
