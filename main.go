package main

import (
	"os"
	"strings"

	"github.com/ghostx31/nativefier-downloader/internal/server"
	"github.com/ghostx31/nativefier-downloader/internal/structs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// main function is used to run the server and interact with other packages
func main() {
	e := echo.New()
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "static/dist",
		Index:  "home.html",
		Browse: false,
		HTML5:  true,
	}))
	e.POST("/save", func(c echo.Context) error {
		Url, Os, widewine := c.FormValue("Url"), c.FormValue("Os"), c.FormValue("widewine")

		urlparams := structs.Urlparams{
			Url:      Url,
			Os:       Os,
			Widewine: widewine,
		}
		file := server.GetUrlFromUser(urlparams)
		defer os.Remove(file) // Remove the zip file
		dirName := strings.Trim(file, ".zip")
		defer os.RemoveAll(dirName) // Remove the folder from which zip was created

		return c.Attachment(file, file)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
