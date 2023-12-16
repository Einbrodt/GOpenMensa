package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetMealsHandler(router *gin.Engine, app *App) {
	router.GET("/meals", func(c *gin.Context) {
		canteen, _ := strconv.Atoi(c.Query("canteen"))
		url := selectMensa(canteen)
		ParsePage(c, url, app)
	})
}

func handleLandingPage(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*")

	router.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "landing.html", nil)
	})
}
