package db

import (
	"errors"

	rt "gopkg.in/gorethink/gorethink.v4"
)

type Session struct {
	*rt.Session
}

func NewSession(opts *rt.ConnectOpts) (*Session, error) {
	return nil, errors.New("not implemented")
}
