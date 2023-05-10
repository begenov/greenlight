package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *application) createMovieHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "create a new movie")
}

func (app *application) showMovieHandler(ctx *gin.Context) {
	params := ctx.Params

	p, ok := params.Get("id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": "showMovieHandler id",
		})
		return
	}
	id, err := strconv.ParseInt(p, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"ERROR": "showMovieHandler Not Found",
		})
		return
	}
	ctx.String(http.StatusOK, "show the details of movie %d", id)
}
