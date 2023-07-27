package model

import (
	"time"

	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Username      string    `db:"username" json:"username"`
	Password      string    `db:"password" json:"password"`
	NamaRekening  string    `db:"nama_rekening" json:"nama_rekening"`
	NomorRekening string    `db:"nomor_rekening" json:"nomor_rekening"`
	NamaBank      string    `db:"nama_bank" json:"nama_bank"`
	SisaSaldo     int64     `db:"sisa_saldo" json:"sisa_saldo"`
	RegisterAt    time.Time `db:"register_at" json:"register_at"`
}
