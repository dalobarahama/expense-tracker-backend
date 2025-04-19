package models

type Expense struct {
	ID          int     `json:"id" gorm:"primary_key"`
	Title       string  `json:"title"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	UserID      uint    `json:"user_id"`
}
