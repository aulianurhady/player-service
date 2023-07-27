package request

import "time"

type RequestCreatePlayer struct {
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	NamaRekening  string    `json:"nama_rekening"`
	NomorRekening string    `json:"nomor_rekening"`
	NamaBank      string    `json:"nama_bank"`
	SisaSaldo     int64     `json:"sisa_saldo"`
	RegisterAt    time.Time `json:"register_at"`
}

type RequestGetPlayer struct {
	Username      string    `json:"username"`
	NamaRekening  string    `json:"nama_rekening"`
	NomorRekening string    `json:"nomor_rekening"`
	NamaBank      string    `json:"nama_bank"`
	SisaSaldo     int64     `json:"sisa_saldo"`
	RegisterAt    time.Time `json:"register_at"`
}

type RequestUpdateSaldoPlayer struct {
	SisaSaldo int64 `json:"sisa_saldo"`
}

type RequestUpdateRekeningBankPlayer struct {
	NamaRekening  string `json:"nama_rekening"`
	NomorRekening string `json:"nomor_rekening"`
	NamaBank      string `json:"nama_bank"`
}

type RequestLoginAccountPlayer struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
