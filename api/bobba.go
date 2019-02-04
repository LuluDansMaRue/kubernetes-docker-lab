package main

type ingredient struct {
	Name   string `json:"name"`
	Calory int    `json:"calory"`
}

// Bobba structure
type Bobba struct {
	Name        string       `json:"name"`
	Ingredients []ingredient `json:"ingredients"`
	Price       float32      `json:"price"`
	Shop        string       `json:"shop"`
	Flavor      string       `json:"flavor"`
}
