package entities

type User struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
}
