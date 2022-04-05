package brand

import "idnmedia/database"

func (b *Brands) FindAll() error {
	db, err := database.ConnectionDB()
	if err != nil {
		return err
	}
	db.Find(&b)
	return nil
}

func (b *Brand) FindByID(id int) error {
	db, err := database.ConnectionDB()
	if err != nil {
		return err
	}
	db.Where("sku = ? ", id).Find(b)
	return nil
}

func (b *Brand) Create() error {
	db, err := database.ConnectionDB()
	if err != nil {
		return err
	}
	db.Create(&b)
	return nil
}