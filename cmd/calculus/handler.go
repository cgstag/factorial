package calculus

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
)

type FactorialResult struct {
	N uint64 `json:"n"`
	Result *big.Int `json:"result"`
}

func NewCalculusHandler(w http.ResponseWriter, r *http.Request) {
	ns := r.URL.Query().Get("n")
	n, err := strconv.ParseUint(ns, 10, 64)
	if err != nil {
		message, _ := json.Marshal("Parameter N should be an unsigned integer")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(message)
		return
	}
	res := Factorial(big.NewInt(int64(n)))
	response, err := json.Marshal(FactorialResult{N:n,Result:res})
	if err != nil {
		message, _ := json.Marshal(fmt.Sprintf("Error calculating the factorial of %d",n))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(message)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}