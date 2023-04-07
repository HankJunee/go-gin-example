package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string  `gorm:"type:varchar(32);not null;unique;comment:用户名"`
	Balance float64 `gorm:"type:decimal(10,2);default:0;comment:余额"`
}

func (u *User) AddBalance(amount float64) {
	u.Balance += amount
}
