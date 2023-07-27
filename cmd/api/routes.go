package api

import (
	playerControllers "player-service/src/controllers/player"

	"github.com/gofiber/fiber/v2"
)

func Load(group fiber.Router) {
	PlayerRoute(group)
}

func PlayerRoute(api fiber.Router) {
	api.Get("/plain", playerControllers.PlainResponseController)

	api.Post("/register", playerControllers.CreatePlayerController)
	api.Get("/list", playerControllers.GetAllListPlayerController)
	api.Get("/list-filter", playerControllers.GetFilteredListPlayerController)
	api.Patch("/:id/topup", playerControllers.TopupSaldoPlayerController)
	api.Put("/:id/rekening-bank", playerControllers.UpdateRekeningBankPlayerController)
	api.Get("/:id/detail", playerControllers.GetDetailPlayerController)

	api.Post("/login", playerControllers.LoginAccountPlayerController)
	api.Put("/:id/logout", playerControllers.LoginAccountPlayerController)
}
