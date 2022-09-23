package db

import(
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

func EnvCouchBase(key string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return ""
	}
	return os.Getenv(key)
}