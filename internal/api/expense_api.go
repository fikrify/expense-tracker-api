package api

import (
	"context"
	"expense-tracker-api/domain"
	"expense-tracker-api/dto"
	"expense-tracker-api/internal/util"
	"time"

	"github.com/gofiber/fiber/v3"
)

type ExpenseApi struct {
	expenseService domain.ExpenseService
}

func NewExpenseApi(app *fiber.App, service domain.ExpenseService) {
	ea := ExpenseApi{
		expenseService: service,
	}

	app.Get("/expenses", ea.GetAllExpenses)
	app.Post("/expenses", ea.CreateExpense)
	app.Get("/expenses/:id", ea.GetExpenseByID)
	app.Put("/expenses/:id", ea.UpdateExpense)
	app.Delete("/expenses/:id", ea.DeleteExpense)
}

func (ea ExpenseApi) GetAllExpenses(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	expenses, err := ea.expenseService.GetAllExpenses(c)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.NewResponseError(err.Error()))
	}

	return ctx.JSON(dto.NewResponseSuccess("Get all expenses succeed.", expenses))
}

func (ea ExpenseApi) GetExpenseByID(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	expenseId := ctx.Params("id")
	if expenseId == "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(dto.NewResponseError("Invalid expense id."))
	}

	expense, err := ea.expenseService.GetExpenseByID(c, expenseId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.NewResponseError(err.Error()))
	}

	return ctx.JSON(dto.NewResponseSuccess("Get expense succeed.", expense))
}

func (ea ExpenseApi) CreateExpense(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var expenseData dto.CreateExpenseData
	if err := ctx.Bind().Body(&expenseData); err != nil {
		return ctx.SendStatus(fiber.StatusUnprocessableEntity)
	}
	fails := util.Validate(expenseData)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(dto.NewResponseErrorWithData("Validation failed.", fails))
	}

	if err := ea.expenseService.CreateExpense(c, expenseData); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(dto.NewResponseError(err.Error()))
	}

	return ctx.
		Status(fiber.StatusCreated).
		JSON(dto.NewResponseSuccess("Create expense succeed.", expenseData))
}

func (ea ExpenseApi) UpdateExpense(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	expenseId := ctx.Params("id")
	if expenseId == "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(dto.NewResponseError("Invalid expense id."))
	}

	var expenseData dto.UpdateExpenseData
	if err := ctx.Bind().Body(&expenseData); err != nil {
		return ctx.SendStatus(fiber.StatusUnprocessableEntity)
	}
	fails := util.Validate(expenseData)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(dto.NewResponseErrorWithData("Validation failed.", fails))
	}

	if err := ea.expenseService.UpdateExpense(c, expenseId, expenseData); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(dto.NewResponseError(err.Error()))
	}

	return ctx.JSON(dto.NewResponseSuccess("Update expense succeed.", expenseData))
}

func (ea ExpenseApi) DeleteExpense(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	expenseId := ctx.Params("id")
	if expenseId == "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(dto.NewResponseError("Invalid expense id."))
	}

	if err := ea.expenseService.DeleteExpense(c, expenseId); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(dto.NewResponseError(err.Error()))
	}

	return ctx.
		Status(fiber.StatusOK).
		JSON(dto.NewResponseSuccess("Delete expense succeed.", ""))
}
