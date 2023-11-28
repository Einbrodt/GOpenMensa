package models

type MensaData struct {
	City         string
	LocationType string
	Day          string
	Open         bool
	Meals        []Meal
}

type Meal struct {
	Name         string
	Price        string
	Vegan        bool
	Vegetarian   bool
	Chicken      bool
	Fish         bool
	Pig          bool
	Cow          bool
	SH_Plate     bool
	PriceByGroup PriceByGroup
	Allergens    []Allergen
}

type Allergen struct {
	Code string
	Name string
}

type PriceByGroup struct {
	Students  string
	Employees string
	Guests    string
}
