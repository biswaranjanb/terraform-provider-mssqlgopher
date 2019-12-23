package gopherbiswa

//Information models the Information structure required to do db operations
type Information struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     int    `json:"age"`
	City    string `json:"city"`
}
