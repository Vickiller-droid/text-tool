package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
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
	fmt.Println("Done! Check", output)
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
				result[j] =binToDecimal(result[j])
			}
		}
	}
	return result
}

func capitalize word(string) string {
	if len(words) == 0 {
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
