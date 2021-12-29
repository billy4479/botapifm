package main

import (
	"io/ioutil"
	"os"

	"github.com/labstack/echo/v4"
)

func NotifyUpload(c echo.Context) error {
	path, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	defer c.Request().Body.Close()

	return os.Remove(string(path))
}
