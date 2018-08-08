package model

import "errors"

type WineCollection struct {
	ID     string    `gorethink:"id,omitempty"`
	Name   string    `gorethink:"name"`
	Owners []Account `gorethink:"owner_ids,reference" gorethink_ref:"id"`
}

func CreateWineCollection(collection *WineCollection) (string, error) {
	return "", errors.New("not implemented")
}

func UpdateWineCollection(collection *WineCollection) error {
	return errors.New("not implemented")
}

func DeleteWineCollection(id string) error {
	return errors.New("not implemented")
}

func GetWineCollection(id string) error {
	return errors.New("not implemented")
}
