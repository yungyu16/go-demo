//go:generate stringer -type=Pill

package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Pill int

const name = len("11")
const (
	Placebo Pill = iota
	Aspirin
	Aspirin2
	Aspirin3
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
	for i2 := range 3 {
		fmt.Println(i2)
	}
}
