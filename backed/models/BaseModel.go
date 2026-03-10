package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

// BaseModel is the base structure for all models.
type BaseModel struct {
	ID        int64                 `gorm:"primarykey;type:int;"`
	CreatedAt time.Time             `gorm:"column:add_time"`
	UpdatedAt time.Time             `gorm:"column:update_time"`
	IsDeleted soft_delete.DeletedAt `gorm:"softDelete:flag;column:is_deleted"`
}
