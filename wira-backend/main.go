// package main

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"

// 	"wira-backend/connection"
// 	"wira-backend/services"

// 	"github.com/gorilla/mux"
// )

// func main() {
// 	connection.InitDB()
// 	defer connection.DB.Close()

// 	r := mux.NewRouter()

// 	r.HandleFunc("/api/login", loginHandler).Methods("POST")

// 	log.Println("Server is running on http://localhost:8080")
// 	log.Fatal(http.ListenAndServe(":8080", r))
// }

// func loginHandler(w http.ResponseWriter, r *http.Request) {
// 	var creds struct {
// 		Username string `json:"username"`
// 		Password string `json:"password"`
// 	}

// 	err := json.NewDecoder(r.Body).Decode(&creds)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	token, err := services.LoginService(creds.Username, creds.Password)
// 	if err != nil {
// 		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(map[string]string{"token": token})
// }

package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"errors"
	"strconv"
	"gorm.io/gorm"

	"github.com/KsoonQuea/wira_assessment/wira-backend/models"
	"github.com/KsoonQuea/wira_assessment/wira-backend/services"
	"github.com/KsoonQuea/wira_assessment/wira-backend/connection"

	// "wira-backend/models"
	// "wira-backend/connection"
	// "wira-backend/services"
)

type Repository struct {
	DB *gorm.DB
}

type todo struct {
	ID int
	Item string
	Completed bool
}

// func (repos *Repository) CreateBook() {

// }

func (repos *Repository) GetAccounts(context *fiber.Ctx) error {
	accModels := &[]models.Accounts{}

	err := repos.DB.Find(accModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get books"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message"	: "books fetched successfully",
		"data"		: accModels,
	})
	return nil
}

func (repos *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	// api.Post("/create_books", repos.CreateBook)
	// api.Delete("delete_book/:id", repos.DeleteBook)
	// api.Get("/get_books/:id", repos.GetBookByID)
	api.Get("/books", repos.GetAccounts)
}

var todos = []todo {
	{ID: 1, Item: "Item 1", Completed: false },
	{ID: 2, Item: "Item 22222", Completed: false },
}

func getTodos (context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func findTodoByID(id int) (*todo, error) {
	for todo_key, todo := range todos {
		if todo.ID == id {
			return &todos[todo_key], nil
		}
	}
	return nil, errors.New("todo not found")
}

func getTodo (context *gin.Context) {
	strid := context.Param("id")

	id, err := strconv.Atoi(strid)

	todo, err := findTodoByID(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		context.IndentedJSON(http.StatusNoContent, gin.H{"message" : "Message Not Found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func main(){

	// router := gin.Default()
	// router.GET("/api/todos", getTodos)
	// router.GET("/api/todo/:id", getTodo)
	// router.Run("localhost:9090")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &connection.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := connection.NewConnection(config)

	if err != nil {
		log.Fatal("could not load the database")
	}
	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	repos := Repository{
		DB: db,
	}
	app := wira.New()

	fmt.Println(app)

	repos.SetupRoutes(app)
	app.Listen(":9090")

}