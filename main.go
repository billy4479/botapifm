package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Debug = true

	// Middlewares
	{
		{
			config := middleware.DefaultLoggerConfig
			config.Format = "${remote_ip} made a ${method} to ${uri} in ${latency_human}: got ${status} ${error}\n"
			// config.Format = "${time_custom} ${remote_ip} made a ${method} to ${uri} in ${latency_human}: got ${status} ${error}\n"
			config.Output = os.Stdout
			// config.CustomTimeFormat = "2006/01/02 15:04:05"
			e.Use(middleware.LoggerWithConfig(config))
		}
		// e.Use(middleware.Logger())
		if !e.Debug {
			e.Use(middleware.Recover())
		}
		e.Use(middleware.Gzip())
		e.Use(middleware.Secure())
	}

	endpoints = make(map[string]string)

	e.POST("/download", MakeEndpoint)
	e.POST("/notifyUpload", NotifyUpload)
	e.GET("/download/:id", Download)
}
