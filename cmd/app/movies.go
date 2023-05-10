package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) createMovieHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "create a new movie")
}

func (app *application) showMovieHandler(ctx *gin.Context) {
	id, err := app.readIDParam(ctx.Params)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"ERROR": "showMovieHandler Not Found",
		})
		return
	}
	ctx.String(http.StatusOK, "show the details of movie %d", id)
}
