package main

import (
	"os"
	"strconv"
	"strings"

	"golang.org/x/time/rate"

	"github.com/ghostx31/nativefier-downloader/internal/server"
	"github.com/ghostx31/nativefier-downloader/internal/structs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// main function is used to run the server and interact with other packages
func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	rateLimit, err := strconv.Atoi(os.Getenv("RATE_LIMIT"))
	if err != nil {
		rateLimit = 20
	}
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(rateLimit))))

	// Serve static files
	e.File("/favicon.ico", "static/dist/assets/favicon.ico")
	e.File("/home.css", "static/dist/home.css")
	e.File("/assets/favicon.ico", "static/dist/assets/favicon.ico")

	// Serve routes
	e.GET("/", func(c echo.Context) error {
		return c.File("static/dist/home.html")
	})

	e.GET("/about", func(c echo.Context) error {
		return c.File("static/dist/about.html")
	})

	e.GET("/usage", func(c echo.Context) error {
		return c.File("static/dist/usage.html")
	})

	e.POST("/save", func(c echo.Context) error {
		Url, Os, widewine, tray := c.FormValue("Url"), c.FormValue("Os"), c.FormValue("widewine"), c.FormValue("tray")

		urlparams := structs.Urlparams{
			Url:      Url,
			Os:       Os,
			Widewine: widewine,
			Tray:     tray,
		}
		file := server.GetUrlFromUser(urlparams)
		defer os.Remove(file) // Remove the zip file
		dirName := strings.Trim(file, ".zip")
		defer os.RemoveAll(dirName) // Remove the folder from which zip was created

		return c.Attachment(file, file)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
