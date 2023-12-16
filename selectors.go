package main

var mensaURLs = []string{
	"https://studentenwerk.sh/de/essen-uebersicht?ort=2&mensa=7#mensaplan",
	"https://studentenwerk.sh/de/essen-uebersicht?ort=2&mensa=14#mensaplan",
	"https://studentenwerk.sh/de/essen-uebersicht?ort=4&mensa=14#mensaplan",
	"https://studentenwerk.sh/de/essen-uebersicht?ort=1&mensa=1#mensaplan",
	"https://studentenwerk.sh/de/essen-uebersicht?ort=1&mensa=2#mensaplan",
	"https://studentenwerk.sh/de/essen-uebersicht?ort=1&mensa=3#mensaplan",
	"https://studentenwerk.sh/de/essen-uebersicht?ort=1&mensa=4#mensaplan",
	"https://studentenwerk.sh/de/essen-uebersicht?ort=1&mensa=5#mensaplan",
	"https://studentenwerk.sh/de/essen-uebersicht?ort=1&mensa=6#mensaplan",
	"https://studentenwerk.sh/de/essen-uebersicht?ort=3&mensa=8#mensaplan",
	"https://studentenwerk.sh/de/essen-uebersicht?ort=3&mensa=9#mensaplan",
	"https://studentenwerk.sh/de/essen-uebersicht?ort=5&mensa=15#mensaplan",
}

func selectMensa(canteen int) string {
	url := mensaURLs[canteen]
	return url
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
