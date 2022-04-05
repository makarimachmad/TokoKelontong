package transaction

import "idnmedia/database"

func (t *Transaction) Create() error {
	db, err := database.ConnectionDB()
	if err != nil {
		return err
	}
	db.Create(&t)
	return nil
}

func (t *Transactions) FindAll() error {
	db, err := database.ConnectionDB()
	if err != nil{
		return err
	}
	db.Find(&t)
	return nil
}

func (t *Transaction) FindByID(id string) error {
	db, err := database.ConnectionDB()
	if err != nil{
		return err
	}
	db.Where("sku = ? ", id).Find(t)
	return nil
}