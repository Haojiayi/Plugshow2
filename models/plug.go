package models

import "github.com/jinzhu/gorm"

type Plug struct {
	gorm.Model
	P_title 		string		`json:"p_title"`		//插件名称
	P_content 		string		`json:"p_content"`		//插件描述
	P_v_dress 		string		`json:"p_v_dress"`		//视频地址
	P_downum 		int			`json:"p_downum"`		//下载次数
	P_img_adress 	string		`json:"p_img_adress"`	//图片地址
	M_id 			int			`json:"m_id"`			//菜单id
	P_d_adress 		string 		`json:"p_d_adress"`		//下载地址
}