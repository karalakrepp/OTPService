package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Func for loading ENV
// Type the name of the variable you will use in .env
func envAccount(name string) string {
	println(godotenv.Unmarshal(".env"))

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln(err)
		log.Println("error loading .env ")

	}

	return os.Getenv(name)
}
