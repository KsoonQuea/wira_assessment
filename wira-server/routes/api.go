package routes

import (
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
)

func Api() {
	userController 	:= controllers.NewUserController()
	accController 	:= controllers.NewAccountController()
	
	facades.Route().Get("/users/{id}", userController.Show)
	
	facades.Route().Get("/accounts", accController.Index)

}
