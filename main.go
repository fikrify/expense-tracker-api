package main

import (
	"expense-tracker-api/internal/api"
	"expense-tracker-api/internal/config"
	"expense-tracker-api/internal/connection"
	"expense-tracker-api/internal/repository"
	"expense-tracker-api/internal/service"

	"github.com/gofiber/fiber/v3"
)

func main() {
	conf := config.Get()
	dbConnection := connection.GetDatabase(conf.Database)

	app := fiber.New()

	expenseRepository := repository.NewExpenseRepository(dbConnection)

	expenseService := service.NewExpenseService(expenseRepository)

	api.NewExpenseApi(app, expenseService)

	_ = app.Listen(conf.Server.Host + ":" + conf.Server.Port)
}
