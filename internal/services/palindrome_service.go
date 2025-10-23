package services

import (
	"regexp"
	"strings"
	"unicode"

	"go-common-logic/internal/models"
)

type PalindromeService struct{}

func NewPalindromeService() *PalindromeService {
	return &PalindromeService{}
}

func (s *PalindromeService) Check(request *models.PalindromeRequest) *models.PalindromeResponse {
	cleanText := s.cleanString(request.Text)
	isPalindrome := s.isPalindrome(cleanText)

	return &models.PalindromeResponse{
		IsPalindrome: isPalindrome,
		CleanText:    cleanText,
	}
}

func (s *PalindromeService) cleanString(text string) string {
	var result strings.Builder

	for _, r := range text {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result.WriteRune(unicode.ToLower(r))
		}
	}

	return result.String()
}

func (s *PalindromeService) isPalindrome(text string) bool {
	if len(text) <= 1 {
		return true
	}

	left := 0
	right := len(text) - 1

	for left < right {
		if text[left] != text[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func (s *PalindromeService) CheckAdvanced(text string) bool {
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	cleaned := reg.ReplaceAllString(text, "")
	cleaned = strings.ToLower(cleaned)

	return s.isPalindrome(cleaned)
}
