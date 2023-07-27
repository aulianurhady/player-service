package api

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func StartWebServer() {
	app := fiber.New(fiber.Config{
		Prefork:   false,
		BodyLimit: 4 * 1024 * 1024,
	})

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, X-Request-With, Content-Type, Accept, nds-token",
		AllowMethods:     "HEAD, GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}))

	apiEndpoint := app.Group("/players")
	Load(apiEndpoint)

	err := app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		panic(err)
	}

}
