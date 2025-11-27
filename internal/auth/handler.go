package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	Service AuthService
}

// NewHandler creates a new instance of Handler with the given AuthService.
func NewHandler(service AuthService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) RegisterRoutes(g *gin.RouterGroup) {
	authGroup := g.Group("/auth")
	{
		authGroup.POST("/signup", h.SignUp)
		authGroup.POST("/login", h.Login)
	}
}

// SignUp handles user registration requests.
func (h *Handler) SignUp(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service to register user
	user, err := h.Service.SignUp(req)
	if err != nil {
		// Check for username collision (GORM unique constraint violation)
		if strings.Contains(err.Error(), "UNIQUE constraint failed") || strings.Contains(err.Error(), "Duplicate entry") {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"}) // 409 Conflict
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User registration failed"}) // 500 Internal Server Error
		}
		return
	}

	// Successful registration
	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"user": gin.H{
				"id":       user.ID,
				"username": user.Username,
				"role":     user.Role,
			},
		},
		"message": "User registered successfully",
	})
}

// Login handles user authentication and JWT token generation.
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service to authenticate user and generate token
	token, err := h.Service.Login(req)
	if err != nil {
		// Check specifically for service errors (invalid credentials, user not found)
		if errors.Is(err, ErrInvalidCredentials) || errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"}) // 401 Unauthorized
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Login failed"}) // 500 Internal Server Error
		return
	}

	// Successful login
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"token": token,
		},
		"message": "Login successful",
	})
}
