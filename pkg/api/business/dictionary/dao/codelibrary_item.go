package result

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
	"zeus/pkg/api/business/dictionary/dto"
	model "zeus/pkg/api/business/dictionary/model"
	baseDao "zeus/pkg/api/dao"
	baseDto "zeus/pkg/api/dto"
)

type CodelibraryItem struct {
}

// GetMenusByIds
func (m CodelibraryItem) GetMenusByIds(ids string) []model.CodelibraryItem {
	var menus []model.CodelibraryItem
	db := baseDao.GetDb()
	db.Where("id in (?) and menu_type=1", strings.Split(ids, ",")).Find(&menus)
	return menus
}

// GetMenusPermByIds - get permissions in menu table
func (m CodelibraryItem) GetMenusPermByIds(ids string) []model.CodelibraryItem {
	var menus []model.CodelibraryItem
	db := baseDao.GetDb()
	db.Where("id in (?) and menu_type=2", strings.Split(ids, ",")).Find(&menus)
	return menus
}

// GetByAlias - get row by alias
func (m CodelibraryItem) GetByAlias(alias string) model.CodelibraryItem {
	var menu model.CodelibraryItem
	db := baseDao.GetDb()
	db.Where("alias = ? and menu_type=2", alias).First(&menu)
	return menu
}

//Get - get single menu info
func (m CodelibraryItem) Get(id int, preload bool) model.CodelibraryItem {
	var menu model.CodelibraryItem
	db := baseDao.GetDb()
	if preload {
		db = db.Preload("tab_codelibrary_item")
	}
	db.Where("id = ?", id).First(&menu)
	return menu
}

//GetSubMenus
func (m CodelibraryItem) GetSubMenus(id int) []model.CodelibraryItem {
	var menus []model.CodelibraryItem
	db := baseDao.GetDb()
	db.Where("parent_id=?", id).First(&menus)
	return menus
}

// List
func (m CodelibraryItem) List(treeDto baseDto.GeneralTreeDto) ([]model.CodelibraryItem, int64) {
	var menus []model.CodelibraryItem
	var total int64
	db := baseDao.GetDb()
	for sk, sv := range baseDto.TransformSearch(treeDto.Q, dto.CodelibraryItemListSearchMapping) {
		db = db.Where(fmt.Sprintf("%s = ?", sk), sv)
	}
	db.Preload("tab_codelibrary_item").Order(" sortvalue asc").Find(&menus)
	db.Model(&model.CodelibraryItem{}).Count(&total)
	return menus, total
}

// queryByCode
func (m CodelibraryItem) QueryByCode(listDto dto.CodelibraryItemDto) []model.CodelibraryItem {
	var menus []model.CodelibraryItem

	db := baseDao.GetDb()
	// sub := db.Table("tab_codelibrary_item").Select("ID")
	// for sk, sv := range baseDto.TransformSearch(listDto.Q, dto.CodelibraryItemQueryByCodeMapping) {
	// 	db.Where(fmt.Sprintf("%s = ?", sk), sv)
	// }
	fmt.Println("******************************" + listDto.Q)
	sub := db.Table("tab_codelibrary_item").Select("ID").Where("code = ?", listDto.Q).SubQuery()
	db.Table("tab_codelibrary_item").Where("parent_id IN ?", sub).Find(&menus)
	return menus
}

// Create - new menu
func (m CodelibraryItem) Create(menu *model.CodelibraryItem) *gorm.DB {
	db := baseDao.GetDb()
	return db.Save(menu)
}

// Update - update menu
func (m CodelibraryItem) Update(menu *model.CodelibraryItem, ups map[string]interface{}) *gorm.DB {
	db := baseDao.GetDb()
	return db.Model(menu).Update(ups)
}

// Delete - delete menu
func (m CodelibraryItem) Delete(menu *model.CodelibraryItem) *gorm.DB {
	db := baseDao.GetDb()
	return db.Delete(menu)
}
