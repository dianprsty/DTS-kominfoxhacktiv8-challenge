package routers

import (
	"Chap2-Challenge2/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/books", controllers.CreateBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.GET("/books/:id", controllers.GetBook)
	router.GET("/books", controllers.GetAllBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	return router
}
