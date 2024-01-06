package parsers

import (
	"fmt"
	"gojsonparser/internal/parsers"
	"gojsonparser/tools/fileHandlers"
	"reflect"
	"testing"
)

func TestLexer(t *testing.T) {
	tests := []struct {
		name          string
		input         []byte
		expected      []parsers.Token
		expectedError error
	}{
		{
			name:     "Step1 Invalid - Empty Input",
			input:    fileHandlers.ReadFile("test-data/step1/invalid.json"),
			expected: []parsers.Token{},
		},
		{
			name:     "Step1 Valid - Valid Empty JSON",
			input:    fileHandlers.ReadFile("test-data/step1/valid.json"),
			expected: []parsers.Token{{TokenType: "BraceOpen", Token: "{"}, {TokenType: "BraceClose", Token: "}"}},
		},
		{
			name:  "Step2 Invalid - Trailing Comma",
			input: fileHandlers.ReadFile("test-data/step2/invalid.json"),
			expected: []parsers.Token{
				{TokenType: "BraceOpen", Token: "{"},
				{TokenType: "String", Token: "\"key\""},
				{TokenType: "Colon", Token: ":"},
				{TokenType: "String", Token: "\"value\""},
				{TokenType: "Comma", Token: ","},
				{TokenType: "BraceClose", Token: "}"},
			},
		},
		{
			name:          "Step2 Invalid2 - Key without double quotes",
			input:         fileHandlers.ReadFile("test-data/step2/invalid2.json"),
			expected:      []parsers.Token{},
			expectedError: fmt.Errorf("index 22, invalid token: k"),
		},
		{
			name:  "Step2 Valid - 1 key:value pair",
			input: fileHandlers.ReadFile("test-data/step2/valid.json"),
			expected: []parsers.Token{
				{TokenType: "BraceOpen", Token: "{"},
				{TokenType: "String", Token: "\"key\""},
				{TokenType: "Colon", Token: ":"},
				{TokenType: "String", Token: "\"value\""},
				{TokenType: "BraceClose", Token: "}"},
			},
		},
		{
			name:  "Step2 Valid - 2 key:value pairs",
			input: fileHandlers.ReadFile("test-data/step2/valid2.json"),
			expected: []parsers.Token{
				{TokenType: "BraceOpen", Token: "{"},
				{TokenType: "String", Token: "\"key\""},
				{TokenType: "Colon", Token: ":"},
				{TokenType: "String", Token: "\"value\""},
				{TokenType: "Comma", Token: ","},
				{TokenType: "String", Token: "\"key2\""},
				{TokenType: "Colon", Token: ":"},
				{TokenType: "String", Token: "\"value\""},
				{TokenType: "BraceClose", Token: "}"},
			},
		},
		{
			name:  "Step3 Valid - Boolean and Number",
			input: fileHandlers.ReadFile("test-data/step3/valid.json"),
			expected: []parsers.Token{
				{TokenType: "BraceOpen", Token: "{"},
				{TokenType: "String", Token: "\"key1\""},
				{TokenType: "Colon", Token: ":"},
				{TokenType: "Boolean", Token: "true"},
				{TokenType: "Comma", Token: ","},
				{TokenType: "String", Token: "\"key2\""},
				{TokenType: "Colon", Token: ":"},
				{TokenType: "Boolean", Token: "false"},
				{TokenType: "Comma", Token: ","},
				{TokenType: "String", Token: "\"key3\""},
				{TokenType: "Colon", Token: ":"},
				{TokenType: "Null", Token: "null"},
				{TokenType: "Comma", Token: ","},
				{TokenType: "String", Token: "\"key4\""},
				{TokenType: "Colon", Token: ":"},
				{TokenType: "String", Token: "\"value\""},
				{TokenType: "Comma", Token: ","},
				{TokenType: "String", Token: "\"key5\""},
				{TokenType: "Colon", Token: ":"},
				{TokenType: "Number", Token: float64(101)},
				{TokenType: "BraceClose", Token: "}"},
			},
		},
		{
			name:          "Step3 Invalid - Invalid Boolean",
			input:         fileHandlers.ReadFile("test-data/step3/invalid.json"),
			expected:      []parsers.Token{},
			expectedError: fmt.Errorf("index 28, invalid token: F"),
		},
		{
			name:  "Step4 Valid - Empty Nested Objects",
			input: fileHandlers.ReadFile("test-data/step4/valid.json"),
			expected: []parsers.Token{
				{TokenType: "BraceOpen", Token: "{"},
				{TokenType: "String", Token: "\"key\""},
				{TokenType: "Colon", Token: ":"},
				{TokenType: "String", Token: "\"value\""},
				{TokenType: "Comma", Token: ","},
				{TokenType: "String", Token: "\"key-n\""},
				{TokenType: "Colon", Token: ":"},
				{TokenType: "Number", Token: float64(101)},
				{TokenType: "Comma", Token: ","},
				{TokenType: "String", Token: "\"key-o\""},
				{TokenType: "Colon", Token: ":"},
				{TokenType: "BraceOpen", Token: "{"},
				{TokenType: "BraceClose", Token: "}"},
				{TokenType: "Comma", Token: ","},
				{TokenType: "String", Token: "\"key-l\""},
				{TokenType: "Colon", Token: ":"},
				{TokenType: "BracketOpen", Token: "["},
				{TokenType: "BracketClose", Token: "]"},
				{TokenType: "BraceClose", Token: "}"},
			},
		},
		{
			name:          "Step4 Invalid - Invalid Nested Object",
			input:         fileHandlers.ReadFile("test-data/step4/invalid.json"),
			expected:      []parsers.Token{},
			expectedError: fmt.Errorf("index 97, invalid token: '"),
		},
		{
			name:  "Step4 Valid2 - Nested Objects",
			input: fileHandlers.ReadFile("test-data/step4/valid2.json"),
			expected: []parsers.Token{
				{TokenType: "BraceOpen", Token: "{"},
				{TokenType: "String", Token: "\"key\""},
				{TokenType: "Colon", Token: ":"},
				{TokenType: "String", Token: "\"value\""},
				{TokenType: "Comma", Token: ","},
				{TokenType: "String", Token: "\"key-n\""},
				{TokenType: "Colon", Token: ":"},
				{TokenType: "Number", Token: float64(101)},
				{TokenType: "Comma", Token: ","},
				{TokenType: "String", Token: "\"key-o\""},
				{TokenType: "Colon", Token: ":"},
				{TokenType: "BraceOpen", Token: "{"},
				{TokenType: "String", Token: "\"inner key\""},
				{TokenType: "Colon", Token: ":"},
				{TokenType: "String", Token: "\"inner value\""},
				{TokenType: "BraceClose", Token: "}"},
				{TokenType: "Comma", Token: ","},
				{TokenType: "String", Token: "\"key-l\""},
				{TokenType: "Colon", Token: ":"},
				{TokenType: "BracketOpen", Token: "["},
				{TokenType: "String", Token: "\"list value\""},
				{TokenType: "BracketClose", Token: "]"},
				{TokenType: "BraceClose", Token: "}"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parsers.Lexer(tt.input)

			if err != nil {
				if tt.expectedError == nil || err.Error() != tt.expectedError.Error() {
					t.Errorf("expected error %v, got %v", tt.expectedError, err)
				}
				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
