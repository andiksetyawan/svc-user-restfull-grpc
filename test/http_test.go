package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	userv1 "svc-user/api/v1"
	"svc-user/internal/repository"
	"svc-user/internal/router"
	"svc-user/internal/service"
	"svc-user/internal/store"
)

func setupRouterTest() *http.ServeMux {
	db := store.NewPostgreeSQLDbTest().Connect()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	router := router.NewRouter("", userService)
	return router.Mux()
}

func TestHTTP1Create(t *testing.T) {
	r := setupRouterTest()
	w := httptest.NewRecorder()

	reqPayload := strings.NewReader("{\"name\":\"John Due Test\"}")
	req, _ := http.NewRequest(http.MethodPost, "/user.v1.UserService/Create", reqPayload)
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	var response userv1.CreateResponse
	json.NewDecoder(w.Body).Decode(&response)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "OK", response.Message)
}
