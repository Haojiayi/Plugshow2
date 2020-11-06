package common

import (
	"ShowWeb/models"
)

func GetPlugByMidLimit(idList []string, page int, size int) (plist []models.Plug, err error) {
	model := db

	for i := 0; i < len(idList); i++ {
		if i == 0 {
			model = model.Where("m_id=?", idList[0])
		} else {
			model = model.Or("m_id=?", idList[i])
		}
	}
	err = model.Limit(size).Offset((page - 1) * size).Find(&plist).Error
	if err != nil {
		return nil, err
	}
	return plist, nil
}

func GetPlugLimit(page int, size int) (plist []models.Plug, err error) {
	err = db.Limit(size).Offset((page - 1) * size).Find(&plist).Error
	if err != nil {
		return nil, err
	}
	return plist, nil
}

func GetPlugByMidCount(idList []string) (count int, err error) {
	model := db.Model(&models.Plug{})

	for i := 0; i < len(idList); i++ {
		if i == 0 {
			model = model.Where("m_id=?", idList[0])
		} else {
			model = model.Or("m_id=?", idList[i])
		}
	}

	err = model.Count(&count).Error
	return
}

func GetPlugCount() (count int, err error) {
	err = db.Model(&models.Plug{}).Count(&count).Error
	return
}

func GetPlugByMid(id string) (plist []models.Plug, err error) {
	err = db.Where("m_id= " + id).Find(&plist).Error
	if err != nil {
		return nil, err
	}
	return plist, nil
}
func GetPlugsByMidHot(id int) (plist []models.Plug, err error) {
	err = db.Where("m_id=?", id).Order("p_downum desc").Find(&plist).Error
	if err != nil {
		return nil, err
	}
	return plist, nil
}
func GetPlugsByPid(id int) (plist models.Plug, err error) {
	err = db.Where("id=?", id).Order("p_downum desc").Find(&plist).Error
	if err != nil {
		return plist, err
	}
	return plist, nil
}
func UpdateDByPid(p *models.Plug) (success string, err error) {
	err = db.Where("id=?", p.ID).Model(&p).Updates(&p).Error
	if err != nil {
		return "fail", err
	}
	return "success", nil
}

func GetPlugByKeyWord(keyword string) (plist []models.Plug, err error) {
	err = db.Where("p_title LIKE ?", keyword).Find(&plist).Error
	if err != nil {
		return nil, err
	}
	return plist, nil
}

func CreatePlug(plug *models.Plug) error {
	return db.Save(plug).Error
}

func GetPlugById(id string) (plug models.Plug, err error) {
	err = db.Find(&plug, id).Error
	return
}

func UpdatePlug(plug *models.Plug) error {
	return db.Save(plug).Error
}

func DeletePlug(id string) error {
	return db.Where("id=?", id).Delete(&models.Plug{}).Error
}
