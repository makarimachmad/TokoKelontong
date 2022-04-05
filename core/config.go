package core

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)	

func SetupConfig(key string) string {
	err := godotenv.Load(".env")
	if err != nil{
		fmt.Println("gagal ambil value di env")
	}
	return os.Getenv(key)
}