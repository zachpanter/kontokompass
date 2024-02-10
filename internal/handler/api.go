package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

// NewAPI is the constructor for the API class
func NewAPI(ctx context.Context, queries *storage.Queries) {
	router := gin.Default()

	// Routes
	router.GET("/greeting", GetGreeting)

	// Reach via: http://localhost:8080/swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run
	router.Run()
}

// GetGreeting godoc
// @Summary Returns a simple greeting message.
// @Description Greets the user with name if provided in query string.
// @Produce  json
// @Param name query string false "Name to greet"
// @Success 200 {object} map[string]string
// @Router /greeting [get]
func GetGreeting(ctx *gin.Context) {
	name := ctx.Query("name")
	message := "Hello"

	if name != "" {
		message += ", " + name
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
