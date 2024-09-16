package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func getRandomNumber(max int) int {
	min := 0
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	return r.Intn(max-min+1) + min
}

func getAvailabePositions(state []string) []int {
	var indices []int
	for index, element := range state {
		if element == "" {
			indices = append(indices, index)
		}
	}
	return indices
}

func nextPos(state []string) int {
	availabePositions := getAvailabePositions(state)
	pos := getRandomNumber(len(availabePositions) - 1)

	return availabePositions[pos]
}

func GetNextPosition(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if req.Method == http.MethodOptions {
		res.WriteHeader(http.StatusOK)
		return
	}

	if req.Method != http.MethodPost {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	var state []string
	err = json.Unmarshal(body, &state)
	if err != nil {
		http.Error(res, "Error parsing request body", http.StatusBadRequest)
		return
	}

	pos := nextPos(state)
	fmt.Fprintf(res, "%d", pos)
}
