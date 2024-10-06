package models

// User представляет данные пользователя.
// swagger:model
type User struct {
	// Имя пользователя
	// example: John Doe
	Name string `json:"name"`
	// Возраст пользователя
	// example: 30
	Age int `json:"age"`
}
