package entity

import "time"

type UserEntity struct {
	UserId   uint   `gorm:"primaryKey" json:"user_id"`
	Email    string `gorm:"type:varchar(100); not null" json:"email"`
	Password string `gorm:"type:varchar(100); optional" json:"password"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
