package http

import (
	"log"
	"net/http/httputil"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func SetupRoutes(router *gin.Engine, inventoryProxy, orderProxy, userProxy *httputil.ReverseProxy) {

	inventory := router.Group("/inventory")
	{
		inventory.Any("/*path", gin.WrapH(inventoryProxy))
	}

	orders := router.Group("/orders")
	{
		orders.Any("/*path", gin.WrapH(orderProxy))
	}

	// User service routes - forward all requests at the users path to the user service
	users := router.Group("/users")
	{
		users.Any("/*path", forwardUserID(), gin.WrapH(userProxy))
	}

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}

// forwardUserID extracts user ID from JWT token and adds it to the request headers
func forwardUserID() gin.HandlerFunc {
	return func(c *gin.Context) {
		if userID, exists := c.Get("userID"); exists {
			c.Request.Header.Set("X-User-ID", userID.(string))
		}
		c.Next()
	}
}

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip auth for login and register endpoints
		if c.Request.URL.Path == "/users/login" || c.Request.URL.Path == "/users/register" || c.Request.URL.Path == "/users/health" || c.Request.URL.Path == "/health" {
			c.Next()
			return
		}

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

		// Extract user ID from token claims and set it in context
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if userID, ok := claims["user_id"].(string); ok {
				c.Set("userID", userID)
			}
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
