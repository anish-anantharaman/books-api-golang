package models

type Books struct {
	Id		  int		`json:"id"`       // book id
	Name      string    `json:"name"`	  // book name
	Author    string    `json:"author"`   // author name
	Category  string    `json:"category"` // book category
	Price     float32   `json:"price"`    // book price
	Isbn      string    `json:"isbn"`     // book isbn
}