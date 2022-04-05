package router

import (
	"idnmedia/module/product"
	"idnmedia/module/transaction"
	"idnmedia/module/user"

	"github.com/labstack/echo"
)

func Routes(app *echo.Echo){
	api := app.Group("/api")
	
	u := api.Group("/user")
	u.GET("/all", user.GetUsers)

	p := api.Group("/product")
	p.GET("/all", product.GetProducts)

	t := api.Group("/transaksi")
	t.GET("/pembelian", transaction.Kasir)
}