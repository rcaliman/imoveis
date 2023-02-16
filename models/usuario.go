package models

import (
	"gorm.io/gorm"
	"time"
)

type Usuario struct {
	gorm.Model
	CreatedAt time.Time `gorm:"<-:create"`
	Usuario   string    `json:"usuario" gorm:"unique"`
	Senha     string    `json:"senha"`
	Tipo      string    `json:"tipo"`
}
