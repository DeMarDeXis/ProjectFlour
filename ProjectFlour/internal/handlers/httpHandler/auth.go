package httpHandler

import (
	"ProjectFlour/internal/handlers/httpHandler/httplib"
	"ProjectFlour/internal/model"
	"encoding/json"
	"net/http"
)

// @Summary SignUp
// @Description create profile for new user/client
// @Tags auth
// @Accept json
// @Produce json
// @Param input body model.User true "credentials"
// @Success 200 {object} map[string]interface{}
// @Router /auth/sign-up [post]
func (h *HTTPHandler) signUp(w http.ResponseWriter, r *http.Request) {
	var input model.User

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		httplib.NewErrorResponse(w, h.logg, http.StatusBadRequest, err.Error())
		return
	}

	if err := input.Validate(); err != nil {
		httplib.NewErrorResponse(w, h.logg, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.AuthService.CreateUser(input)
	if err != nil {
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{"id": id}
	json.NewEncoder(w).Encode(response)
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary SignIn
// @Description get token for user/client
// @Tags auth
// @Accept json
// @Produce json
// @Param input body signInInput true "credentials"
// @Success 200 {object} map[string]interface{}
// @Router /auth/sign-in [post]
func (h *HTTPHandler) signIn(w http.ResponseWriter, r *http.Request) {
	var input signInInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		httplib.NewErrorResponse(w, h.logg, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.service.AuthService.GenerateToken(input.Username, input.Password)
	if err != nil {
		httplib.NewErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{"token": token}
	json.NewEncoder(w).Encode(response)
}
