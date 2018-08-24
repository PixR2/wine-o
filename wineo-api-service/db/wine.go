package db

import "errors"

type Wine struct {
	Winery     string  `json:"winery" gorethink:"id[0]"`
	Varietal   string  `json:"varietal" gorethink:"id[1]"`
	Vintage    int     `json:"vintage" gorethink:"id[2]"`
	BottleSize int     `json:"bottleSize" gorethink:"id[3]"`
	Region     *string `json:"region" gorethink:"region"`
	Peak       *int    `json:"peak" gorethink:"peak"`
	Hold       *int    `json:"hold" gorethink:"hold"`
	Drink      *int    `json:"drink" gorethink:"drink"`
	Taste      *int    `json:"taste" gorethink:"taste"`
	Price      *int    `json:"price" gorethink:"price"`
}

func (s *Session) CreateWine(wine *Wine) error {
	return errors.New("not implemented")
}

func (s *Session) UpdateWine(wine *Wine) error {
	return errors.New("not implemented")
}

func (s *Session) DeleteWine(winery string, varietal string, vintage int, bottleSize int) error {
	return errors.New("not implemented")
}

func (s *Session) GetWine(winery string, varietal string, vintage int, bottleSize int) (*Wine, error) {
	return nil, errors.New("not implemented")
}
