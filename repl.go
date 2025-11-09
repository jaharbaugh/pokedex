package main

import(
	"strings"
)

func cleanInput(text string) []string {
	cleanText := strings.ToLower(text)
	cleanText = strings.TrimSpace(cleanText)
	return strings.Fields(cleanText)
}