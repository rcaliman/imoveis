package models

import (
	"gorm.io/gorm"
	"time"
)

type Cliente struct {
	ID             uint `gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	Nome           string         `json:"nome" gorm:"not null"`
	Locacoes       string         `json:"locacoes" gorm:"-:all"`
	DataNascimento time.Time      `json:"data_nascimento"`
	Ci             string         `json:"ci"`
	Cpf            string         `json:"cpf"`
	Telefone1      string         `json:"telefone_1"`
	Telefone2      string         `json:"telefone_2"`
}
