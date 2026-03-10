package models

import "time"

// gorm is a Go ORM for SQL databases
type User struct {
	BaseModel
	UserID   int64      `gorm:"uniqueIndex;column:user_id;"`
	Email    string     `gorm:"index:idx_email;unique;type:varchar(100);not null;"`
	Password string     `gorm:"type:varchar(100);not null;"`
	UserName string     `gorm:"type:varchar(20);"`
	Birthday *time.Time `gorm:"type:datetime;"`
	Gender   string     `gorm:"default:male;type:varchar(6);comment 'female表示女，male表示男';"`
	Role     int32      `gorm:"default:1;type:int;comment '1表示普通用户，2表示管理员';"`
}
