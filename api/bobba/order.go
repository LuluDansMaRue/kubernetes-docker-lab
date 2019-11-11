package bobba

// Order structure
type Order struct {
	ID       int    `json:"id, omitempty"`
	Bobba_id int    `json:"bobba_id"`
	Client   string `json:"client_name"`
}
