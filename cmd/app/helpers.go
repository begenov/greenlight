package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *application) readIDParam(params gin.Params) (int64, error) {
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

func (app *application) writeJSON(ctx *gin.Context, status int, data any, headers http.Header) error {
	var data1 any
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		ctx.Writer.Header()[key] = value
	}

	err = json.Unmarshal(js, &data1)
	if err != nil {
		app.logger.Printf(err.Error())
		return err
	}
	ctx.JSON(status, data1)
	return nil
}
