package models

type Image struct {
	Id       int    `json:"id"`
	Filename string `json:"filename" example:"Image.jpg"`
}
