package tests

import (
	"testing"

	"go-common-logic/internal/models"
	"go-common-logic/internal/services"

	"github.com/stretchr/testify/assert"
)

func TestTwoSumService_Solve(t *testing.T) {
	service := services.NewTwoSumService()

	tests := []struct {
		name     string
		request  *models.TwoSumRequest
		expected *models.TwoSumResponse
		hasError bool
	}{
		{
			name: "Example 1: [2,7,11,15], target 9",
			request: &models.TwoSumRequest{
				Nums:   []int{2, 7, 11, 15},
				Target: 9,
			},
			expected: &models.TwoSumResponse{
				Indices: []int{0, 1},
			},
			hasError: false,
		},
		{
			name: "Example 2: [3,2,4], target 6",
			request: &models.TwoSumRequest{
				Nums:   []int{3, 2, 4},
				Target: 6,
			},
			expected: &models.TwoSumResponse{
				Indices: []int{1, 2},
			},
			hasError: false,
		},
		{
			name: "Example 3: [3,3], target 6",
			request: &models.TwoSumRequest{
				Nums:   []int{3, 3},
				Target: 6,
			},
			expected: &models.TwoSumResponse{
				Indices: []int{0, 1},
			},
			hasError: false,
		},
		{
			name: "No solution exists",
			request: &models.TwoSumRequest{
				Nums:   []int{1, 2, 3},
				Target: 7,
			},
			expected: nil,
			hasError: true,
		},
		{
			name: "Array too small",
			request: &models.TwoSumRequest{
				Nums:   []int{1},
				Target: 1,
			},
			expected: nil,
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.Solve(tt.request)

			if tt.hasError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.expected.Indices, result.Indices)
			}
		})
	}
}
