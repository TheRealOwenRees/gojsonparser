package parsers

//
//import (
//	"fmt"
//	"strconv"
//	"strings"
//)
//
//var acceptedChars = map[string]string{
//	":": "Colon",
//	",": "Comma",
//	"{": "LeftBrace",
//	"}": "RightBrace",
//	"[": "LeftBracket",
//	"]": "RightBracket",
//}
//
//var numChars = [15]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ".", "-", "+", "e", "E"}
//
//// Token TODO use this struct and pointers, and make methods for it
//type Token struct {
//	TokenType string
//	Token     interface{} // make type with relevant tokens
//}
//
//func isString(str []string) (bool, string) {
//
//	// check if the string is empty
//	if len(str) < 2 {
//		return false, ""
//	}
//
//	// check if the first and last elements are double quotes
//	if str[0] == "\"" && str[len(str)-1] == "\"" {
//		return true, strings.Join(str, "")
//	}
//
//	return false, ""
//}
//
//func isNumber(str []string) (bool, float64) {
//	if len(str) == 0 {
//		return false, 0
//	}
//
//	lastChar := str[len(str)-1]
//	isValidNumberPart := false
//
//	// check if the last character is a valid number character
//	for _, numChar := range numChars {
//		if lastChar == numChar {
//			isValidNumberPart = true
//			break
//		}
//	}
//
//	// if the last character is not a valid number character, parse the string as float64
//	if !isValidNumberPart {
//		// join string except for the last invalid character
//		strJoined := strings.Join(str[:len(str)-1], "")
//
//		// parse string into float64 - this also removes the trailing decimal point
//		num, err := strconv.ParseFloat(strJoined, 64)
//		if err != nil {
//			return false, 0
//		}
//
//		return true, num
//	}
//
//	return false, 0
//}
//
//func isBoolean(str []string) (bool, string) {
//	strJoined := strings.Join(str, "")
//
//	if strJoined == "true" || strJoined == "false" {
//		return true, strJoined
//	}
//
//	// TODO - check if the string will be longer than the accepted values, and return false if so
//
//	return false, ""
//}
//
//func isNull(str []string) (bool, string) {
//	strJoined := strings.Join(str, "")
//
//	if strJoined == "null" {
//		return true, strJoined
//	}
//	// TODO - check if the string will be longer than the accepted values, and return false if so
//
//	return false, ""
//}
//
//func isOtherValidToken(str []string) (bool, string, string) {
//	if len(str) == 0 {
//		return false, "", ""
//	}
//
//	for char, tokenType := range acceptedChars {
//		if str[0] == char {
//			return true, tokenType, str[0]
//		}
//	}
//	return false, "", ""
//}
//
//func isWhiteSpace(str []string) bool {
//	if len(str) == 0 {
//		return false
//	}
//
//	if str[0] == " " || str[0] == "\n" || str[0] == "\t" {
//		return true
//	}
//
//	return false
//}
//
//func Lexer(data []byte) ([]Token, error) {
//	// parsed tokens
//	var tokens []Token
//
//	if len(data) == 0 {
//		return []Token{}, nil
//	}
//
//	// holding string values, to reset after a string or other suitable token is found
//	tempString := make([]string, 0)
//
//	for _, s := range data {
//		tempString = append(tempString, string(s))
//		fmt.Println(strings.Join(tempString, ""))
//
//		if isOtherToken, tokenType, otherTokenResult := isOtherValidToken(tempString); isOtherToken {
//			fmt.Println("Found a valid token")
//			tokens = append(tokens, Token{TokenType: tokenType, Token: otherTokenResult})
//			tempString = make([]string, 0)
//		} else if isWhiteSpace(tempString) {
//			fmt.Println("Found whitespace")
//			tempString = make([]string, 0)
//		} else if isNum, value := isNumber(tempString); isNum {
//			fmt.Println("Found a number")
//			tokens = append(tokens, Token{TokenType: "Number", Token: value})
//			tempString = make([]string, 0)
//		} else if isBool, boolResult := isBoolean(tempString); isBool {
//			fmt.Println("Found a boolean")
//			tokens = append(tokens, Token{TokenType: "Boolean", Token: boolResult})
//			tempString = make([]string, 0)
//		} else if isNullish, nullResult := isNull(tempString); isNullish {
//			fmt.Println("Found null")
//			tokens = append(tokens, Token{TokenType: "Null", Token: nullResult})
//			tempString = make([]string, 0)
//		} else if inProgress, strResult := isString(tempString); inProgress && strResult != "" { // TODO needs to return string and an in-progress status
//			fmt.Println("Found a string")
//			tokens = append(tokens, Token{TokenType: "String", Token: strResult})
//			tempString = make([]string, 0)
//		}
//		//else {
//		//	fmt.Println("Found an invalid token")
//		//	return nil, fmt.Errorf("invalid token: %s", strings.Join(tempString, ""))
//		//}
//	}
//
//	for _, token := range tokens {
//		fmt.Println(token)
//	}
//
//	return tokens, nil
//}
