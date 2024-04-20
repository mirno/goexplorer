package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mirno/goexplorer/internal/domain/stack"
)

var globalStack = stack.Stack{}

func PushHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Push request received")
	w.Header().Set("Content-Type", "application/json")
	var item struct {
		Value interface{} `json:"value"`

	}
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil{
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	globalStack.Push(item.Value)
	w.WriteHeader(http.StatusNoContent) // 204 No Content for successful PUT without response body
}

func PopHandler(w http.ResponseWriter, r*http.Request) {
	log.Println("Pop request received")
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
	}

	item, ok := globalStack.Pop()
	if !ok {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"value": item})
}

func PeekHandler(w http.ResponseWriter, r*http.Request) {
	log.Println("Peek request received")
    item := globalStack.Peek()
    if item == nil {
        w.WriteHeader(http.StatusNoContent) // 204 No content if stack is empty.
        return // Ensure to return immediately to avoid writing to the body.
    }
	log.Printf("Peeking item: %+v\n", item)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{"value": item})
}