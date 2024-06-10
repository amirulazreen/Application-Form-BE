package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func CorsMiddleware() gin.HandlerFunc {
    handler := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    })

    return func(c *gin.Context) {
        handler.HandlerFunc(c.Writer, c.Request)

        if c.Request.Method != http.MethodOptions {
            c.Next()
        } else {
            c.AbortWithStatus(http.StatusOK)
        }
    }
}

