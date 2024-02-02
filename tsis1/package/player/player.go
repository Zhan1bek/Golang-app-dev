package player

type Player struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Position    string `json:"position"`
	Nation      string `json:"nation"`
	ShirtNumber int    `json:"shirtNumber"`
}
