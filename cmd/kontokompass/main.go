package main

import (
	"context"
	_ "github.com/zachpanter/kontokompass/cmd/kontokompass/docs"
	"github.com/zachpanter/kontokompass/internal/config"
	"github.com/zachpanter/kontokompass/internal/handler"
	"github.com/zachpanter/kontokompass/internal/storage"
)

func main() {
	ctx := context.Background()
	conf := config.NewConfig()
	dbConn := storage.OpenDBPool(ctx, conf)
	handler.NewAPI(ctx, dbConn)
}
