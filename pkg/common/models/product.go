package models

type Product struct {
	Cid 		string 	`json:"cid"`
	Name 		string 	`json:"name"`
	Provider 	string 	`json:"provider"`
	Price 		int 	`json:"price"`
	Image 		string	`json:"image"`
}