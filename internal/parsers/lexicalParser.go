package parsers

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var acceptedChars = map[string]string{
	":": "Colon",
	",": "Comma",
	"{": "BraceOpen",
	"}": "BraceClose",
	"[": "BracketOpen",
	"]": "BracketClose",
}

type Token struct {
	TokenType string
	Token     interface{}
}

func isString(str []string) (bool, string) {
	if len(str) > 1 && str[0] == "\"" && str[len(str)-1] == "\"" {
		return true, strings.Join(str, "")
	}
	return false, ""
}

func isValidToken(str []string) (bool, string, string) {
	if len(str) > 0 {
		for char, tokenType := range acceptedChars {
			if str[0] == char {
				return true, tokenType, char
			}
		}
	}
	return false, "", ""
}

func isWhiteSpace(str []string) bool {
	if len(str) > 0 {
		if str[0] == " " || str[0] == "\n" || str[0] == "\t" {
			return true
		}
	}
	return false
}

func isBoolean(str []string) (bool, string, string) {
	if len(str) == 4 || len(str) == 5 {
		if strings.Join(str, "") == "true" || strings.Join(str, "") == "false" {
			return true, "Boolean", strings.Join(str, "")
		} else if strings.Join(str, "") == "null" {
			return true, "Null", strings.Join(str, "")
		}
	}
	return false, "", ""
}

func isNumber(str []string) (bool, float64) {
	if len(str) > 0 {
		strJoined := strings.Join(str, "")
		num, err := strconv.ParseFloat(strJoined, 64)
		if err != nil {
			return false, 0
		}
		return true, num
	}
	return false, 0
}

func Lexer(data []byte) ([]Token, error) {
	// if file is empty, return empty slice
	if len(data) == 0 {
		return []Token{}, nil
	}

	var tokens []Token

	tempString := make([]string, 0)
	inQuotes := false
	inBoolean := false
	inNumber := false

	// iterate through data and add to tempString for processing
	for i, s := range data {
		tempString = append(tempString, string(s))

		// check if the current character is a valid token character
		if isAcceptedChar, tokenType, tokenResult := isValidToken(tempString); isAcceptedChar {
			tokens = append(tokens, Token{TokenType: tokenType, Token: tokenResult})
			tempString = make([]string, 0)
			continue
		}

		// ignore whitespace not inside quotes
		if isWhiteSpace(tempString) {
			tempString = make([]string, 0)
			continue
		}

		// check if the current character is a quote, and then a valid string
		if tempString[len(tempString)-1] == "\"" {
			inQuotes = !inQuotes
		}
		if inQuotes {
			continue
		}
		if isStr, strResult := isString(tempString); isStr {
			tokens = append(tokens, Token{TokenType: "String", Token: strResult})
			tempString = make([]string, 0)
			continue
		}

		// check if the upcoming value is a boolean or null
		if tempString[0] == "t" || tempString[0] == "f" || tempString[0] == "n" {
			inBoolean = true
		}
		if inBoolean && !unicode.IsLetter(rune(data[i+1])) {
			inBoolean = false
		}
		if inBoolean {
			continue
		}
		if isBool, tokenType, boolResult := isBoolean(tempString); isBool {
			tokens = append(tokens, Token{TokenType: tokenType, Token: boolResult})
			tempString = make([]string, 0)
			continue
		}

		// check if the current value is a number, and look ahead to see if it is a number
		if unicode.IsDigit(rune(s)) || s == '.' || s == '-' || s == '+' || s == 'e' || s == 'E' {
			inNumber = true
		}
		if inNumber && (!unicode.IsDigit(rune(data[i+1])) && data[i+1] != '.' && data[i+1] != '-' && data[i+1] != '+' && data[i+1] != 'e' && data[i+1] != 'E') {
			inNumber = false
		}
		if inNumber {
			continue
		}
		if isNum, numResults := isNumber(tempString); isNum {
			tokens = append(tokens, Token{TokenType: "Number", Token: numResults})
			tempString = make([]string, 0)
			continue
		}

		// else invalid token
		return []Token{}, fmt.Errorf("index %d, invalid token: %v", i, strings.Join(tempString, ""))
	}

	return tokens, nil
}
