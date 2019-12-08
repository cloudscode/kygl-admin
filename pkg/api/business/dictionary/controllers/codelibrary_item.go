package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"zeus/pkg/api/business/dictionary/dto"
	"zeus/pkg/api/business/dictionary/service"
	baseController "zeus/pkg/api/controllers"
	baseDto "zeus/pkg/api/dto"
	baseService "zeus/pkg/api/service"
)

var codelibraryItemService = service.CodelibraryItemService{}
var logService = baseService.LogService{}

type CodelibraryItemController struct {
	baseController.BaseController
}

// @Summary 菜单列表
// @Tags menu
// @Security ApiKeyAuth
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{"result":[...],"total":1}}"
// @Router /v1/menus [get]
// List - r of crud
func (m *CodelibraryItemController) List(c *gin.Context) {
	var menuDto baseDto.GeneralTreeDto
	if m.BindAndValidate(c, &menuDto) {
		data, total := codelibraryItemService.List(menuDto)
		baseController.Resp(c, map[string]interface{}{
			"result": data,
			"total":  total,
		})
	}
}

// @Summary 菜单信息
// @Tags menu
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200 {array} model.User "{"code":200,"data":{"id":1,"name":"wutong"}}"
// @Failure 400 {string} json "{"code":10004,"msg": "用户信息不存在"}"
// @Router /v1/menus/:id [get]
func (u *CodelibraryItemController) QueryByCode(c *gin.Context) {
	var listDto dto.CodelibraryItemDto
	if u.BindAndValidate(c, &listDto) {
		data := codelibraryItemService.QueryByCode(listDto)
		baseController.Resp(c, map[string]interface{}{
			"result": data,
		})
	}
}

// @Summary 菜单信息
// @Tags menu
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200 {array} model.User "{"code":200,"data":{"id":1,"name":"wutong"}}"
// @Failure 400 {string} json "{"code":10004,"msg": "用户信息不存在"}"
// @Router /v1/menus/:id [get]
func (m *CodelibraryItemController) Get(c *gin.Context) {
	var gDto baseDto.GeneralGetDto
	if m.BindAndValidate(c, &gDto) {
		data := codelibraryItemService.InfoOfId(gDto)
		//user not found
		if data.Id < 1 {
			baseController.Fail(c, baseController.ErrNoRecord)
			return
		}
		baseController.Resp(c, map[string]interface{}{
			"result": data,
		})
	}
}

// @Tags Role
// @Summary 新增角色
// @Security ApiKeyAuth
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{"id":1}}"
// @Router /v1/roles [post]
func (m *CodelibraryItemController) Create(c *gin.Context) {
	var menuDto dto.CodelibraryItemCreateDto
	if m.BindAndValidate(c, &menuDto) {
		menu := codelibraryItemService.Create(menuDto)
		if menu.Id > 0 {
			// insert operation log
			b, _ := json.Marshal(menuDto)
			orLogDto := baseDto.OperationLogDto{
				UserId:           int(c.Value("userId").(float64)),
				RequestUrl:       c.Request.RequestURI,
				OperationMethod:  c.Request.Method,
				Params:           string(b),
				Ip:               c.ClientIP(),
				IpLocation:       "", //TODO...待接入获取ip位置服务
				OperationResult:  "success",
				OperationSuccess: 1,
				OperationContent: "Create Menu",
			}
			_ = logService.InsertOperationLog(&orLogDto)
		}
		baseController.Resp(c, map[string]interface{}{
			"result": menu,
		})
	}
}

// @Summary 编辑菜单
// @Tags menu
// @Security ApiKeyAuth
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{"id":1}}"
// @Router /v1/menus/:id [put]
// Edit - u of crud
func (u *CodelibraryItemController) Edit(c *gin.Context) {
	var menuDto dto.CodelibraryItemEditDto
	if u.BindAndValidate(c, &menuDto) {
		affected := codelibraryItemService.Update(menuDto)
		if affected > 0 {
			// insert operation log
			b, _ := json.Marshal(menuDto)
			orLogDto := baseDto.OperationLogDto{
				UserId:           int(c.Value("userId").(float64)),
				RequestUrl:       c.Request.RequestURI,
				OperationMethod:  c.Request.Method,
				Params:           string(b),
				Ip:               c.ClientIP(),
				IpLocation:       "", //TODO...待接入获取ip位置服务
				OperationResult:  "success",
				OperationSuccess: 1,
				OperationContent: "Edit Menu",
			}
			_ = logService.InsertOperationLog(&orLogDto)
		}
		baseController.Ok(c, "ok.UpdateDone")
	}
}

// @Summary 删除菜单
// @Tags menu
// @Security ApiKeyAuth
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{"id":1}}"
// @Router /v1/menus/:id [delete]
// Delete - d of crud
func (m *CodelibraryItemController) Delete(c *gin.Context) {
	var menuDto baseDto.GeneralDelDto
	if m.BindAndValidate(c, &menuDto) {
		affected := codelibraryItemService.Delete(menuDto)
		if affected <= 0 {
			if affected == -2 {
				baseController.Fail(c, baseController.ErrHasSubRecord)
			} else {
				baseController.Fail(c, baseController.ErrDelFail)
			}
			return
		}
		// insert operation log
		b, _ := json.Marshal(menuDto)
		orLogDto := baseDto.OperationLogDto{
			UserId:           int(c.Value("userId").(float64)),
			RequestUrl:       c.Request.RequestURI,
			OperationMethod:  c.Request.Method,
			Params:           string(b),
			Ip:               c.ClientIP(),
			IpLocation:       "", //TODO...待接入获取ip位置服务
			OperationResult:  "success",
			OperationSuccess: 1,
			OperationContent: "Delete Menu",
		}
		_ = logService.InsertOperationLog(&orLogDto)
		baseController.Ok(c, "ok.DeletedDone")
	}
}
