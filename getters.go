package main

import (
	"GOpenMensa/models"
	"database/sql"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strings"
	"time"
)

func parseMealInformation(e *colly.HTMLElement) []models.Meal {
	meals := make([]models.Meal, 0, len(".mensa_menu_detail"))
	var meal models.Meal
	lastUpdate, _ := parseDate(e)
	e.ForEach(".mensa_menu_detail", func(i int, element *colly.HTMLElement) {
		meal = models.Meal{
			Name:       extractMealName(element),
			Vegan:      dataArtenContains(element, "ve"),
			Vegetarian: dataArtenContains(element, "vg"),
			LastUpdate: lastUpdate,
		}
		meals = append(meals, meal)
	})
	return meals
}

func parseDate(e *colly.HTMLElement) (time.Time, error) {
	day := e.Attr("data-day")
	date, err := time.Parse("2006-01-02", day)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

func getMealsOnDate(date time.Time, db *sql.DB) ([]models.Meal, error) {
	query := "SELECT * FROM meals WHERE last_update=$1"
	formattedDate := date.Format("2006-01-02")
	rows, err := db.Query(query, formattedDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var meals []models.Meal

	for rows.Next() {
		var meal models.Meal
		canteen := getCanteenInfo(meal.Canteen.ID, db)
		err := rows.Scan(&meal.ID, &meal.Name, &meal.Vegan, &meal.Vegetarian, &meal.Canteen, &meal.Additives, &meal.Allergens, &meal.FoodPreferences)
		if err != nil {
			return nil, err
		}
		meals = append(meals, meal)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return meals, nil
}

func getCanteenInfo(canteenID int, db *sql.DB) models.Canteen {
	query := "SELECT * FROM canteens WHERE id=$1"
	rows, err := db.Query(query, canteenID)
	if err != nil {
		log.Fatal(err)
	}

	var canteen models.Canteen

	if rows.Next() {
		err = rows.Scan(&canteen.ID, &canteen.Name, &canteen.City)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("No rows returned")
	}

	return canteen
}

func parseAllergens(e *colly.HTMLElement) []models.Allergen {
	allergens := make([]models.Allergen, 0)
	var allergen models.Allergen
	e.ForEach(".filterbutton", func(i int, element *colly.HTMLElement) {
		fmt.Println(element.ChildText("span:not(.abk)"))
		allergen = models.Allergen{
			Abbreviation: element.ChildText("span.abk"),
			Name:         element.ChildText("span:not(.abk)"),
		}
		allergens = append(allergens, allergen)
	})
	return allergens
}

func insertAllergensIntoDB(allergens []models.Allergen, db *sql.DB) error {
	for _, allergen := range allergens {
		stmt, err := db.Prepare("INSERT INTO allergens (abbreviation, name) VALUES ($1, $2)")
		if err != nil {
			return err
		}

		_, err = stmt.Exec(strings.TrimSpace(allergen.Abbreviation), strings.TrimSpace(allergen.Name))
		if err != nil {
			return err
		}

	}
	return nil
}

func insertMealsIntoDB(meals []models.Meal, db *sql.DB) error {
	for _, meal := range meals {
		date := time.Now()
		formattedDate := date.Format("2006-01-02")
		stmt, err := db.Prepare("INSERT INTO meals (canteen_id, name, vegetarian, vegan, last_update) VALUES ($1, $2, $3, $4, $5)")
		if err != nil {
			return err
		}

		_, err = stmt.Exec(meal.Canteen.ID, strings.TrimSpace(meal.Name), meal.Vegetarian, meal.Vegan, formattedDate)
		if err != nil {
			return err
		}

	}
	return nil
}

/*
func getPriceByGroup(e *colly.HTMLElement) models.PriceByGroup {
	priceList := e.ChildText(".menu_preis")
	prices := strings.Split(priceList, "/")

	var priceByGroup models.PriceByGroup

	priceByGroup.Students = strings.TrimSpace(prices[0])
	priceByGroup.Employees = strings.TrimSpace(prices[1])
	priceByGroup.Guests = strings.TrimSpace(prices[2])

	return priceByGroup
}
*/
