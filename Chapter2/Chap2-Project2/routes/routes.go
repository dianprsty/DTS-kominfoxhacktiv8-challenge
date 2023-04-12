package routes

import (
	"chap2-project2/controller"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/books", controller.GetAllBook)
	r.POST("/books", controller.CreateBook)
	r.GET("/books/:id", controller.GetBookById)
	r.PUT("/books/:id", controller.UpdateBook)
	r.DELETE("/books/:id", controller.DeleteBook)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
