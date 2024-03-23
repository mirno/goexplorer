package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mirno/goexplorer/internal/domain/stack"
)

var Stack stack.Stack

func PushItem(w http.ResponseWriter, r *http.Request) {
	var item struct {
		Value interface{} `json:"value"`
	}

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Stack.Push(item.Value)
	w.WriteHeader(http.StatusCreated)
}