package main

import (
	"GOpenMensa/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleMeals(router *gin.Engine) {
	router.GET("/meals", func(c *gin.Context) {
		day := c.Query("day")
		city := c.Query("city")
		locationType := c.Query("locationType")
		mensaData := models.MensaData{
			City:         city,
			LocationType: locationType,
			Day:          day,
		}
		selectMensa(c, mensaData, day)
	})
}

func handleLandingPage(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*")

	router.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "landing.html", nil)
	})
}
