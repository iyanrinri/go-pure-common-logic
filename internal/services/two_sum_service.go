package services

import (
	"errors"
	"go-common-logic/internal/models"
)

type TwoSumService struct{}

func NewTwoSumService() *TwoSumService {
	return &TwoSumService{}
}

func (s *TwoSumService) Solve(request *models.TwoSumRequest) (*models.TwoSumResponse, error) {
	if len(request.Nums) < 2 {
		return nil, errors.New("array must contain at least 2 numbers")
	}

	numMap := make(map[int]int)

	for i, num := range request.Nums {
		complement := request.Target - num

		if index, exists := numMap[complement]; exists {
			return &models.TwoSumResponse{
				Indices: []int{index, i},
			}, nil
		}

		numMap[num] = i
	}

	return nil, errors.New("no solution found")
}
