package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/user", func(c echo.Context) error {
		user := User{
			FirstName: "James",
			LastName:  "Cefaratti",
			Email:     "jcefaratti@gmail.com",
			Age:       37,
		}

		user2 := User{
			FirstName: "Crosby",
			LastName:  "Cefaratti",
			Email:     "kingcrosby@gmail.com",
			Age:       4,
		}
		users := [2]User{user, user2}
		return c.JSON(http.StatusOK, users)
	})
	e.Logger.Fatal(e.Start(":9000"))
}
