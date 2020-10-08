package common

import (
	"ShowWeb/models"
)

func GetAllMenu () (mlist []models.Menu,err error)  {
	err = db.Find(&mlist).Error
	if err!=nil {
		return nil,err
	}
	return mlist,nil

}
func GetAllFatherMenu()(mlist []models.Menu,err error)  {
	err = db.Where("m_l_id=0").Find(&mlist).Error
	if err!=nil {
		return nil,err
	}
	return mlist,nil
}
func GetChildrenMenuByFid(id int) (mlist []models.Menu,err error)  {
	err = db.Where("m_l_id=?",id).Find(&mlist).Error
	if err!=nil {
		return nil,err
	}
	return mlist,nil

}
func GetMenuByMid(id int)(mlist []models.Menu,err error)    {
	err = db.Where("id=?",id).Find(&mlist).Error
	if err!=nil {
		return nil,err
	}
	return mlist,nil
}
func CreateMenu(menu *models.Menu) error  {
	return db.Save(menu).Error
}
func UpdateMenu(menu models.Menu) error  {
	return db.Save(menu).Error
}
func GetMenuById(id string) (menu models.Menu,err error)  {
	err = db.Find(&menu,id).Error
	return
}
func DeleteMenu(id string) error  {
	return db.Where("id=?", id).Delete(&models.Menu{}).Error
}
