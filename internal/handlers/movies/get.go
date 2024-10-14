package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *MovieHandlers) Get(c *gin.Context) {
	id := c.Param("id")

	movie, err := h.mhP.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movie)
}
