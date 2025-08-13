package entity

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"gorm.io/gorm"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type User struct {
	ID         int64     `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"`
	Name       string    `gorm:"column:name;type:varchar;size:128;" json:"name"`
	Status     int32     `gorm:"column:status;type:int;" json:"status"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	return nil
}
