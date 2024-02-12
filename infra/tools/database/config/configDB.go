package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DatabaseConnectString = ""
	Port                  = 0
	SecretKey []byte // key used to sign token
)

func LoadInfos() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9090
	}

	fmt.Println("DB_USER:", os.Getenv("DB_USER"))
	fmt.Println("DB_SENHA:", os.Getenv("DB_SENHA"))
	fmt.Println("DB_NAME:", os.Getenv("DB_NAME"))

	DatabaseConnectString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NAME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
