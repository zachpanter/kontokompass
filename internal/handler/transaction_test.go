package handler_test

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/zachpanter/kontokompass/internal/config"
	"github.com/zachpanter/kontokompass/internal/handler"
	"github.com/zachpanter/kontokompass/internal/storage"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMyGinEndpoint(t *testing.T) {
	// SETUP
	router := gin.Default()
	ctx := context.Background()
	conf := config.NewConfig()
	dbConn := storage.OpenDBPool(ctx, conf)

	newHandler := handler.NewAPI(ctx, conf, dbConn)
	router.GET("/transaction", newHandler.TransactionGet)
	//w := httptest.NewRecorder()      // Response recorder
	//c, _ := gin.CreateTestContext(w) // Test context

	request, requestErr := http.NewRequest("GET", "/transaction", nil)
	if requestErr != nil {
		t.Fatal(requestErr)
	}

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "application/json; charset=utf-8", response.Header().Get("Content-Type"))
}
