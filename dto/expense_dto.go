package dto

type ExpenseData struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Amount    float64 `json:"amount"`
	Category  string  `json:"category"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type CreateExpenseData struct {
	Name     string  `json:"name" validate:"required"`
	Amount   float64 `json:"amount" validate:"required"`
	Category string  `json:"category"`
}

type UpdateExpenseData struct {
	Name     string  `json:"name"`
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
}
