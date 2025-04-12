package http

import (
	"log"
	"net/http/httputil"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func SetupRoutes(router *gin.Engine, inventoryProxy, orderProxy *httputil.ReverseProxy) {
	
	inventory := router.Group("/inventory")
	{
		inventory.Any("/*path", gin.WrapH(inventoryProxy))
	}

	
	orders := router.Group("/orders")
	{
		orders.Any("/*path", gin.WrapH(orderProxy))
	}

	
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := extractToken(c)
		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header required"})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}

		c.Next()
	}
}

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		log.Printf("%s %s %s", c.Request.Method, c.Request.URL.Path, c.Request.RemoteAddr)
		c.Next()

		
		log.Printf("Response: %d", c.Writer.Status())
	}
}

func extractToken(c *gin.Context) string {
	bearerToken := c.GetHeader("Authorization")
	if strings.HasPrefix(bearerToken, "Bearer ") {
		return strings.TrimPrefix(bearerToken, "Bearer ")
	}
	return ""
}
