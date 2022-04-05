package product

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func GetProducts(ctx echo.Context) error {
	products := new(ListProducts)
	err := products.FindAll()
	if err != nil{
		fmt.Println("error", err)
		return ctx.JSON(http.StatusNotFound, "error, data enggak ada")
	}
	return ctx.JSON(http.StatusOK, products)
}