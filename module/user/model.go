package user

import "idnmedia/database"

func (u *Users) FindAll() error {
	db, err := database.ConnectionDB()
	if err != nil {
		return err
	}
	db.Find(&u)
	return nil
}

func (u *User) FindByID(id int) error {
	db, err := database.ConnectionDB()
	if err != nil{
		return err
	}
	db.Find(&u, id)
	return nil
}

func (u *User)Create() error{
	db, err := database.ConnectionDB()
	if err != nil {
		return err
	}
	db.Create(&u)
	return nil
}