package runtime

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if gin.Mode() == gin.DebugMode {
		err := godotenv.Load("./.env")
		if err != nil {
			log.Fatal(err)
		}
	}
}
