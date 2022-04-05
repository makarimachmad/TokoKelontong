package user

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetUsers(ctx echo.Context) error {
	user := new(Users)
	err := user.FindAll()
	if err != nil {
		return ctx.JSON(http.StatusNotFound, true)
	}
	return ctx.JSON(http.StatusOK, user)
}

func PostUser(ctx echo.Context) error{
	user := new(User)
	err := ctx.Bind(user)
	if err != nil{
		return ctx.JSON(http.StatusNotFound, true)
	}
	err = user.Create()
	if err != nil{
		return ctx.JSON(http.StatusNotFound, true)
	}
	return ctx.JSON(http.StatusOK, user)
}