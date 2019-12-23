package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"zeus/pkg/api/business/research/dto"
	"zeus/pkg/api/business/research/service"
	baseController "zeus/pkg/api/controllers"
	baseDto "zeus/pkg/api/dto"
	"zeus/pkg/api/log"
	baseService "zeus/pkg/api/service"
)

var disquisionService = service.DisquisionService{}
var logService = baseService.LogService{}

type DisquisionController struct {
	baseController.BaseController
}

// @Tags Role
// @Summary 角色列表[分页+搜索]
// @Security ApiKeyAuth
// @Param limit query int false "条数"
// @Param skip query int false "偏移量"
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{"result":[...],"total":1}}"
// @Router /v1/roles [get]
func (r *DisquisionController) List(c *gin.Context) {
	var listDto baseDto.GeneralListDto
	if r.BindAndValidate(c, &listDto) {
		data, total := disquisionService.List(listDto)
		baseController.Resp(c, map[string]interface{}{
			"result": data,
			"total":  total,
		})
	}
}

// @Tags Users
// @Summary 新增用户
// @Security ApiKeyAuth
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{"id":1}}"
// @Router /v1/users?limit=20&offset=0 [get]
func (u *DisquisionController) Create(c *gin.Context) {
	var userDto dto.DisquisionCreateDto
	log.Info(fmt.Sprintf("%#v", c))
	if u.BindAndValidate(c, &userDto) {
		log.Info(fmt.Sprintf("%#v", userDto))
		user, err := disquisionService.Create(userDto)
		if err != nil {
			baseController.Fail(c, baseController.ErrInputData)
			return
		}
		// TODO insert ldap user
		// insert operation log
		b, _ := json.Marshal(userDto)
		orLogDto := baseDto.OperationLogDto{
			UserId:           int(c.Value("userId").(float64)),
			RequestUrl:       c.Request.RequestURI,
			OperationMethod:  c.Request.Method,
			Params:           string(b),
			Ip:               c.ClientIP(),
			IpLocation:       "", //TODO...待接入获取ip位置服务
			OperationResult:  "success",
			OperationSuccess: 1,
			OperationContent: "Create User",
		}
		_ = logService.InsertOperationLog(&orLogDto)
		baseController.Resp(c, map[string]interface{}{
			"result": user,
		})
	}
}

// @Tags Users
// @Summary 编辑用户
// @Security ApiKeyAuth
// @Produce  json
// @Param id path int true "用户id"
// @Success 200 {string} json "{"code":200,"data":{"id":1}}"
// @Router /v1/users/{id} [put]
func (u *DisquisionController) Edit(c *gin.Context) {
	var userDto dto.DisquisionEditDto
	if u.BindAndValidate(c, &userDto) {
		log.Info(fmt.Sprintf("%#v", userDto))
		affected := disquisionService.Update(userDto)
		// TODO update ldap user
		if affected > 0 {
			// insert operation log
			b, _ := json.Marshal(userDto)
			orLogDto := baseDto.OperationLogDto{
				UserId:           int(c.Value("userId").(float64)),
				RequestUrl:       c.Request.RequestURI,
				OperationMethod:  c.Request.Method,
				Params:           string(b),
				Ip:               c.ClientIP(),
				IpLocation:       "", //TODO...待接入获取ip位置服务
				OperationResult:  "success",
				OperationSuccess: 1,
				OperationContent: "Edit User",
			}
			_ = logService.InsertOperationLog(&orLogDto)
		}
		baseController.Ok(c, "ok.UpdateDone")
	}
}

// @Tags Users
// @Summary 删除用户
// @Security ApiKeyAuth
// @Param id path int true "用户id"
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{"id":1}}"
// @Router /v1/users/{id} [delete]
func (u *DisquisionController) Delete(c *gin.Context) {
	var userDto baseDto.GeneralDelDto
	if u.BindAndValidate(c, &userDto) {
		affected := disquisionService.Delete(userDto)
		if affected <= 0 {
			baseController.Fail(c, baseController.ErrDelFail)
			return
		}
		// insert operation log
		b, _ := json.Marshal(userDto)
		orLogDto := baseDto.OperationLogDto{
			UserId:           int(c.Value("userId").(float64)),
			RequestUrl:       c.Request.RequestURI,
			OperationMethod:  c.Request.Method,
			Params:           string(b),
			Ip:               c.ClientIP(),
			IpLocation:       "", //TODO...待接入获取ip位置服务
			OperationResult:  "success",
			OperationSuccess: 1,
			OperationContent: "Delete User",
		}
		_ = logService.InsertOperationLog(&orLogDto)
		baseController.Ok(c, "ok.DeletedDone")
	}
}

func (u *DisquisionController) Get(c *gin.Context) {
	var gDto baseDto.GeneralGetDto
	if u.BindAndValidate(c, &gDto) {
		data := disquisionService.InfoOfId(gDto)
		//user not found
		if data.Id < 1 {
			baseController.Fail(c, baseController.ErrNoUser)
			return
		}
		baseController.Resp(c, map[string]interface{}{
			"result": data,
		})
	}
}
