package common

import (
	"ShowWeb/models"
)

func GetPlugByMidLimit(id string,page int,lastpage int) (plist []models.Plug,err error)  {
	err = db.Limit(page).Limit(lastpage).Where("m_id= "+id).Find(&plist).Error
	if err!=nil {
		return nil,err
	}
	return plist,nil
}
func GetPlugByMid (id string) (plist []models.Plug,err error)  {
	err = db.Where("m_id= "+id).Find(&plist).Error
	if err!=nil {
		return nil,err
	}
	return plist,nil
}
func GetPlugsByMidHot(id int) (plist []models.Plug,err error)  {
	err = db.Where("m_id=?",id).Order("p_downum desc").Find(&plist).Error
	if err!=nil {
		return nil,err
	}
	return plist,nil
}
func GetPlugsByPid(id int) (plist models.Plug,err error)  {
	err = db.Where("id=?",id).Order("p_downum desc").Find(&plist).Error
	if err!=nil {
		return plist,err
	}
	return plist,nil
}
func UpdateDByPid(p *models.Plug) ( success string,err error)  {
	err = db.Where("id=?",p.ID).Model(&p).Updates(&p).Error
	if err!=nil {
		return "fail",err
	}
	return "success",nil
}

func GetPlugByKeyWord(keyword string)(plist []models.Plug,err error)  {
	err = db.Where("p_title LIKE ?", keyword).Find(&plist).Error
	if err!=nil {
		return nil,err
	}
	return plist,nil
}

func CreatePlug(plug *models.Plug) error  {
	return db.Save(plug).Error
}

func GetPlugById(id string) (plug models.Plug,err error)  {
	err = db.Find(&plug,id).Error
	return
}

func UpdatePlug(plug *models.Plug) error  {
	return db.Save(plug).Error
}

func DeletePlug(id string) error  {
	return db.Where("id=?", id).Delete(&models.Plug{}).Error
}