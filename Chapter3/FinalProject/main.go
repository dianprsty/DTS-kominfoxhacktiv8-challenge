package main

import (
	"github.com/dianprsty/DTS-kominfoxhacktiv8-challenge/tree/main/Chapter3/FinalProject/config"
	"github.com/dianprsty/DTS-kominfoxhacktiv8-challenge/tree/main/Chapter3/FinalProject/docs"
	"github.com/dianprsty/DTS-kominfoxhacktiv8-challenge/tree/main/Chapter3/FinalProject/routes"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/
func main() {

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Movie."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := routes.SetupRouter(db)
	r.Run()
}
