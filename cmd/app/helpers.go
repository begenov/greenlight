package main

import (
	"errors"
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
