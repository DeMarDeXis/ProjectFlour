package test

import (
	"github.com/rs/cors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestCorsAlwaysSuccess always success
func TestCorsAlwaysSuccess(t *testing.T) {
	c := setupCors()
	handler := c.Handler(http.HandlerFunc(mockHandler))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Origin", "http://localhost:3000")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode) // Ожидаем 200
	assert.Equal(t, "http://localhost:3000", res.Header.Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "true", res.Header.Get("Access-Control-Allow-Credentials"))
}

// TestCorsSuccessWithPost, which have POST instead of GET
func TestCorsSuccessWithPost(t *testing.T) {
	c := setupCors()
	handler := c.Handler(http.HandlerFunc(mockHandler))

	req := httptest.NewRequest(http.MethodPost, "/test", nil)
	req.Header.Set("Origin", "http://localhost:3000")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "http://localhost:3000", res.Header.Get("Access-Control-Allow-Origin"))
}

// TestCorsFailureWithInvalidOrigin, which have invalid origin
func TestCorsFailureWithInvalidOrigin(t *testing.T) {
	notOriginURL := "http://malicious-site.com"

	c := setupCors()
	handler := c.Handler(http.HandlerFunc(mockHandler))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Origin", notOriginURL)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Empty(t, res.Header.Get("Access-Control-Allow-Origin"))
}

// TestCorsPreflightFailureWithInvalidOrigin is new test with malicious site
func TestCorsPreflightFailureWithInvalidOrigin(t *testing.T) {
	c := setupCors()
	handler := c.Handler(http.HandlerFunc(mockHandler))

	req := httptest.NewRequest(http.MethodOptions, "/test", nil)
	req.Header.Set("Origin", "http://malicious-site.com")
	req.Header.Set("Access-Control-Request-Method", "POST")
	req.Header.Set("Access-Control-Request-Headers", "Content-Type, Authorization")

	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusNoContent, res.StatusCode)
	assert.Empty(t, res.Header.Get("Access-Control-Allow-Origin"))
}

// TestCorsFailureWithInvalidMethod, which have not allowed method
func TestCorsFailureWithInvalidMethod(t *testing.T) {
	c := setupCors()
	handler := c.Handler(http.HandlerFunc(mockHandler))

	req := httptest.NewRequest(http.MethodDelete, "/test", nil)
	req.Header.Set("Origin", "http://localhost:3000")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode)
	assert.Empty(t, res.Header.Get("Access-Control-Allow-Origin"))
}

func setupCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
}

func mockHandler(w http.ResponseWriter, r *http.Request) {
	allowedMethods := []string{"GET", "POST", "OPTIONS"}
	methodAllowed := false
	for _, m := range allowedMethods {
		if m == r.Method {
			methodAllowed = true
			break
		}
	}

	if !methodAllowed {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
