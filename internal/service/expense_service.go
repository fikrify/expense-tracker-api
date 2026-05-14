package service

import (
	"context"
	"expense-tracker-api/domain"
	"expense-tracker-api/dto"

	"github.com/google/uuid"
)

type ExpenseService struct {
	expenseRepository domain.ExpenseRepository
}

func NewExpenseService(expenseRepository domain.ExpenseRepository) domain.ExpenseService {
	return &ExpenseService{
		expenseRepository: expenseRepository,
	}
}

func (e ExpenseService) GetAllExpenses(c context.Context) ([]dto.ExpenseData, error) {
	expenses, err := e.expenseRepository.FindAll(c)
	if err != nil {
		return nil, err
	}

	var expensesDTO []dto.ExpenseData
	for _, expense := range expenses {
		expensesDTO = append(expensesDTO, dto.ExpenseData{
			ID:        expense.ID,
			Name:      expense.Name,
			Amount:    expense.Amount,
			Category:  expense.Category,
			CreatedAt: expense.CreatedAt.String(),
			UpdatedAt: expense.UpdatedAt.String(),
		})
	}

	return expensesDTO, nil
}

func (e ExpenseService) GetExpenseByID(c context.Context, id string) (dto.ExpenseData, error) {
	expense, err := e.expenseRepository.FindByID(c, id)
	if err != nil {
		return dto.ExpenseData{}, err
	}

	return dto.ExpenseData{
		ID:        expense.ID,
		Name:      expense.Name,
		Amount:    expense.Amount,
		Category:  expense.Category,
		CreatedAt: expense.CreatedAt.String(),
		UpdatedAt: expense.UpdatedAt.String(),
	}, nil
}

func (e ExpenseService) CreateExpense(c context.Context, expense dto.CreateExpenseData) error {
	newExpense := &domain.Expense{
		ID:       uuid.NewString(),
		Name:     expense.Name,
		Amount:   expense.Amount,
		Category: expense.Category,
	}

	return e.expenseRepository.Create(c, newExpense)
}

func (e ExpenseService) UpdateExpense(c context.Context, id string, expense dto.UpdateExpenseData) error {
	_, err := e.expenseRepository.FindByID(c, id)
	if err != nil {
		return err
	}

	updateExpense := &domain.Expense{
		Name:     expense.Name,
		Amount:   expense.Amount,
		Category: expense.Category,
	}

	return e.expenseRepository.Update(c, id, updateExpense)
}

func (e ExpenseService) DeleteExpense(c context.Context, id string) error {
	_, err := e.expenseRepository.FindByID(c, id)
	if err != nil {
		return err
	}

	return e.expenseRepository.Delete(c, id)
}
