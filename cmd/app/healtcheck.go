package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) healthcheckHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello Wolrd")
}
