package main

import (
	"GOpenMensa/models"
	"github.com/gocolly/colly"
	"strings"
)

func getMealInformation(e *colly.HTMLElement) []models.Meal {
	meals := make([]models.Meal, 0, len(".mensa_menu_detail"))
	var meal models.Meal
	e.ForEach(".mensa_menu_detail", func(i int, element *colly.HTMLElement) {
		meal = models.Meal{
			Name:         extractMealName(element),
			Price:        element.ChildText(".menu_preis"),
			Vegan:        dataArtenContains(element, "ve"),
			Vegetarian:   dataArtenContains(element, "vg"),
			Chicken:      dataArtenContains(element, "G"),
			Fish:         dataArtenContains(element, "AGF"),
			Pig:          dataArtenContains(element, "AGS"),
			SH_Plate:     dataArtenContains(element, "SHT"),
			Allergens:    getAllergens(element),
			PriceByGroup: getPriceByGroup(element),
		}
		meals = append(meals, meal)
	})
	return meals
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
			Code: code,
			Name: name,
		}

		allergens = append(allergens, allergen)
	}
	return allergens
}

func getAllergens(e *colly.HTMLElement) []models.Allergen {
	allergenCodes := e.ChildText(".menu_name.mensa_zusatz")
	return getAllergenInfo(allergenCodes)
}

func getPriceByGroup(e *colly.HTMLElement) models.PriceByGroup {
	priceList := e.ChildText(".menu_preis")
	prices := strings.Split(priceList, "/")

	var priceByGroup models.PriceByGroup

	priceByGroup.Students = strings.TrimSpace(prices[0])
	priceByGroup.Employees = strings.TrimSpace(prices[1])
	priceByGroup.Guests = strings.TrimSpace(prices[2])

	return priceByGroup
}
