package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"goProject/middleware"

	"github.com/dgrijalva/jwt-go"
)

func TestAuthMiddleware_ValidToken(t *testing.T) {
	// Set up a request with a valid token
	tokenString, err := generateValidToken()
	if err != nil {
		t.Fatalf("Error generating valid token: %v", err)
	}

	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", tokenString)

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Create an HTTP handler that uses the AuthMiddleware
	handler := middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// Serve the HTTP request
	handler.ServeHTTP(rr, req)

	// Check the status code
	if rr.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %v, but got %v", http.StatusUnauthorized, rr.Code)
	}
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	// Set up a request with an invalid token
	invalidTokenString := "invalid_token"

	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", invalidTokenString)

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Create an HTTP handler that uses the AuthMiddleware
	handler := middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Handler should not be called with an invalid token")
	}))

	// Serve the HTTP request
	handler.ServeHTTP(rr, req)

	// Check the status code
	if rr.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %v, but got %v", http.StatusUnauthorized, rr.Code)
	}
}

// Helper function to generate a valid token
func generateValidToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "John Doe"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MDYxMjI1MjgsInVzZXIiOiJKb2huIERvZSJ9.Ni7icx3noQB_N18y6lkF-FA0qV4yEcCkrjwmjj42tzY"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
