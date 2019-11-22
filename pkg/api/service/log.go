package service

import (
	"zeus/pkg/api/dao"
	"zeus/pkg/api/dto"
	"zeus/pkg/api/model"
)

type LoginLog = model.LoginLog
type OperationLog = model.OperationLog

var loginLogDao = dao.LoginLogDao{}
var operationLogDao = dao.OperationLogDao{}

type LogService struct {
}

func (LogService) LoginLogDetail(dto dto.GeneralGetDto) LoginLog {
	return loginLogDao.Detail(dto.Id)
}

// List - users list with pagination
func (LogService) LoginLogLists(dto dto.LoginLogListDto) ([]dao.LoginLogList, int64) {
	return loginLogDao.Lists(dto)
}

func (LogService) OperationLogDetail(dto dto.GeneralGetDto) OperationLog {
	return operationLogDao.Detail(dto.Id)
}

// List - users list with pagination
func (LogService) OperationLogLists(dto dto.OperationLogListDto) ([]dao.OperationLogList, int64) {
	return operationLogDao.Lists(dto)
}

//Insert Operation Log
func (LogService) InsertOperationLog(orLogDto *dto.OperationLogDto) error {
	return operationLogDao.Create(orLogDto)
}
