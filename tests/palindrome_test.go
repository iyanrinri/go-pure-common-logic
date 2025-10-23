package tests

import (
	"testing"

	"go-common-logic/internal/models"
	"go-common-logic/internal/services"

	"github.com/stretchr/testify/assert"
)

func TestPalindromeService_Check(t *testing.T) {
	service := services.NewPalindromeService()

	tests := []struct {
		name         string
		request      *models.PalindromeRequest
		expected     bool
		expectedText string
	}{
		{
			name: "Simple palindrome: racecar",
			request: &models.PalindromeRequest{
				Text: "racecar",
			},
			expected:     true,
			expectedText: "racecar",
		},
		{
			name: "Case insensitive: RaceCar",
			request: &models.PalindromeRequest{
				Text: "RaceCar",
			},
			expected:     true,
			expectedText: "racecar",
		},
		{
			name: "With spaces: taco cat",
			request: &models.PalindromeRequest{
				Text: "taco cat",
			},
			expected:     true,
			expectedText: "tacocat",
		},
		{
			name: "With punctuation: A man, a plan, a canal: Panama",
			request: &models.PalindromeRequest{
				Text: "A man, a plan, a canal: Panama",
			},
			expected:     true,
			expectedText: "amanaplanacanalpanama",
		},
		{
			name: "Not a palindrome: race a car backwards",
			request: &models.PalindromeRequest{
				Text: "race a car backwards",
			},
			expected:     false,
			expectedText: "raceacarbackwards",
		},
		{
			name: "Single character: a",
			request: &models.PalindromeRequest{
				Text: "a",
			},
			expected:     true,
			expectedText: "a",
		},
		{
			name: "Empty string",
			request: &models.PalindromeRequest{
				Text: "",
			},
			expected:     true,
			expectedText: "",
		},
		{
			name: "Numbers: 12321",
			request: &models.PalindromeRequest{
				Text: "12321",
			},
			expected:     true,
			expectedText: "12321",
		},
		{
			name: "Mixed alphanumeric: Was it a car or a cat I saw?",
			request: &models.PalindromeRequest{
				Text: "Was it a car or a cat I saw?",
			},
			expected:     true,
			expectedText: "wasitacaroracatisaw",
		},
		{
			name: "Not palindrome: hello",
			request: &models.PalindromeRequest{
				Text: "hello",
			},
			expected:     false,
			expectedText: "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.Check(tt.request)

			assert.NotNil(t, result)
			assert.Equal(t, tt.expected, result.IsPalindrome)
			assert.Equal(t, tt.expectedText, result.CleanText)
		})
	}
}

func TestPalindromeService_CheckAdvanced(t *testing.T) {
	service := services.NewPalindromeService()

	tests := []struct {
		name     string
		text     string
		expected bool
	}{
		{
			name:     "Complex palindrome with special characters",
			text:     "Madam, I'm Adam!",
			expected: true,
		},
		{
			name:     "Not a palindrome with numbers",
			text:     "12345",
			expected: false,
		},
		{
			name:     "Unicode palindrome",
			text:     "Able was I ere I saw Elba",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.CheckAdvanced(tt.text)
			assert.Equal(t, tt.expected, result)
		})
	}
}
