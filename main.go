package main

import (
	"github.com/shundahoy/techready/controller"
	"github.com/shundahoy/techready/db"
	"github.com/shundahoy/techready/repository"
	"github.com/shundahoy/techready/router"
	"github.com/shundahoy/techready/usecase"
	"github.com/shundahoy/techready/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
