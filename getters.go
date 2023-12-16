package main

import (
	"GOpenMensa/models"
	"database/sql"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strings"
)

func parseMealInformation(e *colly.HTMLElement, canteen models.Canteen) []models.Meal {
	meals := make([]models.Meal, 0, len(".mensa_menu_detail"))
	var meal models.Meal
	e.ForEach(".mensa_menu_detail", func(i int, element *colly.HTMLElement) {
		meal = models.Meal{
			Name:       extractMealName(element),
			Price:      element.ChildText(".menu_preis"),
			Vegan:      dataArtenContains(element, "ve"),
			Vegetarian: dataArtenContains(element, "vg"),
			Canteen:    canteen,
		}
		meals = append(meals, meal)
	})
	return meals
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

func getAllergenInfo(allergenCodes string) []models.Allergen {
	allergenCodes = strings.Trim(allergenCodes, "()")
	codes := strings.Split(allergenCodes, ",")
	allergens := make([]models.Allergen, 0, len(codes))

	for _, code := range codes {
		code = strings.TrimSpace(code)

		name, exists := allergenCodeToName[code]
		if !exists {
			name = "Unknown"
		}

		allergen := models.Allergen{
			Name: name,
		}

		allergens = append(allergens, allergen)
	}
	return allergens
}

func parseAllergens(e *colly.HTMLElement) []models.Allergen {
	allergens := make([]models.Allergen, 0)
	var allergen models.Allergen
	e.ForEach(".filterbutton.span", func(i int, element *colly.HTMLElement) {
		fmt.Println(element.ChildText(".abk"))
		allergen = models.Allergen{
			Abbreviation: element.ChildText(".abk"),
			Name:         element.ChildText("not(.abk)"),
		}
		allergens = append(allergens, allergen)
	})
	return allergens

	//allergenCodes := e.ChildText(".menu_name.mensa_zusatz")
	//return getAllergenInfo(allergenCodes)
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
