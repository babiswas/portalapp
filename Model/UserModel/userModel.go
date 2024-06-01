package UserModel

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName  string    `json:"username"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"password"`
	RoleID    uint      `gorm:"not null;DEFAULT:3" json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Role      Role      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
