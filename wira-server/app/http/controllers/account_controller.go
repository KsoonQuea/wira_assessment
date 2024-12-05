package controllers

import (
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
)

type AccountController struct {
	AccountRepo *models.AccountRepository

	//Dependent services
}

func NewAccountController() *AccountController {
	return &AccountController{
		//Inject services
	}
}

func (r *AccountController) Index(ctx http.Context)  {

	var accounts []models.Account

	repo := &models.AccountRepository{}
	
	// Retrieve accounts
	accounts, err := repo.Index()
	if err != nil {
		// Handle error
		ctx.Response().Status(500).Json(map[string]string{
			"error": "Failed to retrieve accounts",
		})
		return
	}

	// Return successful response
	ctx.Response().Json(200, accounts)
}	

func (r *AccountController) Show(ctx http.Context) http.Response {

	return ctx.Response().Success().Json(http.Json{
		"acc_show": "This is account show page",
	})
}

func (r *AccountController) Store(ctx http.Context) http.Response {
	return nil
}

func (r *AccountController) Update(ctx http.Context) http.Response {
	return nil
}

func (r *AccountController) Destroy(ctx http.Context) http.Response {
	return nil
}
