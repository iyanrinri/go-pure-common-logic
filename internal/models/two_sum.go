package models

type TwoSumRequest struct {
	Nums   []int `json:"nums" binding:"required"`
	Target int   `json:"target" binding:"required"`
}

type TwoSumResponse struct {
	Indices []int `json:"indices"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// PalindromeRequest represents the input data for palindrome checking
type PalindromeRequest struct {
	Text string `json:"text" binding:"required"`
}

// PalindromeResponse represents the output data for palindrome checking
type PalindromeResponse struct {
	IsPalindrome bool   `json:"is_palindrome"`
	CleanText    string `json:"clean_text"`
}
