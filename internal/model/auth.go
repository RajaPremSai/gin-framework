package internal

import(

)

type User struct {
	Id int `gorm:"primaryKey" json:"id"`
	Password string `json:"password"`
	Email string `json:"email"`
}

func (u *User) TableName() string {
	return "users"
}