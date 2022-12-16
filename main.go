package main

import (
	"fmt"
	"os"

	"github.com/ghostx31/nativefier-downloader/server"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	loc, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	f, err := os.Open(loc)
	if err != nil {
		panic(err)
	}
	idk, err := f.ReadDir(0)
	if err != nil {
		panic(err)
	}

	for _, v := range idk {
		fmt.Println(v.Name(), v.IsDir())
	}
	e.File("/", "static/file.html")

	e.POST("/save", func(c echo.Context) error {
		url := c.FormValue("Url")
		os := c.FormValue("Os")

		file := server.GetUrlFromUser(url, os)
		return c.File(file)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
