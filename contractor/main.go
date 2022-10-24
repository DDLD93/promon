package main

import (
	"fmt"
	"log"
	"os"

	controller "github.com/ddld93/promon/contractor/src/controller"
	"github.com/ddld93/promon/contractor/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("APP_PORT")
	host := os.Getenv("DATABASE_HOST")
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "3000"
	}
	contractorCtrl := controller.ConnectDB(host, 27017)
	route := routes.ContractorRoute{ContractorCtrl: contractorCtrl}
	router := gin.New()
	//router.Use(middleware.Logger())
	v1 := router.Group("/api/v1/contractor")
	{
		v1.POST("/", route.Create())
		v1.GET("/", route.GetAll())
		v1.GET("/:id", route.GetOne())
		v1.PUT("/:id", route.GetOne())
		v1.DELETE("/:id", route.GetOne())
		// v1.POST("/post", route.FormTest())
	}

	err := router.Run(fmt.Sprintf("%s:%v", "0.0.0.0", port))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server running on port %s", port)
}
