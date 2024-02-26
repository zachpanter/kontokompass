package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zachpanter/kontokompass/internal/config"
	"github.com/zachpanter/kontokompass/internal/storage"
	"net/http"
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
	router *gin.Engine
	ctx    context.Context
	dbConn *storage.Queries
	conf   *config.Config
}

// NewAPI is the constructor for the API class
func NewAPI(ctx context.Context, conf *config.Config, dbConn *storage.Queries) *API {

	api := &API{
		router: gin.Default(),
		ctx:    ctx,
		dbConn: dbConn,
		conf:   conf,
	}

	// Routes
	api.router.GET("/greeting", api.TransactionGet)

	api.router.POST("/transaction", api.TransactionPost)

	// Reach via: http://localhost:8080/swagger/index.html
	api.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run
	go func() {
		err := api.router.Run()
		if err != nil {
			panic(err)
		}
	}()
	return api
}

// TransactionGet godoc
// @Summary Selects a transaction.
// @Description Gets a transaction using it's id.
// @Produce  json
// @Param name query string false "Name to greet"
// @Success 200 {object} map[string]string
// @Router /greeting [get]
func (a *API) TransactionGet(ctx *gin.Context) {
	name := ctx.Query("name")
	message := "Hello"

	if name != "" {
		message += ", " + name
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

// TransactionPost godoc
// @Summary Inserts a transaction into the DB
// @Description Receives a transaction payload via a POST and then inserts it into the DB
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /transaction [post]
func (a *API) TransactionPost(c *gin.Context) {
	var payload storage.TransactionInsertParams

	// Bind the incoming JSON to the payload struct
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Do something with the received data

	insertTransactionErr := a.dbConn.TransactionInsert(a.ctx, payload)
	if insertTransactionErr != nil {
		fmt.Printf("%e", insertTransactionErr)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": insertTransactionErr,
		})
	}

	// Response:
	c.JSON(http.StatusOK, gin.H{
		"message": "Payload received!",
		"data":    payload,
	})
}
