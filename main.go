package main

import (
	"go-gin-crud/config"
	"go-gin-crud/routes"

	"github.com/gin-gonic/gin"
)


func main() {
	// initialize gin router
	router := gin.Default()

	//initialize db connection
	db := config.InitDB()
	defer db.Close()

	//register routes with the router and pass the db connection
	routes.RegisterRoutes(router, db)

	//start http server
	router.Run(":8080")
}