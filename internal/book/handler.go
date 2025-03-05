package book

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

// CreateBook POST /books
func (h *Handler) CreateBook(c *gin.Context) {
	var req struct {
		Title       string `json:"title"`
		Author      string `json:"author"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	b, err := h.service.CreateBook(req.Title, req.Author, req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, b)
}

// ListBooks GET /books
func (h *Handler) ListBooks(c *gin.Context) {
	books, err := h.service.ListBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

// LoanBook POST /books/:id/loan
func (h *Handler) LoanBook(c *gin.Context) {
	userIDInterface, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no user in context"})
		return
	}
	userID := userIDInterface.(uint)

	bookIDParam := c.Param("id")
	bookID64, err := strconv.ParseUint(bookIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}
	bookID := uint(bookID64)

	var req struct {
		DueDate string `json:"due_date"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	due, err := time.Parse("2006-01-02", req.DueDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "due_date format should be YYYY-MM-DD"})
		return
	}

	lr, err := h.service.LoanBook(bookID, userID, due)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lr)
}

// ReturnBook POST /books/:id/return
func (h *Handler) ReturnBook(c *gin.Context) {
	userIDInterface, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no user in context"})
		return
	}
	_ = userIDInterface.(uint) // можно проверить что этот пользователь был тем, кто взял книгу

	bookIDParam := c.Param("id")
	bookID64, err := strconv.ParseUint(bookIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}
	bookID := uint(bookID64)

	lr, err := h.service.ReturnBook(bookID, 0) // userID=0 - упрощённо
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lr)
}
