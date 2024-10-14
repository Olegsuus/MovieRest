package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *MovieHandlers) Remove(c *gin.Context) {
	id := c.Param("id")

	if err := h.mhP.Remove(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie removed successfully"})
}
