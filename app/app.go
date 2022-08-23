package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

var (
	router = gin.Default()
	port   = os.Getenv("PORT")
)

func Start() {
	// define routes
	router.GET("/health", health)
	router.POST("/sftp", isAuthorized(), SFTP)

	// define default ingress port
	if port == "" {
		port = "9089"
	}

	// configure ingress security policy
	allowedHeaders := handlers.AllowedHeaders([]string{"access-control-allow-headers", "access-control-allow-methods", "access-control-allow-origin", "access-control-max-age", "content-type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST"})

	// Start server
	fmt.Printf("Server started on port %s.\n", port)
	log.Panic(http.ListenAndServe(":"+port, handlers.CORS(allowedHeaders, allowedMethods, allowedOrigins)(router)))
}
