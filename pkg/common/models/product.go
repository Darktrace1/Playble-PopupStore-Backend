package models

type Product struct {
	Id 			string 	`json:"id"`
	Name 		string 	`json:"name"`
	Provider 	string 	`json:"provider"`
	Price 		int 	`json:"price"`
	Image 		string	`json:"image"`
}