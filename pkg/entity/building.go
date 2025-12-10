package entity

import (
	"time"

	"github.com/google/uuid"
)

type Building struct {
	ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Address         string    `gorm:"not null"             json:"address"`
	Floors          int16     `gorm:"not null"             json:"floors"`           // количество этажей
	CadastralNumber string    `gorm:"uniqueIndex;not null" json:"cadastral_number"` // кадастровый номер
	YearBuilt       int16     `gorm:"not null"             json:"year_built"`       // год постройки
	BuildingType    string    `gorm:"not null;size:32"     json:"building_type"`    // тип (кирпичный, монолит, панельный и т.п.)
	Entrances       int16     `gorm:"not null"             json:"entrances"`        // количество подъездов
	Apartments      *int16    `json:"apartments,omitempty"`                         // количество квартир

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	BuildingUnits []*BuildingUnit `gorm:"foreignKey:BuildingID" json:"-"`
}

// 1) может ли пользователь быть в нескольких организациях? - да

// к одному строению может быть привязано несколько пользователей
// причем каждый пользователь привязывается к строению с определенной ролью (пока указываем просто роль, а собственник указывает серию, номер и дату выдачи свидетельства о праве собственности)
