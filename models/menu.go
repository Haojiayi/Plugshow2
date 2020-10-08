package models

import "github.com/jinzhu/gorm"

type Menu struct {
	gorm.Model
	Name 	string	`json:"name" form:"name" binding:"required"`		//菜单名称
	M_l_id 	int		`json:"m_l_id" form:"m_l_id" binding:"required"`		//上级菜单id
	M_url 	string	`json:"m_url" form:"m_url" binding:"required"`		//菜单地址
}