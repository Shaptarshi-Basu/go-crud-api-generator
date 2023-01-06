package util

import "strings"

func toCamelCase(s string) string {
	words := strings.Split(s, " ")
	for i, word := range words {
		if i > 0 {
			word = strings.Title(word)
		}
		words[i] = word
	}
	return strings.Join(words, "")
}

func CovertPathtoCamelCaseMethodName(path string, method string) string {
	return method + toCamelCase(strings.ReplaceAll(path, "/", " "))
}
