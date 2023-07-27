package player

import (
	"context"
	"errors"
	"player-service/infrastructure/redis"
	"player-service/src/domain/player/model"
	"player-service/src/domain/player/repository"
	"player-service/src/util/helper"
	"player-service/src/util/jwt"
	"player-service/src/util/request"
	"reflect"
	"strconv"
)

var PlayerService playerService

type playerService struct {
}

func (ps *playerService) CreatePlayerService(request request.RequestCreatePlayer) error {
	err := repository.PlayerRepo.CreatePlayerRepo(request)
	if err != nil {
		return err
	}

	return nil
}

func (ps *playerService) GetAllListPlayerService() ([]model.Player, error) {
	res, _ := repository.PlayerRepo.GetAllRepo()
	return res, nil
}

func (ps *playerService) GetFilteredListPlayerService(req request.RequestGetPlayer) ([]model.Player, error) {
	res, _ := repository.PlayerRepo.GetListByFilterRepo(req)
	return res, nil
}

func (ps *playerService) TopupSaldoPlayerService(id string, request request.RequestUpdateSaldoPlayer) error {
	uId, _ := strconv.ParseUint(id, 0, 64)
	err := repository.PlayerRepo.UpdateSaldoRepo(uint(uId), request)
	if err != nil {
		return err
	}

	return nil
}

func (ps *playerService) UpdateRekeningBankPlayerService(id string, request request.RequestUpdateRekeningBankPlayer) error {
	uId, _ := strconv.ParseUint(id, 0, 64)
	err := repository.PlayerRepo.UpdateRekeningBankRepo(uint(uId), request)
	if err != nil {
		return err
	}

	return nil
}

func (ps *playerService) GetDetailPlayerService(id string) (model.Player, error) {
	uId, _ := strconv.ParseUint(id, 0, 64)
	res, err := repository.PlayerRepo.GetDetailPlayerRepo(uint(uId))
	if err != nil {
		return model.Player{}, err
	}

	if reflect.DeepEqual(res, model.Player{}) {
		return model.Player{}, errors.New("data player tidak ditemukan")
	}

	return res, nil
}

func (ps *playerService) LoginAccountService(request request.RequestLoginAccountPlayer) (string, error) {
	dataPlayer, err := repository.PlayerRepo.LoginAccountRepo(request)
	if err != nil {
		return "", err
	}

	if reflect.DeepEqual(dataPlayer, model.Player{}) {
		return "", errors.New("data player tidak ditemukan")
	}

	tokenString := jwt.CreateClaimJwt(dataPlayer.ID, dataPlayer.Username, dataPlayer.RegisterAt.String())

	key := helper.GenerateRedisKey(dataPlayer.ID)

	redis.RedisClient.Set(context.Background(), key, tokenString, 0)

	return tokenString, nil
}

func (ps *playerService) LogoutAccountService(id string) error {
	uId, _ := strconv.ParseUint(id, 0, 64)
	key := helper.GenerateRedisKey(uint(uId))

	data, err := redis.RedisClient.Get(context.Background(), key).Result()
	if err != nil {
		return err
	}

	if data == "" {
		return errors.New("tidak ada data login")
	}

	redis.RedisClient.Del(context.Background(), key)

	return nil
}
