package test

import (
	"bytes"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"net/http"
	"testing"
)

type user struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type signInResp struct {
	Token string `json:"token"`
}

func TestSignUp(t *testing.T) {
	baseURL := "http://localhost:8080"
	signUpURL := "/auth/sign-up"

	users := []user{
		{"John Marston", "JohnRDR1", "qazwsxedc"},
		{"Dominic Torretto", "TheLastRide", "pormifamilia"},
		{"James Bond", "JamesBond007", "mi6mi6mi6"},
		{"Arthur Morgan", "SweetyBoy", "Pinketrons"},
		{"Nikola Jokic", "Joker", "Denver5124"},
	}

	for _, u := range users {
		u.Username = "test_" + u.Username
		body, _ := json.Marshal(u)
		req, err := http.NewRequest("POST", baseURL+signUpURL, bytes.NewBuffer(body))
		if err != nil {
			t.Fatalf("Error in [SignUp cycle]: %v", err)
		}

		req.Header.Set("Origin", "http://localhost:3000")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Error in [SignUp cycle]: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Error in [SignUp cycle]: %v", resp.Status)
		}

		resp.Body.Close()
	}

	t.Logf("All users registered successfully.")
}

func TestSignIn(t *testing.T) {
	baseURL := "http://localhost:8080"
	signInURL := "/auth/sign-in"

	users := []user{
		{"John Marston", "JohnRDR1", "qazwsxedc"},
		{"Dominic Torretto", "TheLastRide", "pormifamilia"},
		{"James Bond", "JamesBond007", "mi6mi6mi6"},
		{"Arthur Morgan", "SweetyBoy", "Pinketrons"},
		{"Nikola Jokic", "Joker", "Denver5124"},
	}

	var tokens []string

	for _, u := range users {
		loginBody := map[string]string{
			"username": "test_" + u.Username,
			"password": u.Password,
		}
		body, _ := json.Marshal(loginBody)
		req, err := http.NewRequest("POST", baseURL+signInURL, bytes.NewBuffer(body))
		if err != nil {
			t.Fatalf("Error in [SignIn cycle]: %v", err)
		}

		req.Header.Set("Origin", "http://localhost:3000")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Error in [SignIn cycle]: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf(". Current user %s. Error in [SignIn cycle]: %v", u.Username, resp.Status)
		}

		var r signInResp
		if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
			t.Fatalf("Failed to decode response body: %v", err)
		}
		resp.Body.Close()

		if r.Token == "" {
			t.Fatalf("Token is empty for user %s", u.Username)
		}

		tokens = append(tokens, r.Token)
	}

	tokenSet := make(map[string]struct{})
	for _, token := range tokens {
		if _, exists := tokenSet[token]; exists {
			t.Fatalf("Duplicate token found: %s", token)
		}
		tokenSet[token] = struct{}{}
	}

	t.Logf("All tokens are unique")
}

func rollbackUsersTable(t *testing.T) {
	connDB := "host=localhost port=5432 user=postgres password=qwerty dbname=postgres sslmode=disable"
	dbname := "postgres"
	db, err := sqlx.Open(dbname, connDB)
	if err != nil {
		t.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(`DELETE FROM users WHERE username LIKE 'test_%'`)
	if err != nil {
		t.Fatalf("Error deleting rows: %v", err)
	}
}
