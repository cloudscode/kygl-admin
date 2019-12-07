package service

import (
	"time"
	dao "zeus/pkg/api/business/dictionary/dao"
	"zeus/pkg/api/business/dictionary/dto"
	model "zeus/pkg/api/business/dictionary/model"
	baseDto "zeus/pkg/api/dto"
	"zeus/pkg/api/log"
)

var codelibraryItemDao = dao.CodelibraryItem{}

// CodelibraryItemService
type CodelibraryItemService struct {
}

// InfoOfId - get menu info by id
func (CodelibraryItemService) InfoOfId(dto baseDto.GeneralGetDto) model.CodelibraryItem {
	return codelibraryItemDao.Get(dto.Id, true)
}

// List - users list with pagination
func (CodelibraryItemService) List(treeDto baseDto.GeneralTreeDto) ([]model.CodelibraryItem, int64) {
	return codelibraryItemDao.List(treeDto)
}

// Create - create a menu item
func (ms CodelibraryItemService) Create(codelibraryItemCreateDto dto.CodelibraryItemCreateDto) model.CodelibraryItem {
	codelibraryItemCreateModel := model.CodelibraryItem{
		Name:           codelibraryItemCreateDto.Name,
		Code:           codelibraryItemCreateDto.Code,
		ParentId:       codelibraryItemCreateDto.ParentId,
		Sortvalue:      codelibraryItemCreateDto.Sortvalue,
		CreateTime:     time.Now(),
		LastUpdateTime: time.Now(),
		Deleted:        0,
	}
	c := codelibraryItemDao.Create(&codelibraryItemCreateModel)
	if c.Error != nil {
		log.Error(c.Error.Error())
	}
	// for _, alias := range strings.Split(codelibraryItemCreateDto.Alias, ",") {
	// 	menuPermAliasDao.Create(&model.MenuPermAlias{
	// 		Perms:       codelibraryItemCreateDto.Perms,
	// 		Alias:       alias,
	// 		DomainId:    codelibraryItemCreateDto.DomainId,
	// 		CreatedTime: time.Now().Unix(),
	// 		UpdatedTime: time.Now().Unix(),
	// 	})
	// }
	return codelibraryItemCreateModel
}

// Update
func (ms CodelibraryItemService) Update(codelibraryItemCreateDto dto.CodelibraryItemEditDto) int64 {
	codelibraryItemCreateModel := codelibraryItemDao.Get(codelibraryItemCreateDto.Id, false)
	c := codelibraryItemDao.Update(&codelibraryItemCreateModel, map[string]interface{}{
		"name":      codelibraryItemCreateDto.Name,
		"code":      codelibraryItemCreateDto.Code,
		"parent_id": codelibraryItemCreateDto.ParentId,
	})
	// 1.Remove all alias
	//menuPermAliasDao.Delete(model.MenuPermAlias{Perms: codelibraryItemCreateModel.Perms, DomainId: codelibraryItemCreateModel.DomainId})
	// 2.Save new alias again
	// for _, alias := range strings.Split(codelibraryItemCreateDto.Alias, ",") {
	// 	menuPermAliasDao.Create(&model.MenuPermAlias{
	// 		Perms:       codelibraryItemCreateDto.Perms,
	// 		Alias:       alias,
	// 		DomainId:    codelibraryItemCreateDto.DomainId,
	// 		CreatedTime: time.Now().Unix(),
	// 		UpdatedTime: time.Now().Unix(),
	// 	})
	// }
	return c.RowsAffected
}

// Delete - delete menu
func (ms CodelibraryItemService) Delete(dto baseDto.GeneralDelDto) int64 {
	codelibraryItemModel := model.CodelibraryItem{
		Id: dto.Id,
	}
	c := codelibraryItemDao.Delete(&codelibraryItemModel)
	return c.RowsAffected
	// codelibraryItemCreateModel := codelibraryItemDao.Get(dto.Id, false)
	// if codelibraryItemCreateModel.Id < 1 {
	// 	return -1
	// }
	// //If menu has sub menus , stop removing node
	// // subMenus := codelibraryItemDao.GetSubMenus(dto.Id)
	// // if len(subMenus) > 0 {
	// // 	return -2
	// // }
	// // //1. delete menu
	// // perms := codelibraryItemCreateModel.Perms
	// // c := codelibraryItemDao.Delete(&codelibraryItemCreateModel)
	// // if c.RowsAffected > 0 {
	// // 	//2. delete related policies
	// // 	//Empty perms puts in delete method would cause all casbin rules removed!
	// // 	if perms != "" {
	// // 		perm.DelFilteredPerm(1, perms)
	// // 	}
	// // 	// Remove all alias
	// // 	menuPermAliasDao.Delete(model.MenuPermAlias{Perms: codelibraryItemCreateModel.Perms, DomainId: codelibraryItemCreateModel.DomainId})
	// // }
	// //return c.RowsAffected
	// return 0
}
