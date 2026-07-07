package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Aadi022/Backend_Golang/internal/dto"
	"github.com/Aadi022/Backend_Golang/internal/response"
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
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	err := h.service.Register( //we register in db using the service
		req.Name,
		req.Email,
		req.Password,
	)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(w, http.StatusCreated, map[string]string{
		"message": "User registered successfully",
	})
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest //Holds the json request body

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { // Parse JSON into the struct
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	token, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err.Error())
		return
	}

	response.Success(w, http.StatusOK, map[string]string{ //respond with the jwt token
		"access_token": token,
	})
}
