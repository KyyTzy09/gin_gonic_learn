package entity

import "time"

type UserEntity struct {
	UserId   uint    `gorm:"primaryKey" json:"user_id"`
	Email    string  `gorm:"uniqueIndex type:varchar(100); not null" json:"email"`
	Password *string `gorm:"type:varchar(100);default:null" json:"password"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
