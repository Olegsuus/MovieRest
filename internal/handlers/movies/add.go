package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddMovieDTO struct {
	Title string `json:"title" binding:"required"`
}

func (h *MovieHandlers) Add(c *gin.Context) {
	var dto AddMovieDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	id, err := h.mhP.Add(c.Request.Context(), &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
