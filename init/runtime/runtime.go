package runtime

import (
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(err)
	}
}
