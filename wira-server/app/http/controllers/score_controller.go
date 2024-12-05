package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type ScoreController struct {
	//Dependent services
}

func NewScoreController() *ScoreController {
	return &ScoreController{
		//Inject services
	}
}

func (r *ScoreController) Index(ctx http.Context) http.Response {
	return nil
}	

func (r *ScoreController) Show(ctx http.Context) http.Response {
	return nil
}

func (r *ScoreController) Store(ctx http.Context) http.Response {
	return nil
}

func (r *ScoreController) Update(ctx http.Context) http.Response {
	return nil
}

func (r *ScoreController) Destroy(ctx http.Context) http.Response {
	return nil
}
