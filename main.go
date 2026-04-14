package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: go run . input.txt output.txt")
		return
	}

	input := os.Args[1]
	output := os.Args[2]

	text, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result := processText(string(text))

	err = os.WriteFile(output, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}

func processText(text string) string {
	words := strings.Fields(text)
	words = applyTransformations(words)
	result := strings.Join(words, " ")
	result = fixPunctuation(result)
	result = fixApostrophes(result)
	result = fixArticles(result)
	return result
}

func applyTransformations(tokens []string) []string {
	result := make([]string, 0)

	for i := 0; i < len(tokens); i++ {
		tagName, count, isTag := parseTag(tokens[i])

		if !isTag {
			result = append(result, tokens[i])
			continue
		}

		if count > len(result) {
			count = len(result)
		}

		for j := len(result) - count; j < len(result); j++ {
			switch tagName {
			case "up":
				result[j] = strings.ToUpper(result[j])
			case "low":
				result[j] = strings.ToLower(result[j])
			case "cap":
				result[j] = capitalize(result[j])
			case "hex":
				result[j] = hexToDecimal(result[j])
			case "bin":
				result[j] = binToDecimal(result[j])
			}
		}
	}
	return result
}

func parseTag(token string) (string, int, bool) {
	t := strings.ToLower(strings.ReplaceAll(token, " ", ""))

	switch t {
	case "(up)":
		return "up", 1, true
	case "(low)":
		return "low", 1, true
	case "(cap)":
		return "cap", 1, true
	case "(hex)":
		return "hex", 1, true
	case "(bin)":
		return "bin", 1, true
	}

	if strings.HasPrefix(t, "(") && strings.HasSuffix(t, ")") {
		inner := t[1 : len(t)-1]
		parts := strings.SplitN(inner, ",", 2)
		if len(parts) == 2 {
			name := strings.TrimSpace(parts[0])
			numStr := strings.TrimSpace(parts[1])
			num, err := strconv.Atoi(numStr)
			if err == nil && (name == "up" || name == "low" || name == "cap") {
				return name, num, true
			}
		}
	}
	return "", 0, false
}

func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(word[0:1]) + strings.ToLower(word[1:])
}

func hexToDecimal(word string) string {
	num, err := strconv.ParseInt(word, 16, 64)
	if err != nil {
		return word
	}
	return strconv.Itoa(int(num))
}

func binToDecimal(word string) string {
	num, err := strconv.ParseInt(word, 2, 64)
	if err != nil {
		return word
	}
	return strconv.Itoa(int(num))
}

func fixPunctuation(text string) string {
	puncts := []string{".", ",", "!", "?", ":", ";"}
	for _, p := range puncts {
		for strings.Contains(text, " "+p) {
			text = strings.ReplaceAll(text, " "+p, p)
		}
	}

	var builder strings.Builder
	runes := []rune(text)

	for i := 0; i < len(runes); i++ {
		builder.WriteRune(runes[i])
		if isPunctuation(runes[i]) && i+1 < len(runes) {
			next := runes[i+1]
			if next != ' ' && next != '\n' && !isPunctuation(next) {
				builder.WriteRune(' ')
			}
		}
	}
	return builder.String()
}

func isPunctuation(r rune) bool {
	return r == '.' || r == ',' || r == '!' || r == '?' || r == ':' || r == ';'
}

func fixApostrophes(text string) string {
	parts := strings.Split(text, "'")
	if len(parts) < 3 {
		return text
	}

	var builder strings.Builder
	for i, part := range parts {
		if i == 0 {
			builder.WriteString(strings.TrimRight(part, " "))
		} else if i%2 == 1 {
			builder.WriteRune('\'')
			builder.WriteString(strings.TrimSpace(part))
		} else {
			builder.WriteRune('\'')
			builder.WriteString(strings.TrimLeft(part, " "))
		}
	}
	return builder.String()
}

func fixArticles(text string) string {
	words := strings.Fields(text)
	vowelsAndH := "aeiouAEIOUhH"

	for i := 0; i < len(words)-1; i++ {
		if strings.ToLower(words[i]) == "a" {
			next := words[i+1]
			if len(next) > 0 && strings.ContainsRune(vowelsAndH, rune(next[0])) {
				if words[i] == "A" {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			}
		}
	}
	return strings.Join(words, " ")
}
