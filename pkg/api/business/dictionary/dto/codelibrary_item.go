package dto

// CodelibraryItemListSearchMapping - define search query keys in CodelibraryItem list page
var CodelibraryItemListSearchMapping = map[string]string{
	"n": "name",
	"d": "domain_id",
	"r": "role_name",
}

// CodelibraryItemCreateDto - dto for CodelibraryItem creation
type CodelibraryItemCreateDto struct {
	Name string `form:"name" json:"name" binding:"required"`
	// DomainId    int    `form:"domain_id" json:"domain_id" binding:"required,gte=1"`
	// RoleName    string `form:"role_name" json:"role_name" binding:"required"`
	// Remark      string `form:"remark" json:"remark"`
	// MenuIds     string `form:"menu_ids" json:"menu_ids"`
	// MenuIdsEle  string `form:"menu_ids_ele" json:"menu_ids_ele"`
	// DataPermIds string `form:"data_perm_ids" json:"data_perm_ids"`
}

// CodelibraryItemEditDto - dto for CodelibraryItem modification
type CodelibraryItemEditDto struct {
	Id   int    `uri:"id" json:"id" binding:"required,gte=1"`
	Name string `form:"name" json:"name" binding:"required"`
	// DomainId    int    `form:"domain_id" json:"domain_id" binding:"required,gte=1"`
	// Remark      string `form:"remark" json:"remark"`
	// MenuIds     string `form:"menu_ids" json:"menu_ids"`
	// MenuIdsEle  string `form:"menu_ids_ele" json:"menu_ids_ele"`
	// DataPermIds string `form:"data_perm_ids" json:"data_perm_ids"`
}
