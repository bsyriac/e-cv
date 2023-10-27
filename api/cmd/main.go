package main

import (
	"e-cv/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// create router
	router := gin.Default()

	handlers.RegisterJWTHandler(router)

	router.Run(":8443")
}
