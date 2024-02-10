package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/zachpanter/kontokompass/cmd/kontokompass/docs"
	"net/http"
)

// @title           KontoKompass
// @version         1.0
// @description     A Personal Finance solution for planning and analytics.
// @contact.name   Zach Panter
// @contact.email  zachpanter@gmail.com
// @license.name  N/A
// @license.url   https://choosealicense.com/
// @host      localhost:8080
// @BasePath  /api/v1
// @securityDefinitions.basic  BasicAuth
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

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
