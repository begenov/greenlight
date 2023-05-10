package main

import (
	"net/http"
	"time"

	"github.com/begenov/greenlight/internal/data"
	"github.com/gin-gonic/gin"
)

func (app *application) healthcheckHandler(ctx *gin.Context) {

	env := envelope{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}
	err := app.writeJSON(ctx, http.StatusOK, env, nil)
	if err != nil {
		app.errorResponse(ctx, 500, serverError, err)
	}
}

func (app *application) createMovieHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "create a new movie")
}

func (app *application) showMovieHandler(ctx *gin.Context) {
	id, err := app.readIDParam(ctx.Params)
	if err != nil {
		app.errorResponse(ctx, http.StatusNotFound, notFoundError, err)
	}

	movie := data.Movie{
		ID:       id,
		CreateAt: time.Now(),
		Title:    "Casablanca",
		Runtime:  102,
		Genres:   []string{"drama", "romance", "war"},
		Version:  1,
	}

	err = app.writeJSON(ctx, http.StatusOK, envelope{"movies": movie}, nil)
	if err != nil {
		app.errorResponse(ctx, http.StatusInternalServerError, serverError, err)
	}
}
