package routers

import (
	handlers "github.com/Olegsuus/MovieRest/internal/handlers/movies"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, movieHandler *handlers.MovieHandlers) {
	router.POST("/movies", movieHandler.Add)
	router.GET("/movies/:id", movieHandler.Get)
	router.GET("/movies", movieHandler.GetMany)
	router.PATCH("/movies/:id", movieHandler.Update)
	router.DELETE("/movies/:id", movieHandler.Remove)
}
