package models

type User struct {
	Id   int    `json:"id"`
	Age  int    `json:"age"  example:"10"`
	Name string `json:"name" example:"Bill"`
}
