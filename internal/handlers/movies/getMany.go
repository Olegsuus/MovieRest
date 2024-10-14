package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *MovieHandlers) GetMany(c *gin.Context) {
	movies, err := h.mhP.GetMany(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}
