package model

import "errors"

type WinePosition struct {
	CollectionID string `gorethink:"id[0]"`
	Compartment  string `gorethink:"id[1]"`

	Wine Wine `gorethink:"id[2],reference" gorethink_ref:"id"`

	BuyPrice int
	Bottles  int
}

func CreateWinePosition(pos *WinePosition) error {
	return errors.New("not implemented")
}

func UpdateWinePosition(pos *WinePosition) error {
	return errors.New("not implemented")
}

func DeleteWinePosition(email string) error {
	return errors.New("not implemented")
}

func GetWinePosition(email string) error {
	return errors.New("not implemented")
}
