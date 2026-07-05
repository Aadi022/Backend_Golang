package handler

import (
	"encoding/json"
	"net/http"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler { //Constructor and returns the struct HealthHandler. We are using pointer here as we don't want to copy struct causing more memory usage
	return &HealthHandler{} //while creating a healthhandler, return the pointer of it(memory address), so caller gets a reference of the same object not copy
}

func (h *HealthHandler) Health( //method for HealthHandler, pointer here as it will point to the original copy of the struct
	w http.ResponseWriter, //writing to respinse to the client, no pointer because this is an interface
	r *http.Request, //reading the client's request, pointer as its a struct
) {
	w.Header().Set("Content-Type", "application/json") //setting the header of the response
	//creates a json encoder, encodes directly to response
	json.NewEncoder(w).Encode(map[string]string{ //creating a map with key string and value string
		"status": "ok",
	})
}

func (h *HealthHandler) Ready(
	w http.ResponseWriter,
	r *http.Request,
) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"status": "ready",
	})
}
