package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"zeus/pkg/api/business/research/dto"
	"zeus/pkg/api/business/research/model"
	baseDao "zeus/pkg/api/dao"
	baseDto "zeus/pkg/api/dto"
)

type Disquision struct {
}

var tableDisquisition string = "tab_disquisition"

//Get - get single disquision info
func (u Disquision) Get(id int, preload bool) model.Disquision {
	var disquision model.Disquision
	db := baseDao.GetDb()
	if preload {
		db = db.Preload(tableDisquisition)
	}
	db.Where("id = ?", id).First(&disquision)
	return disquision
}

// List - Disquision list
func (u Disquision) List(listDto baseDto.GeneralListDto) ([]model.Disquision, int64) {
	var disquision []model.Disquision
	var total int64
	db := baseDao.GetDb()
	for sk, sv := range baseDto.TransformSearch(listDto.Q, dto.DisquisionListSearchMapping) {
		db = db.Where(fmt.Sprintf("%s = ?", sk), sv)
	}
	db.Preload(tableDisquisition).Offset(listDto.Skip).Limit(listDto.Limit).Find(&disquision)
	db.Model(&model.Disquision{}).Count(&total)
	fmt.Printf(">>>>>>>>>>>>" + disquision[0].Title + "<<<<<<<<<<<<<<<<")
	return disquision, total
}

// Create - new Disquision
func (u Disquision) Create(disquision *model.Disquision) *gorm.DB {
	db := baseDao.GetDb()
	return db.Create(disquision)
}

// Update - update Disquision
func (u Disquision) Update(disquision *model.Disquision, ups map[string]interface{}) *gorm.DB {
	db := baseDao.GetDb()
	return db.Model(disquision).Update(ups)
}

// Delete - delete Disquision
func (u Disquision) Delete(disquision *model.Disquision) *gorm.DB {
	db := baseDao.GetDb()
	return db.Delete(disquision)
}
