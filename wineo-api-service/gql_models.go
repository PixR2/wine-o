package wineo_api_service

type User struct {
	Email string `json:"email" gorethink:"id"`

	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
}

type Wine struct {
	Winery     string  `json:"winery"`
	Varietal   string  `json:"varietal"`
	Vintage    int     `json:"vintage"`
	BottleSize int     `json:"bottleSize"`
	Region     *string `json:"region"`
	Peak       *int    `json:"peak"`
	Hold       *int    `json:"hold"`
	Drink      *int    `json:"drink"`
	Taste      *int    `json:"taste"`
	Price      *int    `json:"price"`
}

type WineCollection struct {
	ID    string  `json:"id" gorethink:"price"`
	Name  *string `json:"name" gorethink:"price"`
	Owner []User  `json:"owner" gorethink:"price"`
}

type WinePosition struct {
	CollectionID string  `json:"collectionID"`
	Compartment  *string `json:"compartment"`
	Wine         `json:"wine"`
	NumBottles   int `json:"numBottles"`
}
