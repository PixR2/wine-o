package db

import "errors"

type Position struct {
	ID           *string `json:"id" gorethink:"id,omitempty"`
	CollectionID *string `json:"collectionID" gorethink:"collectionID,omitempty"`

	Winery     *string `json:"winery" gorethink:"wineID[0],omitempty"`
	Varietal   *string `json:"varietal" gorethink:"wineID[1],omitempty"`
	Vintage    *int    `json:"vintage" gorethink:"wineID[2],omitempty"`
	BottleSize *int    `json:"bottleSize" gorethink:"wineID[3],omitempty"`
	Region     *string `json:"region" gorethink:"region,omitempty"`
	Peak       *int    `json:"peak" gorethink:"peak,omitempty"`
	Hold       *int    `json:"hold" gorethink:"hold,omitempty"`
	Drink      *int    `json:"drink" gorethink:"drink,omitempty"`
	Taste      *int    `json:"taste" gorethink:"taste,omitempty"`
	Price      *int    `json:"price" gorethink:"price,omitempty"`
}

func (s *Session) CreatePosition(winery string, varietal string, vintage int, bottleSize int, numBottles int) (string, error) {
	return "", errors.New("not implemented")
}

func (s *Session) UpdatePosition(positionID string, winery string, varietal string, vintage int, bottleSize int, numBottles int) (string, error) {
	return "", errors.New("not implemented")
}

func (s *Session) SubscribeToPositionsOfCollection(collectionID string) (<-chan Position, error) {
	return nil, errors.New("not implemented")
}

func (s *Session) RemovePosition(positionID string) error {
	return errors.New("not implemented")
}
