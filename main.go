package main

import (
	"github.com/ghostx31/nativefier-downloader/server"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.File("/", "static/file.html")

	e.POST("/save", func(c echo.Context) error {
		url := c.FormValue("Url")
		os := c.FormValue("Os")

		file := server.GetUrlFromUser(url, os)
		return c.File(file)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
