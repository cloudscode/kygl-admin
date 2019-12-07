package model

type Codelibrary struct {
	Id        string `json:"id" xorm:"not null pk VARCHAR(16)"`
	Name      string `json:"name" xorm:"not null VARCHAR(300)"`
	Code      string `json:"code" xorm:"not null VARCHAR(50)"`
	Kind      string `json:"kind" xorm:"not null VARCHAR(16)"`
	Sortvalue int    `json:"sortvalue" xorm:"not null INT(11)"`
	Flagtree  int    `json:"flagtree" xorm:"not null INT(11)"`
}

func (Codelibrary) TableName() string {
	return "tab_codelibrary"
}
