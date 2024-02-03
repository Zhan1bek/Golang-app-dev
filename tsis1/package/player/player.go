package player

type Player struct {
	ID          string `json:"id"`
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	Age         int    `json:"age"`
	Position    string `json:"position"`
	Nation      string `json:"nation"`
	ShirtNumber int    `json:"shirtNumber"`
}
