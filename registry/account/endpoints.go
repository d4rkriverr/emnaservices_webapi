package account

import (
	"emnaservices/webapi/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Handler struct {
	service *Service
}

func newHandler(s *Service) *Handler {
	return &Handler{service: s}
}

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	var creds UserCredentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil || creds.Username == "" || creds.Password == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	userID, err := h.service.Authenticate(creds.Username, creds.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Store the access token in the database
	accessToken, err := h.service.StoreAccessToken(userID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to store access token")
		return
	}

	utils.RespondWithSuccess(w, accessToken)
}

func (handler *Handler) HandleUserCreate(w http.ResponseWriter, r *http.Request) {
	var creds UserCredentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil || creds.Username == "" || creds.Password == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err = handler.service.UserCreate(creds.Username, creds.Password)
	if err != nil {
		fmt.Println(err)
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid credentials")
		return
	}
	utils.RespondWithSuccess(w, "created successfully")
}

func (h *Handler) HandleUserInfo(w http.ResponseWriter, r *http.Request) {
	token, err := GetAuthorizationToken(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid Authorization ErrNoRows")
		return
	}
	user, err := h.service.ValidToken(token)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid Authorization ErrQuery")
		return
	}
	utils.RespondWithSuccess(w, map[string]any{"username": user.Username, "role": user.Role})
}

// GET TOKEN FROM REQUEST
func GetAuthorizationToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("Missing Authorization header")
	}
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return "", fmt.Errorf("Invalid Authorization format")
	}
	return tokenParts[1], nil
}
