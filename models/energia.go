package models

import (
	"gorm.io/gorm"
	"time"
)

type Energia struct {
	ID             uint      `gorm:"primaryKey"`
	CreatedAt      time.Time `gorm:"<-:create"`
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	Data           time.Time      `json:"data"`
	Relogio1       int            `json:"relogio_1"`
	ValorConta1    float64        `json:"valor_conta_1" gorm:"-:all"`
	Relogio2       int            `json:"relogio_2"`
	ValorConta2    float64        `json:"valor_conta_2" gorm:"-:all"`
	Relogio3       int            `json:"relogio_3"`
	ValorConta3    float64        `json:"valor_conta_3" gorm:"-:all"`
	ValorKwh       float64        `json:"valor_kwh"`
	ValorConta     float64        `json:"valor_conta"`
	UltimoRegistro bool           `json:"ultimo_registro" gorm:"-:all"`
}
