package playerControllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"player-service/src/util/jwt"
	"player-service/src/util/response"

	"player-service/src/util/request"

	"github.com/gofiber/fiber/v2"

	FeaturePlayer "player-service/src/domain/player"
)

func PlainResponseController(c *fiber.Ctx) error {
	return response.RestSuccess(c, "ini player service", 200, nil)
}

func CreatePlayerController(c *fiber.Ctx) error {
	var req request.RequestCreatePlayer

	if err := c.BodyParser(&req); err != nil {
		return response.RestError(c, "Fail to parse body json", http.StatusBadRequest, err)
	}

	err := FeaturePlayer.PlayerService.CreatePlayerService(req)

	if err != nil {
		return response.RestError(c, "Internal Server Error", http.StatusInternalServerError, err)
	} else {
		return response.RestSuccess(c, "Player Data", http.StatusOK, req)
	}
}

func GetAllListPlayerController(c *fiber.Ctx) error {
	token := c.Get("token")

	if token == "" {
		return response.RestError(c, "Unauthorized", http.StatusUnauthorized, errors.New("akses tidak diijinkan, silahkan login terlebih dahulu"))
	}

	_, _, _, err := jwt.ClaimToken(token)
	if err != nil {
		return response.RestError(c, "Unauthorized", http.StatusUnauthorized, errors.New("token tidak valid"))
	}

	resp, err := FeaturePlayer.PlayerService.GetAllListPlayerService()
	if err != nil {
		return response.RestError(c, "Internal Server Error", http.StatusInternalServerError, err)
	}

	return response.RestSuccess(c, "Player Data", http.StatusOK, resp)
}

func GetFilteredListPlayerController(c *fiber.Ctx) error {
	token := c.Get("token")

	if token == "" {
		return response.RestError(c, "Unauthorized", http.StatusUnauthorized, errors.New("akses tidak diijinkan, silahkan login terlebih dahulu"))
	}

	_, _, _, err := jwt.ClaimToken(token)
	if err != nil {
		return response.RestError(c, "Unauthorized", http.StatusUnauthorized, errors.New("token tidak valid"))
	}

	var req request.RequestGetPlayer

	queries := c.Queries()

	d, _ := json.Marshal(queries)

	json.Unmarshal(d, &req)

	resp, err := FeaturePlayer.PlayerService.GetFilteredListPlayerService(req)
	if err != nil {
		return response.RestError(c, "Internal Server Error", http.StatusInternalServerError, err)
	}

	return response.RestSuccess(c, "Player Data", http.StatusOK, resp)
}

func TopupSaldoPlayerController(c *fiber.Ctx) error {
	token := c.Get("token")

	if token == "" {
		return response.RestError(c, "Unauthorized", http.StatusUnauthorized, errors.New("akses tidak diijinkan, silahkan login terlebih dahulu"))
	}

	_, _, _, err := jwt.ClaimToken(token)
	if err != nil {
		return response.RestError(c, "Unauthorized", http.StatusUnauthorized, errors.New("token tidak valid"))
	}

	var req request.RequestUpdateSaldoPlayer

	id := c.Params("id")
	if id == "" {
		return response.RestError(c, "id is empty", http.StatusBadRequest, errors.New("id not found"))
	}

	if err := c.BodyParser(&req); err != nil {
		return response.RestError(c, "Fail to parse body json", http.StatusBadRequest, err)
	}

	err = FeaturePlayer.PlayerService.TopupSaldoPlayerService(id, req)

	if err != nil {
		return response.RestError(c, "Internal Server Error", http.StatusInternalServerError, err)
	} else {
		return response.RestSuccess(c, "Player Data", http.StatusOK, req)
	}
}

func UpdateRekeningBankPlayerController(c *fiber.Ctx) error {
	token := c.Get("token")

	if token == "" {
		return response.RestError(c, "Unauthorized", http.StatusUnauthorized, errors.New("akses tidak diijinkan, silahkan login terlebih dahulu"))
	}

	_, _, _, err := jwt.ClaimToken(token)
	if err != nil {
		return response.RestError(c, "Unauthorized", http.StatusUnauthorized, errors.New("token tidak valid"))
	}

	var req request.RequestUpdateRekeningBankPlayer

	id := c.Params("id")
	if id == "" {
		return response.RestError(c, "id is empty", http.StatusBadRequest, errors.New("id not found"))
	}

	if err := c.BodyParser(&req); err != nil {
		return response.RestError(c, "Fail to parse body json", http.StatusBadRequest, err)
	}

	err = FeaturePlayer.PlayerService.UpdateRekeningBankPlayerService(id, req)

	if err != nil {
		return response.RestError(c, "Internal Server Error", http.StatusInternalServerError, err)
	} else {
		return response.RestSuccess(c, "Player Data", http.StatusInternalServerError, req)
	}
}

func GetDetailPlayerController(c *fiber.Ctx) error {
	token := c.Get("token")

	if token == "" {
		return response.RestError(c, "Unauthorized", http.StatusUnauthorized, errors.New("akses tidak diijinkan, silahkan login terlebih dahulu"))
	}

	_, _, _, err := jwt.ClaimToken(token)
	if err != nil {
		return response.RestError(c, "Unauthorized", http.StatusUnauthorized, errors.New("token tidak valid"))
	}

	id := c.Params("id")
	if id == "" {
		return response.RestError(c, "id is empty", http.StatusBadRequest, errors.New("id not found"))
	}

	resp, err := FeaturePlayer.PlayerService.GetDetailPlayerService(id)
	if err != nil {
		return response.RestError(c, "Internal Server Error", http.StatusInternalServerError, err)
	}

	return response.RestSuccess(c, "Player Data", http.StatusOK, resp)
}

func LoginAccountPlayerController(c *fiber.Ctx) error {
	var req request.RequestLoginAccountPlayer

	if err := c.BodyParser(&req); err != nil {
		return response.RestError(c, "Fail to parse body json", http.StatusBadRequest, err)
	}

	token, err := FeaturePlayer.PlayerService.LoginAccountService(req)

	if err != nil {
		return response.RestError(c, "Internal Server Error", http.StatusInternalServerError, err)
	} else {
		return response.RestSuccess(c, "Player Data", http.StatusOK, token)
	}
}

func LogoutAccountPlayerController(c *fiber.Ctx) error {
	token := c.Get("token")

	if token == "" {
		return response.RestError(c, "Unauthorized", http.StatusUnauthorized, errors.New("akses tidak diijinkan, silahkan login terlebih dahulu"))
	}

	_, _, _, err := jwt.ClaimToken(token)
	if err != nil {
		return response.RestError(c, "Unauthorized", http.StatusUnauthorized, errors.New("token tidak valid"))
	}

	id := c.Params("id")
	if id == "" {
		return response.RestError(c, "id is empty", http.StatusBadRequest, errors.New("id not found"))
	}

	err = FeaturePlayer.PlayerService.LogoutAccountService(id)

	if err != nil {
		return response.RestError(c, "Internal Server Error", http.StatusInternalServerError, err)
	} else {
		return response.RestSuccess(c, "Player Data", http.StatusOK, "logout sukses")
	}
}
