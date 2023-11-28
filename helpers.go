package main

import (
	"github.com/gocolly/colly"
	"regexp"
	"strings"
)

func extractMealName(e *colly.HTMLElement) string {
	mealName, _ := e.DOM.Find(".menu_name").Html()
	mealName = excludeMensaZusatz(mealName)
	return mealName
}

func excludeMensaZusatz(input string) string {
	// Remove content within parentheses (mensa_zusatz)
	re := regexp.MustCompile(`\s*\([^)]*\)\s*`)
	cleaned := re.ReplaceAllString(input, "")

	// Remove <span> tags
	cleaned = strings.ReplaceAll(cleaned, "<span class=\"mensa_zusatz\">", "")
	cleaned = strings.ReplaceAll(cleaned, "</span>", "")

	// Replace <br/> with a comma and space
	cleaned = strings.ReplaceAll(cleaned, "<br/>", ", ")
	cleaned = strings.ReplaceAll(cleaned, ", ", ",")
	cleaned = strings.ReplaceAll(cleaned, " ,", ", ")
	cleaned = strings.ReplaceAll(cleaned, ",", ", ")

	// Trim spaces around commas
	cleaned = strings.TrimSpace(cleaned)

	return cleaned
}

func dataArtenContains(e *colly.HTMLElement, allergenCode string) bool {
	return strings.Contains(e.Attr("data-arten"), allergenCode)
}
