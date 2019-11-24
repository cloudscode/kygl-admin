package result

import (
	"fmt"
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
