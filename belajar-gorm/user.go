package main

import "time"

type Name struct{

	FirstName	string	`gorm:"column:first_name"`
	MiddleName	string	`gorm:"column:middle_name"`
	LastName	string	`gorm:"column:last_name"`

}

type User struct {
	ID        string 	`gorm:"primary_key;column:id"`
	Password  string 	`gorm:"column:password"`
	Name      Name 		`gorm:"embedded"`
	CreatedAt time.Time	`gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time	`gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

// membuat nama table manual
func (u *User)	TableName() string{
	return "users"
}