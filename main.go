package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"strconv"
	"time"
)

type App struct {
	DB *sql.DB
}

func main() {
	loadConfig()
	app := connectToDB()
	fmt.Println("Connected to Postgres!")

	router := gin.Default()
	GetMealsHandler(router, app)
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func ParsePage(c *gin.Context, url string, app *App) {
	col := colly.NewCollector(colly.AllowedDomains("studentenwerk.sh"))

	col.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	col.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})

	canteenID, _ := strconv.Atoi(c.Query("canteen"))
	col.OnHTML(selectDay(time.Now().Weekday().String()), func(e *colly.HTMLElement) {
		fmt.Println("Meals:")
		canteen := getCanteenInfo(canteenID, app.DB)
		meals := GetMealInformation(e, canteen)
		c.JSON(200, gin.H{"meals": meals})
	})

	err := col.Visit(url)
	if err != nil {
		fmt.Println("Error visiting URL:", err)
		c.JSON(500, gin.H{"error": "Error visiting URL"})
		return
	}
}

/*
Loads config for database connection. Currently only available to me. Sad.
*/
func loadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading config file: ", err)
	}
}

func connectToDB() *App {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.dbname"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Postgres!")

	return &App{DB: db}
}
