package route

import (
	"fmt"
	"goProject/dictionary"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationRoutes(t *testing.T) {
	// Créez un dictionnaire pour les tests
	d := dictionary.New("test_route.txt")
	defer d.Close()

	// Créez le routeur en utilisant le dictionnaire de test
	router := NewRouter(d)

	// Test d'ajout
	t.Run("AddEntry", func(t *testing.T) {
		word := "testWord"
		definition := "testDefinition"
		url := fmt.Sprintf("/add/%s/%s", word, definition)

		req, err := http.NewRequest("POST", url, nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	// Test de récupération
	t.Run("GetDefinition", func(t *testing.T) {
		word := "testWord"
		url := fmt.Sprintf("/get/%s", word)

		req, err := http.NewRequest("GET", url, nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Definition of")
	})

	// Test de suppression
	t.Run("RemoveEntry", func(t *testing.T) {
		word := "testWord"
		url := fmt.Sprintf("/remove/%s", word)

		req, err := http.NewRequest("DELETE", url, nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

}
