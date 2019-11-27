package model

type CodelibraryItem struct {
	Id            string `json:"id" xorm:"not null pk index VARCHAR(16)"`
	Name          string `json:"name" xorm:"index VARCHAR(150)"`
	Codelibraryid string `json:"codelibraryid" xorm:"not null pk VARCHAR(16)"`
	Parentid      string `json:"parentid" xorm:"index VARCHAR(16)"`
	Flagdefault   int    `json:"flagdefault" xorm:"not null INT(11)"`
	Flaglevel     int    `json:"flaglevel" xorm:"index INT(11)"`
	Sortvalue     int    `json:"sortvalue" xorm:"not null INT(11)"`
}

func (CodelibraryItem) TableName() string {
	return "tab_codelibrary_item"
}
