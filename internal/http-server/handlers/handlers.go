package handlers

import (
	"encoding/json"
	"factorial_calculate/internal/lib/algorithms"
	"factorial_calculate/internal/lib/pr"
	"factorial_calculate/internal/lib/rp"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sync"
)

func FactorialHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	reqBody, err := pr.ParseRequest(r)
	fmt.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(rp.Error("Failed to get body"))
		return
	}

	var wg sync.WaitGroup
	var factA, factB int

	wg.Add(1)

	go func() {
		defer wg.Done()
		factA = algorithms.FactorialRecursive(reqBody.A)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		factB = algorithms.FactorialRecursive(reqBody.B)
	}()

	wg.Wait()

	response := map[string]int{"a!": factA, "b!": factB}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return
}
