package main

import (
	"context"
	_ "github.com/zachpanter/kontokompass/docs" // Import for Swagger docs
	"github.com/zachpanter/kontokompass/internal/config"
	"github.com/zachpanter/kontokompass/internal/handler"
	"github.com/zachpanter/kontokompass/internal/storage"
)

// @title Gin Swagger API
// @version 1.0
// @description This is a sample Gin-based API server.
// @contact.name API Support
// @contact.email support@your-domain.com
// @license.name MIT
// @host localhost:8080
// @BasePath /api/v1
func main() {
	ctx := context.Background()
	conf := config.NewConfig()
	dbConn := storage.OpenDBPool(ctx, conf)
	apiHandler := handler.NewAPI(ctx, conf, dbConn)
	routerRunErr := apiHandler.Router.Run()
	if routerRunErr != nil {
		// TODO: Log it
		panic(routerRunErr)
	}
}
