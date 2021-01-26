package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sysrex/wows/controllers"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("views/*.html")

	// Routes
	r.GET("/players", controllers.FindPlayers)
	r.POST("/players", controllers.CreatePlayers)

	// Run the server
	r.Run()
}
