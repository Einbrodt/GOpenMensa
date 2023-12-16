package models

type Price struct {
	Value    int    `json:"value"`
	Currency string `json:"currency"`
}

type Canteen struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

type Meal struct {
	Price           string           `json:"price"`
	Name            string           `json:"name"`
	Vegetarian      bool             `json:"vegetarian"`
	Vegan           bool             `json:"vegan"`
	Canteen         Canteen          `json:"canteen"`
	Allergens       []Allergen       `json:"allergens"`
	FoodPreferences []FoodPreference `json:"foodPreferences"`
	Additives       []Additive       `json:"additives"`
}

type Allergen struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}

type FoodPreference struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}

type Additive struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Number int    `json:"number"`
}
