package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *MovieHandlers) Update(c *gin.Context) {
	id := c.Param("id")
	var movieUpdate map[string]interface{}
	if err := c.ShouldBindJSON(&movieUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	movie, err := h.mhP.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if title, ok := movieUpdate["title"].(string); ok {
		movie.Title = title
	}

	if description, ok := movieUpdate["description"].(string); ok {
		movie.Description = description
	}
	if year, ok := movieUpdate["year"].(float64); ok {
		movie.Year = int32(year)
	}
	if country, ok := movieUpdate["country"].(string); ok {
		movie.Country = country
	}
	if genres, ok := movieUpdate["genres"].([]interface{}); ok {
		var genresStr []string
		for _, g := range genres {
			if genreStr, ok := g.(string); ok {
				genresStr = append(genresStr, genreStr)
			}
		}
		movie.Genres = genresStr
	}
	if posterURL, ok := movieUpdate["poster_url"].(string); ok {
		movie.PosterURL = posterURL
	}
	if rating, ok := movieUpdate["rating"].(float64); ok {
		movie.Rating = float32(rating)
	}

	err = h.mhP.Update(c.Request.Context(), id, movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie updated successfully"})
}
