package db

import "errors"

type WineCollection struct {
	ID   string `json:"id" gorethink:"id"`
	Name string `json:"name" gorethink:"name"`

	Owner         string   `json:"owner" gorethink:"owner"`
	Collaborators []string `json:"collaborators" gorethink:"collaborators"`
}

func (s *Session) CreateCollection(name string) (string, error) {
	return "", errors.New("not implemented")
}

func (s *Session) GetCollectionsForUser(email string) ([]WineCollection, error) {
	return nil, errors.New("not implemented")
}

func (s *Session) UpdateCollection(id string, name *string) error {
	return errors.New("not implemented")
}
