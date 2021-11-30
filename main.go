package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io/ioutil"
	"net/http"
	"strings"
	"unicode"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "public")
	e.POST("/upload", upload)

	e.Logger.Fatal(e.Start(":1323"))
}

func upload(c echo.Context) error {

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	b, err := ioutil.ReadAll(src)

	words := parseText(string(b))
	return c.HTML(http.StatusOK, fmt.Sprintf("<p>%s</p>", strings.Join(words, "<br>")))
}

func parseText(text string) []string {
	var result []string
	words := map[string]struct{}{}
	for _, field := range strings.Fields(text) {
		if unicode.IsLetter(rune(field[0])) {
			if _, ok := words[field]; !ok {
				words[field] = struct{}{}
				result = append(result, field)
			}
		}
	}

	return result
}
