package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	serverError           = "the server encoutered a probleb and could not process your request"
	notFoundError         = "the requested resource could not be found"
	methodNotAllowedError = "the %s method is not supported for this resource"
)

func (app *application) logError(r *http.Request, err error) {
	app.logger.Print(err)

}

func (app *application) errorResponse(ctx *gin.Context, status int, message any, e error) {
	app.logError(ctx.Request, e)
	env := envelope{"error": message}
	err := app.writeJSON(ctx, status, env, nil)
	if err != nil {
		app.logError(ctx.Request, err)
		ctx.Status(http.StatusInternalServerError)
	}
}
