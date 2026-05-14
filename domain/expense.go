package domain

import (
	"context"
	"expense-tracker-api/dto"
	"time"

	"gorm.io/gorm"
)

type Expense struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ExpenseRepository interface {
	FindAll(c context.Context) ([]*Expense, error)
	FindByID(c context.Context, id string) (*Expense, error)
	Create(c context.Context, expense *Expense) error
	Update(c context.Context, expense *Expense) error
	Delete(c context.Context, id string) error
}

type ExpenseService interface {
	GetAllExpenses(c context.Context) ([]dto.ExpenseData, error)
	GetExpenseByID(c context.Context, id string) (dto.ExpenseData, error)
	CreateExpense(c context.Context, expense dto.CreateExpenseData) error
	UpdateExpense(c context.Context, id string, expense dto.UpdateExpenseData) error
	DeleteExpense(c context.Context, id string) error
}
