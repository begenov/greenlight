package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) healthcheckHandler(ctx *gin.Context) {

	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}
	err := app.writeJSON(ctx, http.StatusOK, data, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"ERROR": "The server encountered a problem and could not process your request",
		})
		return
	}
}

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
