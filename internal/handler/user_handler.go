package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Aadi022/Backend_Golang/internal/dto"
	"github.com/Aadi022/Backend_Golang/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) { //HTTP handler for user registration
	var req dto.RegisterRequest                                  //Create a DTO to hold incoming JSON request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { //Read JSON from request body and store it in req
		http.Error(w, "Invalid request Body", http.StatusBadGateway) //Return 400 if JSON is invalid
		return                                                       //Stop execution
	}

	err := h.service.Register( //we register in db using the service
		req.Name,
		req.Email,
		req.Password,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated) //sets the http status code to 201 created before sending the response body

	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
}
