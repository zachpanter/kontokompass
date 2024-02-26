package handler_test

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/zachpanter/kontokompass/internal/config"
	"github.com/zachpanter/kontokompass/internal/handler"
	"github.com/zachpanter/kontokompass/internal/storage"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type MockDB struct {
}

func (mdb MockDB) TransactionSelect(ctx context.Context, taID int32) (storage.TransactionSelectRow, error) {
	return storage.TransactionSelectRow{
		TaPostdate:           time.Time{},
		TaDescription:        "",
		TaDebit:              sql.NullFloat64{},
		TaCredit:             sql.NullFloat64{},
		TaBalance:            0,
		TaClassificationText: "",
	}, nil
}

func (mds MockDB) TransactionInsert(ctx context.Context, arg storage.TransactionInsertParams) error {
	return nil
}

func TestMyGinEndpoint(t *testing.T) {
	// SETUP
	ctx := context.Background()
	conf := config.NewConfig()
	//dbConn := storage.OpenDBPool(ctx, conf)
	dbConn := MockDB{}

	newHandler := handler.NewAPI(ctx, conf, dbConn)

	// Set Gin to Test mode to disable logging
	gin.SetMode(gin.TestMode)

	// Create a test request
	req, err := http.NewRequest("GET", "/transaction/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder
	rr := httptest.NewRecorder()

	// Serve the request through the router
	newHandler.Router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json; charset=utf-8", rr.Header().Get("Content-Type"))
}
