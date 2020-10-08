package models

import "github.com/jinzhu/gorm"

type DownHistory struct {
	gorm.Model
	user_id 	string	`json:"user_id"`		//下载用户
}
