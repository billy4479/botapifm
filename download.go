package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var endpoints map[string]string

func MakeEndpoint(c echo.Context) error {
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	defer c.Request().Body.Close()

	path := string(b)
	UUID := uuid.New().String()

	endpoints[UUID] = path

	return c.String(http.StatusOK, UUID)
}

func Download(c echo.Context) error {
	UUID := c.Param("id")
	path, ok := endpoints[UUID]
	if !ok {
		return c.NoContent(http.StatusNotFound)
	}

	err := c.File(path)
	if err != nil {
		return err
	}

	delete(endpoints, UUID)

	return os.Remove(path)
}
