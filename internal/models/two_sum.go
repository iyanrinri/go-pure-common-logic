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
