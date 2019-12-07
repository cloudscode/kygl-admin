package dto

import (
	"gopkg.in/go-playground/validator.v8"
	"reflect"
	"strings"
)

// UserListSearchMapping - define search query keys in user list page
var CodelibraryItemListSearchMapping = map[string]string{
	"d": "parent_id",
}

// MenuCreateDto - dto for menu's creation
type CodelibraryItemCreateDto struct {
	Name      string `form:"name" json:"name" binding:"required"`
	Code      string `form:"code" json:"code"`
	ParentId  int    `form:"parent_id,default=0" json:"parent_id" binding:"gte=0"`
	Sortvalue int    `form:"sortvalue,default=1" json:"sortvalue"`
}

// MenuEditDto - dto for menu's modification
type CodelibraryItemEditDto struct {
	Id        int    `uri:"id" json:"id" binding:"required,gte=1"`
	Name      string `form:"name" json:"name" binding:"required"`
	Code      string `form:"code" json:"code"`
	ParentId  int    `form:"parent_id,default=0" json:"parent_id" binding:"gte=0"`
	Sortvalue int    `form:"sortvalue,default=1" json:"sortvalue"`
}

func permsValidate(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if val, ok := field.Interface().(string); ok {
		if strings.Contains(val, ",") {
			return false
		}
	}
	return true
}
