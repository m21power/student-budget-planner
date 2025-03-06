package data

import (
	"encoding/json"
	"net/http"
	"strconv"
	"student-planner/usecases"
	"student-planner/util"
)

type UserHandler struct {
	usecase usecases.UserUsecase
}

func NewUserHandler(usecase usecases.UserUsecase) *UserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	userID, err := strconv.Atoi(id)
	if err != nil {
		util.WriteError(w, err)
		return
	}

	user, err := h.usecase.GetUser(userID)
	if err != nil {
		util.WriteError(w, err)
		return
	}

	util.WriteJSON(w, http.StatusOK, user)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	type LoginPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var payload LoginPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		util.WriteError(w, err)
		return
	}

	user, err := h.usecase.Login(payload.Email, payload.Password)
	if err != nil {
		util.WriteError(w, err)
		return
	}

	util.WriteJSON(w, http.StatusOK, user)
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	type RegisterPayload struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var payload RegisterPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		util.WriteError(w, err)
		return
	}
	err = h.usecase.Register(payload.Name, payload.Email, payload.Password)
	if err != nil {
		util.WriteError(w, err)
	}

	util.WriteJSON(w, http.StatusOK, payload)
}

func (h *UserHandler) UpdateBadge(w http.ResponseWriter, r *http.Request) {
	type UpdateBadgePayload struct {
		ID    int    `json:"id"`
		Badge string `json:"badge"`
	}
	var payload UpdateBadgePayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		util.WriteError(w, err)
		return
	}

	err = h.usecase.UpdateBadge(payload.ID, payload.Badge)
	if err != nil {
		util.WriteError(w, err)
		return
	}

	util.WriteJSON(w, http.StatusOK, payload)
}

func (h *UserHandler) AskGemini(w http.ResponseWriter, r *http.Request) {
	type AskGeminiPayload struct {
		Message string `json:"message"`
	}
	var payload AskGeminiPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		util.WriteError(w, err)
		return
	}

	response, err := h.usecase.AskGemini(payload.Message)
	if err != nil {
		util.WriteError(w, err)
		return
	}

	util.WriteJSON(w, http.StatusOK, response)
}
