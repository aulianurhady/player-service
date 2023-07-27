package repository

import (
	"errors"
	"player-service/infrastructure/database"
	"player-service/src/domain/player/model"
	"player-service/src/util/request"
	"time"
)

var PlayerRepo playerRepo

type playerRepo struct {
}

func (pr *playerRepo) CreatePlayerRepo(req request.RequestCreatePlayer) error {
	data := model.Player{
		Username:      req.Username,
		NamaRekening:  req.NamaRekening,
		NomorRekening: req.NomorRekening,
		NamaBank:      req.NamaBank,
		SisaSaldo:     req.SisaSaldo,
		RegisterAt:    time.Now(),
	}

	result := database.DB.Create(&data)
	if result.RowsAffected <= 0 {
		return errors.New("gagal membuat player")
	}

	return nil
}

func (pr *playerRepo) GetAllRepo() ([]model.Player, error) {
	var data []model.Player
	database.DB.Find(&data)

	if len(data) == 0 {
		return data, errors.New("data kosong")
	}

	return data, nil
}

func (pr *playerRepo) GetListByFilterRepo(req request.RequestGetPlayer) ([]model.Player, error) {
	var data []model.Player

	database.DB.Where(model.Player{
		Username:      req.Username,
		NamaRekening:  req.NamaRekening,
		NamaBank:      req.NamaBank,
		SisaSaldo:     req.SisaSaldo,
		RegisterAt:    req.RegisterAt,
		NomorRekening: req.NomorRekening,
	}).Find(&data)

	return data, nil
}

func (pr *playerRepo) UpdateSaldoRepo(id uint, req request.RequestUpdateSaldoPlayer) error {
	result := database.DB.Model(&model.Player{}).Where("id = ?", id).Update("sisa_saldo", req.SisaSaldo)
	if result.RowsAffected <= 0 {
		return errors.New("gagal topup saldo player")
	}
	return nil
}

func (pr *playerRepo) UpdateRekeningBankRepo(id uint, req request.RequestUpdateRekeningBankPlayer) error {
	var data model.Player
	data.ID = id

	result := database.DB.Model(&data).Where("id = ?", id).Updates(map[string]interface{}{"nama_rekening": req.NamaRekening, "nomor_rekening": req.NomorRekening, "nama_bank": req.NamaBank})
	if result.RowsAffected <= 0 {
		return errors.New("gagal update data rekening player")
	}
	return nil
}

func (pr *playerRepo) GetDetailPlayerRepo(id uint) (model.Player, error) {
	var data model.Player
	result := database.DB.First(&data, id)

	if result.RowsAffected <= 0 {
		return model.Player{}, errors.New("gagal mendapatkan detail player")
	}

	return data, nil
}

func (pr *playerRepo) LoginAccountRepo(request request.RequestLoginAccountPlayer) (model.Player, error) {
	var data model.Player
	result := database.DB.Where("username = ? AND password = ?", request.Username, request.Password).First(&data)

	if result.RowsAffected <= 0 {
		return model.Player{}, errors.New("gagal mendapatkan data login player")
	}

	return data, nil
}
