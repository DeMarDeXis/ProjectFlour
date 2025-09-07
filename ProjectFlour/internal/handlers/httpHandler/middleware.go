package httpHandler

import (
	"ProjectFlour/internal/handlers/httpHandler/httplib"
	"context"
	"errors"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"
	userCtx    = "userID"
)

func (h *HTTPHandler) userIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authHeader)
		if header == "" {
			httplib.NewErrorResponse(w, h.logg, http.StatusUnauthorized, "empty auth header")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			httplib.NewErrorResponse(w, h.logg, http.StatusUnauthorized, "invalid auth header")
			return
		}

		userID, err := h.service.AuthService.ParseToken(headerParts[1])
		if err != nil {
			httplib.NewErrorResponse(w, h.logg, http.StatusUnauthorized, err.Error())
			return
		}

		//r.Header.Set(userCtx, strconv.Itoa(userID))
		ctx := context.WithValue(r.Context(), userCtx, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *HTTPHandler) getUserId(r *http.Request) (int, error) {
	id, ok := r.Context().Value(userCtx).(int)
	if !ok {
		return 0, errors.New("user id not found")
	}

	return id, nil
}
