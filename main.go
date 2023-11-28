package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func main() {
	router := gin.Default()

	handleMeals(router)
	handleLandingPage(router)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func ParsePage(c *gin.Context, url string, day string) {
	col := colly.NewCollector(colly.AllowedDomains("studentenwerk.sh"))

	col.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	col.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})

	col.OnHTML(selectDay(day), func(e *colly.HTMLElement) {
		fmt.Println("Meals:")
		meals := getMealInformation(e)
		c.JSON(200, gin.H{"meals": meals})
	})

	err := col.Visit(url)
	if err != nil {
		fmt.Println("Error visiting URL:", err)
		c.JSON(500, gin.H{"error": "Error visiting URL"})
		return
	}
}
