package routers

import (
	"wira-backend/controllers"

	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type AccountRouter struct{
	DB *gorm.DB
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Change '*' to your frontend's URL in production
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight OPTIONS request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Proceed to the next middleware or handler
		next.ServeHTTP(w, r)
	})
}

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.Use(CORSMiddleware)

	account_controller 		:= controllers.NewAccountController()
	dashboard_controller 	:= controllers.NewDashboardController()

	r.HandleFunc("/accounts", account_controller.Index).Methods("GET")

	r.HandleFunc("/dashboard/ranking_table", dashboard_controller.RankingTable).Methods("GET")
	// r.HandleFunc("/items", account_controller.AddItem).Methods("POST")
	// r.HandleFunc("/items/{id}", account_controller.UpdateItem).Methods("PUT")
	// r.HandleFunc("/items/{id}", account_controller.DeleteItem).Methods("DELETE")

	return r
}
