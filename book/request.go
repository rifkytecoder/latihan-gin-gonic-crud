package book

import "encoding/json"

// ? Struct untuk PostBookHandler `luar kedalam from postman ke code`
// Book Holder/contains information untuk body di postman
type BookRequest struct {
	Title string      `json:"title" binding:"required"`        //validate
	Price json.Number `json:"price" binding:"required,number"` //json string/int yg penting angka
	//SubTitle string `json:"sub_title"`
	Description string      `json:"description" binding:"required"`
	Rating      json.Number `json:"rating" binding:"required"`
	Discount    json.Number `json:"discount" binding:"required"`
}
