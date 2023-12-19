package models

import "time"

type Prices struct {
	Currency      string    `json:"currency"`
	StudentPrice  float64   `json:"studentPrice"`
	LecturerPrice float64   `json:"lecturerPrice"`
	GuestPrice    float64   `json:"guestPrice"`
	LastUpdate    time.Time `json:"lastUpdate"`
}

type Canteen struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

type Meal struct {
	ID              int              `json:"id"`
	Name            string           `json:"name"`
	Vegetarian      bool             `json:"vegetarian"`
	Vegan           bool             `json:"vegan"`
	LastUpdate      time.Time        `json:"lastUpdate"`
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
