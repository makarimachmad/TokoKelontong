package transaction

import (
	"fmt"
	"net/http"
	"strings"

	"idnmedia/helper"
	"idnmedia/module/product"
	"idnmedia/module/user"

	"github.com/labstack/echo"
)

func Kasir(ctx echo.Context) error{
	//setup response
	resp := new(helper.Response)

	//get data request
	pembelian := new(Pembelian)
	ctx.Bind(&pembelian)

	//find user by id from mysql
	user := new(user.User)
	user.FindByID(pembelian.UserID)

	//set value response user from struct user
	resp.User.Name = user.Name
	resp.User.Email = user.Email
	resp.User.Address = user.Address

	//setup value qty for count each qty sku
	qtyfr1 := 0
	qtysr1 := 0
	qtycf1 := 0

	hargafr1 := 0
	hargasr1 := 0
	hargacf1 := 0

	total := 0
	hargatotal := 0
	diskon := 0
	
	reason := []string{}

	//check item and count qty
	//find product by sku from mysql
	for _, j := range pembelian.Items{
		if j.Sku == "SR1"{

			qtysr1 = qtysr1 + j.Qty
			if hargasr1 == 0 {
				product := new(product.Product)
				product.FindBySKU(j.Sku)
				hargasr1 = product.Price
			}
		} else if j.Sku == "FR1"{
			
			qtyfr1 = qtyfr1 + j.Qty
			if hargafr1 == 0 {
				product := new(product.Product)
				product.FindBySKU(j.Sku)
				hargafr1 = product.Price
			}
		}else if j.Sku == "CF1"{
			qtycf1 = qtycf1 + j.Qty
			if hargacf1 == 0 {
				product := new(product.Product)
				product.FindBySKU(j.Sku)
				hargacf1 = product.Price
			}
		}
	}

	//check discount
	//discount strawberries
	if qtysr1 >= 3{
		fmt.Println("beli 3 atau lebih strawberries, dapat diskon Rp 4500")
		hargasr1 = 4500
		hargasr1 = (qtysr1 * hargasr1)
		diskon = 500 * qtyfr1
		reason = append(reason, "beli 3 atau lebih strawberries, dapat diskon Rp 4500")
	}
	
	//discount fruit tea
	sisa := 0
	if qtyfr1 %2 != 0{
		sisa = qtyfr1 %2
	}

	fmt.Println("beli 1 gratis 1 fruit tea")
	qtyfr1 = qtyfr1-sisa
	hargafr1 = (qtyfr1 * hargafr1) / 2 + (hargafr1 * sisa)
	diskon = diskon + hargafr1
	reason = append(reason, "beli 1 gratis 1 fruit tea")

	hargatotal = hargafr1 + hargasr1 + hargacf1
	total = diskon + hargatotal

	//set value response items from struct array item
	if qtycf1 != 0{
		respitem := new(helper.Item)
		respitem.Name = "Coffe"
		respitem.Sku = "CF1"
		respitem.Qty = qtycf1

		resp.Items = append(resp.Items, *respitem)
	}
	if qtyfr1 != 0{
		respitem := new(helper.Item)
		respitem.Name = "Fruit tea"
		respitem.Sku = "FR1"
		respitem.Qty = qtyfr1
		
		resp.Items = append(resp.Items, *respitem)
	}
	if qtysr1 != 0{
		respitem := new(helper.Item)
		respitem.Name = "Strawberries"
		respitem.Sku = "SR1"
		respitem.Qty = qtysr1

		resp.Items = append(resp.Items, *respitem)
	}

	resp.Total = total
	resp.Discount = diskon
	resp.GrandTotal = hargatotal
	resp.Reason = strings.Join(reason, ";")
	
	return ctx.JSON(http.StatusOK, resp)
}