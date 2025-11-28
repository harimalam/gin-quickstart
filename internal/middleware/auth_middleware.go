package middleware

import (
	"gin-quickstart/internal/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const BearerSchema = "Bearer "
const contextClaimsKey = "claims"

func AuthMiddleware(secret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), BearerSchema)

		claims, err := auth.VerifyToken(tokenString, secret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing token"})
			return
		}
		// Store claims in context for further handlers to use
		c.Set(contextClaimsKey, claims)
		// Proceed to the next handler
		c.Next()
	}
}

func Authorize(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Get the Claims from the context (set by AuthMiddleware)
		claimsRaw, exists := c.Get(contextClaimsKey)

		// Safety check (shouldn't happen if AuthMiddleware ran)
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied. Claims missing."})
			return
		}

		// 2. Cast the claims back to the *auth.Claims type
		claims, ok := claimsRaw.(*auth.Claims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied. Claims malformed."})
			return
		}

		// 3. Check for role match
		if claims.Role != requiredRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden. Insufficient role privileges."})
			return
		}

		// 4. Proceed if authorized
		c.Next()
	}
}
