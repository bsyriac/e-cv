package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JWTHandler struct {
}

func RegisterJWTHandler(r *gin.Engine) {

	jwtStruct := JWTHandler{}

	jwt := r.Group("/jwt")

	jwt.GET("/", jwtStruct.test)

}

func (jwt *JWTHandler) test(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
	return
}
