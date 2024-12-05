package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type CharacterController struct {
	//Dependent services
}

func NewCharacterController() *CharacterController {
	return &CharacterController{
		//Inject services
	}
}

func (r *CharacterController) Index(ctx http.Context) http.Response {
	return nil
}	

func (r *CharacterController) Show(ctx http.Context) http.Response {
	return nil
}

func (r *CharacterController) Store(ctx http.Context) http.Response {
	return nil
}

func (r *CharacterController) Update(ctx http.Context) http.Response {
	return nil
}

func (r *CharacterController) Destroy(ctx http.Context) http.Response {
	return nil
}
