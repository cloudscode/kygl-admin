package result

import (
	"fmt"
	"zeus/pkg/api/business/dictionary/dto"
	model "zeus/pkg/api/business/dictionary/model"
	baseDao "zeus/pkg/api/dao"
	baseDto "zeus/pkg/api/dto"
)

type CodelibraryItem struct {
}

var tabCodelibrary string = "tab_codelibrary"

//Get - get single CodelibraryItem info
func (u CodelibraryItem) Get(id int, preload bool) model.CodelibraryItem {
	var codelibraryItem model.CodelibraryItem
	db := baseDao.GetDb()
	if preload {
		db = db.Preload(tabCodelibrary)
	}
	db.Where("id = ?", id).First(&codelibraryItem)
	return codelibraryItem
}

// List - CodelibraryItem list
func (u CodelibraryItem) List(listDto baseDto.GeneralListDto) ([]model.CodelibraryItem, int64) {
	var codelibraryItem []model.CodelibraryItem
	var total int64
	db := baseDao.GetDb()
	for sk, sv := range baseDto.TransformSearch(listDto.Q, dto.CodelibraryItemListSearchMapping) {
		db = db.Where(fmt.Sprintf("%s = ?", sk), sv)
	}
	db.Preload(tabCodelibrary).Offset(listDto.Skip).Limit(listDto.Limit).Find(&codelibraryItem)
	db.Model(&model.CodelibraryItem{}).Count(&total)
	//fmt.Printf(">>>>>>>>>>>>" + codelibraryItem[0].Name + "<<<<<<<<<<<<<<<<")
	return codelibraryItem, total
}

// // GetRolesByIds
// func (Role) GetRolesByIds(ids string) []model.Role {
// 	var roles []model.Role
// 	db := GetDb()
// 	db.Where("id in (?)", strings.Split(ids, ",")).Find(&roles)
// 	return roles
// }

// // GetRolesByNames
// func (Role) GetRolesByNames(names []string) []model.Role {
// 	var roles []model.Role
// 	db := GetDb()
// 	db.Where("role_name in (?)", names).Find(&roles)
// 	return roles
// }

// //Get - get single roel infoD
// func (u Role) GetByName(name string) model.Role {
// 	var role model.Role
// 	db.Where("role_name = ?", name).Preload("Domain").First(&role)
// 	return role
// }

// // Create - new role
// func (r Role) Create(role *model.Role) *gorm.DB {
// 	var row model.Role
// 	db := GetDb()
// 	db.Where("name = ? or role_name = ?", role.Name, role.RoleName).First(&row)
// 	if row.Id > 0 {
// 		return nil
// 	}
// 	return db.Create(role)
// }

// // Update - update role
// func (r Role) Update(role *model.Role, ups map[string]interface{}) *gorm.DB {
// 	db := GetDb()
// 	return db.Model(role).Update(ups)
// }

// // Delete - delete role
// func (r Role) Delete(role *model.Role) *gorm.DB {
// 	db := GetDb()
// 	return db.Delete(role)
// }
