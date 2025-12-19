// server/internal/router/routes_test.go
package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := Routes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Pong")
}

func TestLoginRoute_Exists(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := Routes()

	req, _ := http.NewRequest("POST", "/api/auth/login", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.NotEqual(t, http.StatusNotFound, w.Code)
}

func TestRegisterRoute_Exists(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := Routes()
	req, _ := http.NewRequest("POST", "/api/auth/register", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.NotEqual(t, http.StatusNotFound, w.Code)
}

func TestGetAllPastesRoute_Exists(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := Routes()
	req, _ := http.NewRequest("POST", "/api/auth/register", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.NotEqual(t, http.StatusNotFound, w.Code)
}
