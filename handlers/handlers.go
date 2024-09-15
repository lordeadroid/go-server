package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func getRandomNumber() int {
	max := 8
	min := 0
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	return r.Intn(max-min+1) + min
}

func nextPos(state []string) int {
	pos := getRandomNumber()

	if state[pos] == "" {
		return pos
	}

	return nextPos(state)

}

func GetNextPosition(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(req.Body)
	defer req.Body.Close()

	var state []string

	err = json.Unmarshal(body, &state)
	pos := nextPos(state)

	if err != nil {
		http.Error(res, "Error reading request body", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(res, "%d", pos)
}
