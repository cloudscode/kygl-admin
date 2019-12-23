package model

import "time"

type Disquision struct {
	Id              int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Kind            string    `json:"kind" xorm:"not null VARCHAR(16)"`
	Title           string    `json:"title" xorm:"not null VARCHAR(256)"`
	Publications    string    `json:"publications" xorm:"not null VARCHAR(256)"`
	Levels          string    `json:"levels" xorm:"VARCHAR(16)"`
	Volumeno        string    `json:"volumeno" xorm:"VARCHAR(64)"`
	Period          string    `json:"period" xorm:"VARCHAR(64)"`
	Publishingdate  time.Time `json:"publishingdate" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Wordlength      float32   `json:"wordlength" xorm:"FLOAT"`
	Knowledgeclass  string    `json:"knowledgeclass" xorm:"VARCHAR(16)"`
	Firstknowledge  string    `json:"firstknowledge" xorm:"VARCHAR(16)"`
	Source          string    `json:"source" xorm:"VARCHAR(16)"`
	Publishingrange string    `json:"publishingrange" xorm:"VARCHAR(16)"`
	Issn            string    `json:"issn" xorm:"VARCHAR(64)"`
	Cn              string    `json:"cn" xorm:"VARCHAR(64)"`
	Authorname      string    `json:"authorname" xorm:"VARCHAR(256)"`
	Flagapprove     int       `json:"flagapprove" xorm:"INT(11)"`
	Flagschool      int       `json:"flagschool" xorm:"INT(11)"`
	Approvecomments string    `json:"approvecomments" xorm:"VARCHAR(500)"`
	Systemtype      int       `json:"systemtype" xorm:"INT(11)"`
	Disname         string    `json:"disname" xorm:"VARCHAR(100)"`
	Orgid           int       `json:"orgid" xorm:"INT(11)"`
	Deptid          int       `json:"deptid" xorm:"INT(11)"`
	Createid        int       `json:"createid" xorm:"INT(11)"`
	Createtime      time.Time `json:"createtime" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Updateid        int       `json:"updateid" xorm:"INT(11)"`
	Updatetime      time.Time `json:"updatetime" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

func (Disquision) TableName() string {
	return "tab_disquisition"
}

//for更新创建
type DisquisionEntity struct {
	Id              int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Kind            string    `json:"kind" xorm:"not null VARCHAR(16)"`
	Title           string    `json:"title" xorm:"not null VARCHAR(256)"`
	Publications    string    `json:"publications" xorm:"not null VARCHAR(256)"`
	Levels          string    `json:"levels" xorm:"VARCHAR(16)"`
	Volumeno        string    `json:"volumeno" xorm:"VARCHAR(64)"`
	Period          string    `json:"period" xorm:"VARCHAR(64)"`
	Publishingdate  time.Time `json:"publishingdate" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Wordlength      float32   `json:"wordlength" xorm:"FLOAT"`
	Knowledgeclass  string    `json:"knowledgeclass" xorm:"VARCHAR(16)"`
	Firstknowledge  string    `json:"firstknowledge" xorm:"VARCHAR(16)"`
	Source          string    `json:"source" xorm:"VARCHAR(16)"`
	Publishingrange string    `json:"publishingrange" xorm:"VARCHAR(16)"`
	Issn            string    `json:"issn" xorm:"VARCHAR(64)"`
	Cn              string    `json:"cn" xorm:"VARCHAR(64)"`
	Authorname      string    `json:"authorname" xorm:"VARCHAR(256)"`
	Flagapprove     int       `json:"flagapprove" xorm:"INT(11)"`
	Flagschool      int       `json:"flagschool" xorm:"INT(11)"`
	Approvecomments string    `json:"approvecomments" xorm:"VARCHAR(500)"`
	Systemtype      int       `json:"systemtype" xorm:"INT(11)"`
	Disname         string    `json:"disname" xorm:"VARCHAR(100)"`
	Orgid           int       `json:"orgid" xorm:"INT(11)"`
	Deptid          int       `json:"deptid" xorm:"INT(11)"`
	Updateid        int       `json:"updateid" xorm:"INT(11)"`
	Updatetime      time.Time `json:"updatetime" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
