package models

type Expense struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}
