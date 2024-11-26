package test

import (
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID int
	Item string
	Completed bool
}

var todos = []todo {
	{ID: 1, Item: "Item 1", Completed: false },
	{ID: 2, Item: "Item 22222", Completed: false },
}

func test(){

	router := gin.Default()

	router.Run("localhost:9090")

}