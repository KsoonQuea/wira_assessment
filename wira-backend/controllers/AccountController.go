package controllers

import (
	"encoding/json"
	"net/http"

	"wira-backend/models"
)

type AccountController struct {
	// AccountRepo *models.AccountRepository

	//Dependent services
}

func NewAccountController() *AccountController {
	return &AccountController{
		//Inject services
	}
}

func (ctl *AccountController) Index(w http.ResponseWriter, r *http.Request)  {

	items, err := models.GetAccounts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(items)
}	

// func (r *AccountController) Show(ctx http.Context) http.Response {

// 	return ctx.Response().Success().Json(http.Json{
// 		"acc_show": "This is account show page",
// 	})
// }

// func (r *AccountController) Store(ctx http.Context) http.Response {
// 	return nil
// }

// func (r *AccountController) Update(ctx http.Context) http.Response {
// 	return nil
// }

// func (r *AccountController) Destroy(ctx http.Context) http.Response {
// 	return nil
// }
