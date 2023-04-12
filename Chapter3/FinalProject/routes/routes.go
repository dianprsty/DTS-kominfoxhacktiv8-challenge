package routes

import (
	"github.com/dianprsty/DTS-kominfoxhacktiv8-challenge/tree/main/Chapter3/FinalProject/controller"
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

	r.GET("/photos", controller.GetAllPhoto)
	r.POST("/photos", controller.CreatePhoto)
	r.GET("/photos/:id", controller.GetPhotoById)
	r.PATCH("/photos/:id", controller.UpdatePhoto)
	r.DELETE("photos/:id", controller.DeletePhoto)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
