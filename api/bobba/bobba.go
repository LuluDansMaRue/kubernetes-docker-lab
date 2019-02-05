package bobba

// Bobba structure
type Bobba struct {
	ID     int     `json:"id,omitempty"`
	Name   string  `json:"name"`
	Price  float32 `json:"price,omitempty"`
	Shop   string  `json:"shop"`
	Flavor string  `json:"flavor"`
	Calory float32 `json:"calory,omitempty"`
}
