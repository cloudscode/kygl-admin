package controllers

import (
	"github.com/gin-gonic/gin"
	"zeus/pkg/api/business/research/service"
	baseController "zeus/pkg/api/controllers"
	"zeus/pkg/api/dto"
)

var disquisionService = service.DisquisionService{}

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
	var listDto dto.GeneralListDto
	if r.BindAndValidate(c, &listDto) {
		data, total := disquisionService.List(listDto)
		baseController.Resp(c, map[string]interface{}{
			"result": data,
			"total":  total,
		})
	}
}

// // @Tags Role
// // @Summary 新增角色
// // @Security ApiKeyAuth
// // @Produce  json
// // @Success 200 {string} json "{"code":200,"data":{"id":1}}"
// // @Router /v1/roles [post]
// func (r *RoleController) Create(c *gin.Context) {
// 	var roleDto dto.RoleCreateDto
// 	if r.BindAndValidate(c, &roleDto) {
// 		newRole, err := roleService.Create(roleDto)
// 		if err != nil {
// 			ErrAddFail.Moreinfo = err.Error()
// 			fail(c, ErrAddFail)
// 			ErrAddFail.Moreinfo = ""
// 			return
// 		}
// 		// insert operation log
// 		b, _ := json.Marshal(roleDto)
// 		orLogDto := dto.OperationLogDto{
// 			UserId:           int(c.Value("userId").(float64)),
// 			RequestUrl:       c.Request.RequestURI,
// 			OperationMethod:  c.Request.Method,
// 			Params:           string(b),
// 			Ip:               c.ClientIP(),
// 			IpLocation:       "", //TODO...待接入获取ip位置服务
// 			OperationResult:  "success",
// 			OperationSuccess: 1,
// 			OperationContent: "Create Role",
// 		}
// 		_ = logService.InsertOperationLog(&orLogDto)
// 		resp(c, map[string]interface{}{
// 			"result": newRole,
// 		})
// 	}
// }

// // @Summary 更新角色信息
// // @Tags Role
// // @Security ApiKeyAuth
// // @Produce  json
// // @Success 200 {string} json "{"code":200,"data":{"result":[...],"total":1}}"
// // @Router /v1/roles/:id [put]
// // Edit - u of crud
// func (r *RoleController) Edit(c *gin.Context) {
// 	var roleDto dto.RoleEditDto
// 	if r.BindAndValidate(c, &roleDto) {
// 		affected := roleService.Update(roleDto)
// 		if affected < 0 {
// 			fail(c, ErrNoRecord)
// 			return
// 		}
// 		// insert operation log
// 		b, _ := json.Marshal(roleDto)
// 		orLogDto := dto.OperationLogDto{
// 			UserId:           int(c.Value("userId").(float64)),
// 			RequestUrl:       c.Request.RequestURI,
// 			OperationMethod:  c.Request.Method,
// 			Params:           string(b),
// 			Ip:               c.ClientIP(),
// 			IpLocation:       "", //TODO...待接入获取ip位置服务
// 			OperationResult:  "success",
// 			OperationSuccess: 1,
// 			OperationContent: "Edit Role",
// 		}
// 		_ = logService.InsertOperationLog(&orLogDto)
// 		ok(c, "ok.UpdateDone")
// 	}
// }

// // @Summary 删除角色信息
// // @Tags Role
// // @Security ApiKeyAuth
// // @Produce  json
// // @Success 200 {string} json "{"code":200,"data":{"result":[...],"total":1}}"
// // @Router /v1/roles/:id [delete]
// // Delete - d of crud
// func (r *RoleController) Delete(c *gin.Context) {
// 	var roleDto dto.GeneralDelDto
// 	if r.BindAndValidate(c, &roleDto) {
// 		if roleService.Delete(roleDto) < 1 {
// 			fail(c, ErrNoRecord)
// 			return
// 		}
// 		// insert operation log
// 		b, _ := json.Marshal(roleDto)
// 		orLogDto := dto.OperationLogDto{
// 			UserId:           int(c.Value("userId").(float64)),
// 			RequestUrl:       c.Request.RequestURI,
// 			OperationMethod:  c.Request.Method,
// 			Params:           string(b),
// 			Ip:               c.ClientIP(),
// 			IpLocation:       "", //TODO...待接入获取ip位置服务
// 			OperationResult:  "success",
// 			OperationSuccess: 1,
// 			OperationContent: "Delete Role",
// 		}
// 		_ = logService.InsertOperationLog(&orLogDto)
// 		ok(c, "ok.DeleteDone")
// 	}
// }
