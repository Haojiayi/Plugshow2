package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Account 	string	`json:"account"`		//账号
	Password 	string	`json:"password"`		//密码
	Salt 		int		`json:"salt"`			//盐
}
