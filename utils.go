package main

import "strings"

func removeLinks(text string) string {
	words := strings.Split(text, " ")
	var cleaned []string
	for i := 0; i < len(words); i++ {
		if strings.Contains(words[i], "http") {

		} else {
			cleaned = append(cleaned, words[i])
		}
	}
	return strings.Join(cleaned, " ")
}
