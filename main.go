package main

import (
	"fmt"
	"os"
	"player-service/cmd/api"
	"player-service/infrastructure/database"
	"player-service/infrastructure/redis"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func main() {
	err := database.SetupDatabase()
	if err != nil {
		panic(err)
	}

	redis.NewRedisClient(os.Getenv("REDIS_HOST"), "")

	fmt.Println("webserver running")

	api.StartWebServer()
}
