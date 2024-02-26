package handler_test

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
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

func SetupTesting() *handler.API {
	// SETUP
	ctx := context.Background()
	conf := config.NewConfig()
	//dbConn := storage.OpenDBPool(ctx, conf)
	dbConn := MockDB{}

	newHandler := handler.NewAPI(ctx, conf, dbConn)

	// Set Gin to Test mode to disable logging
	gin.SetMode(gin.TestMode)
	return newHandler
}

func TestTransactionSelect(t *testing.T) {
	newHandler := SetupTesting()

	// Create a test request
	req, err := http.NewRequest("GET", "/transaction/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder
	rr := httptest.NewRecorder()

	// Serve the request through the router
	newHandler.Router.ServeHTTP(rr, req)

	// ASSERTS
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json; charset=utf-8", rr.Header().Get("Content-Type"))
}

func TestTransactionInsert(t *testing.T) {
	newHandler := SetupTesting()

	//payload := map[string]interface{}{
	//	"name":  "John Doe",
	//	"email": "john.doe@example.com",
	//}

	var payload storage.TransactionInsertParams
	payload.TaClassificationText = "foo_class"
	payload.TaDescription = "foo_desc"
	payload.TaCredit.Float64 = 0.0
	payload.TaDebit.Float64 = 0.0
	payload.TaBalance = 0.0
	payload.TaPostdate = time.Now()

	payloadBytes, _ := json.Marshal(payload)

	// Create a test request
	req, err := http.NewRequest("POST", "/transaction", bytes.NewBuffer(payloadBytes))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder
	rr := httptest.NewRecorder()

	// Serve the request through the router
	newHandler.Router.ServeHTTP(rr, req)

	// ASSERTS
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json; charset=utf-8", rr.Header().Get("Content-Type"))
}
