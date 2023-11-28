package main

import (
	"GOpenMensa/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

var mensaURLs = map[string]map[string]string{
	"Flensburg": {
		"Mensa":     "https://studentenwerk.sh/de/essen-uebersicht?ort=2&mensa=7#mensaplan",
		"Cafeteria": "https://studentenwerk.sh/de/essen-uebersicht?ort=2&mensa=14#mensaplan",
	},
	"Heide": {
		"": "https://studentenwerk.sh/de/essen-uebersicht?ort=4&mensa=14#mensaplan",
	},
	"Kiel": {
		"Mensa_1":          "https://studentenwerk.sh/de/essen-uebersicht?ort=1&mensa=1#mensaplan",
		"Mensa_2":          "https://studentenwerk.sh/de/essen-uebersicht?ort=1&mensa=2#mensaplan",
		"Mensa_Gaarden":    "https://studentenwerk.sh/de/essen-uebersicht?ort=1&mensa=3#mensaplan",
		"Mensa_Kesselhaus": "https://studentenwerk.sh/de/essen-uebersicht?ort=1&mensa=4#mensaplan",
		"Mensa_Schwentine": "https://studentenwerk.sh/de/essen-uebersicht?ort=1&mensa=5#mensaplan",
		"Cafeteria":        "https://studentenwerk.sh/de/essen-uebersicht?ort=1&mensa=6#mensaplan",
	},
	"Lübeck": {
		"Mensa":           "https://studentenwerk.sh/de/essen-uebersicht?ort=3&mensa=8#mensaplan",
		"Musikhochschule": "https://studentenwerk.sh/de/essen-uebersicht?ort=3&mensa=9#mensaplan",
	},
	"Wedel": {
		"": "https://studentenwerk.sh/de/essen-uebersicht?ort=5&mensa=15#mensaplan",
	},
}

func selectMensa(c *gin.Context, data models.MensaData, day string) {
	url, found := mensaURLs[data.City][data.LocationType]
	if !found {
		fmt.Println("Invalid mensa selection")
		c.JSON(400, gin.H{"error": "Invalid mensa selection"})
		return
	}

	ParsePage(c, url, day)
}

func selectDay(day string) string {
	switch day {
	case "Monday":
		return ".tag_headline:first-child"
	case "Tuesday":
		return ".tag_headline:nth-child(2)"
	case "Wednesday":
		return ".tag_headline:nth-child(3)"
	case "Thursday":
		return ".tag_headline:nth-child(4)"
	case "Friday":
		return ".tag_headline:nth-child(5)"
	default:
		return ".tag_headline:first-child"
	}
}

func selectWeek(current bool) string {
	if current {
		return ".tag_headline:first-child"
	} else {
		return ".tag_headline:nth-child(2)"
	}
}
