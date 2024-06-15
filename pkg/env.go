package pkg

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv(file string) {
	err := godotenv.Load(file)
	if err != nil {
		fmt.Println("Please set the environment variables target file")
	}
}