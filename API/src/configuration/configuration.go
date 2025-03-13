package configuration

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	ApiPort            = ""
	DbStringConnection = ""
)

func LoadConfigurations() {
	log.Println("Loading environment configurations")

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	ApiPort = os.Getenv("API_PORT")

	DbStringConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NOME"),
	)
}
