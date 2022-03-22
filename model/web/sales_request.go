package web

type SalesAddRequest struct {
	Name           string `json:"Name"`
	Email          string `json:"Email"`
	Quantity       int    `json:"Quantity"`
	Book_ID        int    `json:"Book_ID"`
	Author_ID      int
	Price_Per_Unit float64
	Price_Total    float64
}
