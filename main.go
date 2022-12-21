package main

import (
	"os"
	"strings"

	"github.com/ghostx31/nativefier-downloader/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "static",
		Index:  "file.html",
		Browse: false,
		HTML5:  true,
	}))
	e.POST("/save", func(c echo.Context) error {
		url, Os := c.FormValue("Url"), c.FormValue("Os")
		file := server.GetUrlFromUser(url, Os)

		defer os.Remove(file) // Remove the zip file
		dirName := strings.Trim(file, ".zip")
		defer os.RemoveAll(dirName) // Remove the folder from which zip was created

		return c.Attachment(file, file)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
