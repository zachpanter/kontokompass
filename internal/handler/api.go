package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zachpanter/kontokompass/internal/config"
	"github.com/zachpanter/kontokompass/internal/storage"
	"net/http"
	"strconv"
)

// TODO: GET transactions (with filtering options (date range, account, category)

// TODO: categories
// GET: Lists available budgeting categories.
// POST: Creates a new custom category.

// TODO: /categories/{category_id}:
// GET: Gets details of a specific category.
// PUT: Edits a category.
// DELETE: Deletes a category.

type API struct {
	Router *gin.Engine
	ctx    context.Context
	dbConn *storage.Queries
	conf   *config.Config
}

// NewAPI is the constructor for the API class
func NewAPI(ctx context.Context, conf *config.Config, dbConn *storage.Queries) *API {

	api := &API{
		Router: gin.Default(),
		ctx:    ctx,
		dbConn: dbConn,
		conf:   conf,
	}

	// Routes
	api.Router.GET("/transaction/:transaction_id", api.TransactionGet)

	api.Router.POST("/transaction", api.TransactionPost)

	// Reach via: http://localhost:8080/swagger/index.html
	api.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return api
}

// TransactionGet godoc
// @Summary Selects a transaction.
// @Description Gets a transaction using it's id.
// @Produce  json
// @Param name query string false "Name to greet"
// @Success 200 {object} map[string]string
// @Router /transaction [get]
func (a *API) TransactionGet(ctx *gin.Context) {
	idString := ctx.Param("transaction_id")
	id, convErr := strconv.Atoi(idString)
	if convErr != nil {
		// TODO: Log it
		ctx.AbortWithError(http.StatusInternalServerError, convErr)
		return
	}
	ta, taSelectErr := a.dbConn.TransactionSelect(ctx, int32(id))
	if taSelectErr != nil {
		// TODO: Log it
		ctx.AbortWithError(http.StatusInternalServerError, taSelectErr)
		return
	}

	ctx.JSON(http.StatusOK, ta)
}

// TransactionPost godoc
// @Summary Inserts a transaction into the DB
// @Description Receives a transaction payload via a POST and then inserts it into the DB
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /transaction [post]
func (a *API) TransactionPost(ctx *gin.Context) {
	var payload storage.TransactionInsertParams

	// Bind the incoming JSON to the payload struct
	if bindErr := ctx.ShouldBindJSON(&payload); bindErr != nil {
		// TODO: Log it
		ctx.AbortWithError(http.StatusBadRequest, bindErr)
		return
	}

	// Do something with the received data

	insertTransactionErr := a.dbConn.TransactionInsert(a.ctx, payload)
	if insertTransactionErr != nil {
		// TODO: Log it
		ctx.AbortWithError(http.StatusInternalServerError, insertTransactionErr)
		return
	}

	// Response
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Payload received!",
		"data":    payload,
	})
}
