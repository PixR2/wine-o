package model

type Wine struct {
	Winery string `gorethink:"id[0]"`

	Vintage   int `gorethink:"id[1]"`
	Peak      int `gorethink:"peak"`
	HoldUntil int `gorethink:"hold"`
	DrinkBy   int `gorethink:"drink"`
	TasteNext int `gorethink:"taste"`

	BottleSize int `gorethink:"id[2]"`
	Price      int `gorethink:"price"`

	Region   string `gorethink:"region"`
	Varietal string `gorethink:"varietal"`
}
