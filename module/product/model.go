package product

import (
	"idnmedia/database"
)

func (p *ListProducts) FindAll() error {
	db, err := database.ConnectionDB()
	if err != nil{
		return err
	}
	db.Find(&p)
	return nil
}

func (p *Product) FindBySKU(sku string) error {
	db, err := database.ConnectionDB()
	if err != nil{
		return err
	}
	db.Where("sku = ? ", sku).Find(p)
	return nil
}

func (p *Product)Create() error{
	db, err := database.ConnectionDB()
	if err != nil {
		return err
	}
	db.Create(&p)
	return nil
}