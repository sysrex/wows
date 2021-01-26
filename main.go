package main

import (
	"./controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("views/*.html")

	// Routes
	r.GET("/", controllers.Home)

	// Run the server
	r.Run(":8080")
}
