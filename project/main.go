package main

import (
	"fmt"
	"log"
	"os"

	controller"github.com/ddld93/promon/project/src/controller"
	"github.com/ddld93/promon/project/src/routes"

	"github.com/gin-gonic/gin"
)
func main()  {
	port := os.Getenv("PORT")
	host := os.Getenv("DATABASE_HOST")
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "3000"
	}
	projectCtrl := controller.ConnectDB(host, 27017)
	route := routes.ProjectRoute{ProjectCtrl: projectCtrl}
	router := gin.New()
	//router.Use(middleware.Logger())
	v1 := router.Group("/api/v1/project")
		{
			v1.POST("/", route.Create())
			v1.GET("/", route.GetAll())
			v1.GET("/:id", route.GetOne())
			v1.PATCH("/:id", route.GetOne())
			v1.DELETE("/:id", route.GetOne())
			// v1.POST("/post", route.FormTest())
		}

	err := router.Run(fmt.Sprintf("%s:%v", "0.0.0.0", port))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server running on port %s", port)
}