package database

import (
	"time"

	"gorm.io/gorm"

	"github.com/niko0xdev/nx-ddns/internal/utils"
)

type BaseModel struct {
	ID        string    `gorm:"type:varchar(36);primary_key;" json:"id"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
}

func (model *BaseModel) BeforeCreate(tx *gorm.DB) error {

	if model.ID == "" {
		model.ID = utils.GenerateID().String()
	}

	return nil
}
