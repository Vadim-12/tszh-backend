package entity

import (
	"time"

	"github.com/google/uuid"
)

type Building struct {
	ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Address         string    `gorm:"not null"`
	Floors          int16     `gorm:"not null"`             // количество этажей
	CadastralNumber string    `gorm:"uniqueIndex;not null"` // кадастровый номер
	YearBuilt       int16     `gorm:"not null"`             // год постройки
	BuildingType    string    `gorm:"not null;size:32"`     // тип (кирпичный, монолит, панельный и т.п.)
	Entrances       int16     `gorm:"not null"`             // количество подъездов
	Apartments      *int16    ``                            // количество квартир
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`

	BuildingUnits []*BuildingUnit `gorm:"foreignKey:BuildingID"`
}
