package main

import (
	"Library/controllers"
	"Library/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // init the router
	models.ConnectDataBase() // connect to GORM (in memory SqlLite)

	//get all, and cruds
	r.GET("/books", controllers.GetBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.GetBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.Run()
}