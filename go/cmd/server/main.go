package main

import "github.com/labstack/echo/v4"

func main() {
	// This is a comment
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
