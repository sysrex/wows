package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home Page
func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}
