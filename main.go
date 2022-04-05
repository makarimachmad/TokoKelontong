package main

import (
	"fmt"
	"idnmedia/router"
	"net/http"

	"github.com/labstack/echo"
)

//import database
//buat config koneksi
//buat routing endpoint
//buat response

func main() {
	e := echo.New()

	e.GET("/", func(ctx echo.Context) error {
		fmt.Println("IDN media API")
		return ctx.JSON(http.StatusOK, "oke mantap tersambung....")
	})

	router.Routes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
