package main

import (
	//"net/http"

	"bookstore/configs"
	"bookstore/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(router)
	routes.BookRoute(router)
	routes.CategoryRoute(router)
	routes.AuthorRoute(router)
	routes.AdminRoute(router)
	routes.OrderRoute(router)
	routes.CartRoute(router)
	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"data": "Hello from Gin-gonic & MongoDB",
	// 	})
	// })

	router.Run()
}
