package repository

import (
	"context"
	"expense-tracker-api/domain"

	"gorm.io/gorm"
)

type ExpenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) domain.ExpenseRepository {
	return &ExpenseRepository{
		db: db,
	}
}

func (e ExpenseRepository) FindAll(c context.Context) ([]*domain.Expense, error) {
	var expenses []*domain.Expense

	if err := e.db.
		WithContext(c).
		Find(&expenses).
		Error; err != nil {
		return nil, err
	}

	return expenses, nil
}

func (e ExpenseRepository) FindByID(c context.Context, id string) (*domain.Expense, error) {
	var expense domain.Expense
	if err := e.db.WithContext(c).
		First(&expense, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &expense, nil
}

func (e ExpenseRepository) Create(c context.Context, expense *domain.Expense) error {
	if err := e.db.WithContext(c).
		Create(expense).Error; err != nil {
		return err
	}
	return nil
}

func (e ExpenseRepository) Update(c context.Context, id string, expense *domain.Expense) error {
	return e.db.WithContext(c).
		Model(&domain.Expense{}).
		Where("id = ?", id).
		Updates(expense).Error
}

func (e ExpenseRepository) Delete(c context.Context, id string) error {
	return e.db.WithContext(c).
		Delete(&domain.Expense{}, "id = ?", id).Error
}
