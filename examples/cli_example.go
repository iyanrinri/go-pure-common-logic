package main

import (
	"encoding/json"
	"fmt"
	"go-common-logic/internal/models"
	"go-common-logic/internal/services"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run examples/cli_example.go '<json_input>'")
	}

	input := os.Args[1]

	var request models.TwoSumRequest
	if err := json.Unmarshal([]byte(input), &request); err != nil {
		log.Fatal("Invalid JSON input:", err)
	}

	service := services.NewTwoSumService()
	result, err := service.Solve(&request)

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	output, _ := json.Marshal(result)
	fmt.Println(string(output))
}
