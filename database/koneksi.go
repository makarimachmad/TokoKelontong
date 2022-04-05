package database

import (
	"fmt"
	"idnmedia/core"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDB()(*gorm.DB, error){
	dsn := core.SetupConfig("DB_USER") + ":" + core.SetupConfig("DB_PASS") + "@(" + core.SetupConfig("DB_HOST") + ":" + core.SetupConfig("DB_PORT") + ")/" + core.SetupConfig("DB_NAME")
	
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}
	fmt.Println("berhasil terhubung database")
	return db, nil
}