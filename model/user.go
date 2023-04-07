package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	name    string  `gorm:"type:varchar(32);not null;unique;comment:用户名"`
	balance float64 `gorm:"type:decimal(10,2);default:0;comment:余额"`
}

func (u *User) AddBalance(amount float64) {
	u.balance += amount
}

func (u *User) Balance() float64 {
	return u.balance
}

func (u *User) SetBalance(balance float64) {
	u.balance = balance
}

func (u *User) Name() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}
