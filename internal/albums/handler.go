package albums

import (
	"gin-quickstart/internal/middleware"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Handler holds the necessary dependencies for the album handlers.
type Handler struct {
	service Service
}

// NewHandler is the constructor for Handler.
func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

// RegisterRoutes attaches album routes to the Gin engine.
func (h *Handler) RegisterRoutes(g *gin.RouterGroup) {
	albumGroup := g.Group("/albums")
	{
		// 1. READ routes (accessible to anyone with a valid token: 'user' or 'admin')
		albumGroup.GET("/", h.GetAlbums)
		albumGroup.GET("/:id", h.GetAlbumByID)

		// 2. WRITE routes (only accessible to 'admin')
		adminAlbumGroup := albumGroup.Group("/")
		adminAlbumGroup.Use(middleware.Authorize("admin"))
		{
			adminAlbumGroup.POST("/", h.CreateAlbum)
			adminAlbumGroup.PUT("/:id", h.UpdateAlbum)
			adminAlbumGroup.DELETE("/:id", h.DeleteAlbum)
		}
	}
}

// GetAlbums retrieves all albums via the service and returns a 200 OK response.
func (h *Handler) GetAlbums(c *gin.Context) {
	// Call service to get all albums
	albums, err := h.service.FindAll()
	if err != nil {
		// Handle error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return albums with 200 OK status
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"albums": albums,
		},
		"message": "Albums retrieved successfully",
	})
}

// CreateAlbum processes a POST request to add a new album.
func (h *Handler) CreateAlbum(c *gin.Context) {
	var album Album

	// Bind JSON body to album struct
	if err := c.BindJSON(&album); err != nil {
		// Handle JSON binding error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service to create album
	created, err := h.service.Create(album)
	if err != nil {
		//	 Handle creation error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return created album with 201 Created status
	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"album": created,
		},
		"message": "Album created successfully",
	})
}

// GetAlbumByID retrieves a single album by its ID.
func (h *Handler) GetAlbumByID(c *gin.Context) {
	idStr := c.Param("id")
	// 1. Convert string URL param to uint
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format. Must be an integer."})
		return
	}

	// 2. Call service to find by ID
	album, err := h.service.FindById(uint(idUint))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Handle not found error
			c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
		} else {
			// Handle other errors
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// 3. Return found album
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"album": album,
		},
		"message": "Album retrieved successfully",
	})
}

// UpdateAlbum processes a PUT request to update an existing album.
func (h *Handler) UpdateAlbum(c *gin.Context) {
	var album Album
	idStr := c.Param("id")

	// 1. Convert string URL param to uint
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format. Must be an integer."})
		return
	}

	// 2. Bind JSON body to album struct
	if err := c.BindJSON(&album); err != nil {
		// Handle JSON binding error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 3. Set the ID from the URL param
	album.ID = uint(idUint)

	// 4. Call service to update
	updated, err := h.service.Update(album)
	if err != nil {
		// Handle not found error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 5. Return updated album
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"album": updated,
		},
		"message": "Album updated successfully",
	})
}

// delete album by ID
func (h *Handler) DeleteAlbum(c *gin.Context) {
	idStr := c.Param("id")

	// 1. Convert string URL param to uint
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format. Must be an integer."})
		return
	}

	// 2. Call service to delete
	if err := h.service.Delete(uint(idUint)); err != nil {
		// Handle not found error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 3. Return no content status
	c.Status(http.StatusNoContent)
}
