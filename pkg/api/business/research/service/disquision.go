package service

import (
	"errors"
	"strconv"
	"zeus/pkg/api/business/research/dao"
	"zeus/pkg/api/business/research/dto"
	"zeus/pkg/api/business/research/model"
	"zeus/pkg/api/domain/user"
	baseDto "zeus/pkg/api/dto"
	"zeus/pkg/api/log"
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
	c := disquisionDao.Create(&userModel)
	if c.Error != nil {
		log.Error(c.Error.Error())
		return nil, errors.New("create user failed")
	}
	return &userModel, nil
}

// Update - update Disquision information
func (us DisquisionService) Update(userDto dto.DisquisionEditDto) int64 {
	userModel := model.Disquision{
		Id: userDto.Id,
	}
	c := disquisionDao.Update(&userModel, map[string]interface{}{})

	return c.RowsAffected
}

// Delete - delete Disquision
func (us DisquisionService) Delete(dto baseDto.GeneralDelDto) int64 {
	userModel := model.Disquision{
		Id: dto.Id,
	}
	c := disquisionDao.Delete(&userModel)
	if c.RowsAffected > 0 {
		user.DeleteUser(strconv.Itoa(dto.Id))
	}
	return c.RowsAffected
}
