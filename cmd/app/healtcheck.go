package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) healthcheckHandler(ctx *gin.Context) {

	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}
	var data1 map[string]string
	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Printf(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"ERROR": "The server encountered a problem and could not process your request",
		})
		return
	}

	js = append(js, '\n')

	err = json.Unmarshal(js, &data1)
	if err != nil {
		app.logger.Printf(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"ERROR": "The server encountered a problem and could not process your request",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"INFO": data1,
	})
}
