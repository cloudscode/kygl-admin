package service

import (
	"errors"
	"fmt"
	"zeus/pkg/api/business/research/dao"
	"zeus/pkg/api/business/research/dto"
	"zeus/pkg/api/business/research/model"
	baseDto "zeus/pkg/api/dto"
	"zeus/pkg/api/log"
	"zeus/pkg/api/utils/copier"
)

var disquisionDao = dao.Disquision{}

// DisquisionService
type DisquisionService struct {
}

// List - users list with pagination
func (DisquisionService) List(dto baseDto.GeneralListDto) ([]model.Disquision, int64) {
	return disquisionDao.List(dto)
}

// Create - create a new Disquision
func (us DisquisionService) Create(userDto dto.DisquisionCreateDto) (*model.Disquision, error) {

	userModel := model.Disquision{}
	copier.Copy(&userModel, &userDto)
	log.Info(fmt.Sprintf("create > %#v", &userModel))
	c := disquisionDao.Create(&userModel)
	if c.Error != nil {
		log.Error(c.Error.Error())
		return nil, errors.New("create user failed")
	}
	return &userModel, nil
}

// Update - update Disquision information
func (us DisquisionService) Update(userDto dto.DisquisionEditDto) int64 {
	// userModel := model.Disquision{
	// 	Id: userDto.Id,
	// }
	userModel := model.Disquision{}
	copier.Copy(&userModel, &userDto)
	log.Info(fmt.Sprintf("Update > %#v", &userModel))
	c := disquisionDao.Update(&userModel, map[string]interface{}{})

	return c.RowsAffected
}

// Delete - delete Disquision
func (us DisquisionService) Delete(dto baseDto.GeneralDelDto) int64 {
	userModel := model.Disquision{
		Id: dto.Id,
	}
	c := disquisionDao.Delete(&userModel)

	return c.RowsAffected
}

func (us DisquisionService) InfoOfId(dto baseDto.GeneralGetDto) model.Disquision {
	return disquisionDao.Get(dto.Id, true)
}
