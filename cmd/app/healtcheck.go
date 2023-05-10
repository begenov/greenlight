package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) healthcheckHandler(ctx *gin.Context) {

	js := fmt.Sprintf(`{"status": "available", "enironment": %q, "version": %q}`, app.config.env, version)

	ctx.JSON(http.StatusOK, gin.H{
		"INFO": js,
	})
}
