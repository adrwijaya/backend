package models

type Users struct {
	Id       int    `json:"id"`
	FullName string `json:"fullName" gorm:"type: varchar(255)"`
	UserName string `json:"userName" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
}
