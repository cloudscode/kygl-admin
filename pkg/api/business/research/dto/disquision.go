package dto

import "time"

// DisquisionListSearchMapping - define search query keys in disquision list page
var DisquisionListSearchMapping = map[string]string{
	"n": "name",
	"d": "domain_id",
	"r": "role_name",
}

// DisquisionCreateDto - dto for disquision creation
type DisquisionCreateDto struct {
	Title           string    `form:"title" json:"title" binding:"required"`
	Kind            string    `form:"kind" json:"kind"`
	Publications    string    `form:"publications" json:"publications"`
	Levels          string    `form:"levels" json:"levels"`
	Volumeno        string    `form:"volumeno" json:"volumeno"`
	Period          string    `form:"period" json:"period"`
	Publishingdate  time.Time `form:"publishingdate" json:"publishingdate"`
	Wordlength      float32   `form:"wordlength" json:"wordlength"`
	Knowledgeclass  string    `form:"knowledgeclass" json:"knowledgeclass"`
	Firstknowledge  string    `form:"firstknowledge" json:"firstknowledge"`
	Source          string    `form:"source" json:"source"`
	Publishingrange string    `form:"publishingrange" json:"publishingrange"`
	Issn            string    `form:"issn" json:"issn"`
	Cn              string    `form:"cn" json:"cn"`
	Authorname      string    `form:"authorname" json:"authorname"`
	Flagapprove     int       `form:"flagapprove" json:"flagapprove"`
	Flagschool      int       `form:"flagschool" json:"flagschool"`
	Approvecomments string    `form:"approvecomments" json:"approvecomments"`
	Systemtype      int       `form:"systemtype" json:"systemtype"`
	Disname         string    `form:"disname" json:"disname"`
	Orgid           int       `form:"orgid" json:"orgid"`
	Deptid          int       `form:"deptid" json:"deptid"`
	Createid        int       `form:"createid" json:"createid"`
	Createtime      time.Time `form:"createtime" json:"createtime"`
	Updateid        int       `form:"updateid" json:"updateid"`
	Updatetime      time.Time `form:"updatetime" json:"updatetime"`
}

// DisquisionEditDto - dto for disquision modification
type DisquisionEditDto struct {
	Id              int       `form:"id" json:"id"`
	Title           string    `form:"title" json:"title"`
	Kind            string    `form:"kind" json:"kind"`
	Publications    string    `form:"publications" json:"publications"`
	Levels          string    `form:"levels" json:"levels"`
	Volumeno        string    `form:"volumeno" json:"volumeno"`
	Period          string    `form:"period" json:"period"`
	Publishingdate  time.Time `form:"publishingdate" json:"publishingdate"`
	Wordlength      float32   `form:"wordlength" json:"wordlength"`
	Knowledgeclass  string    `form:"knowledgeclass" json:"knowledgeclass"`
	Firstknowledge  string    `form:"firstknowledge" json:"firstknowledge"`
	Source          string    `form:"source" json:"source"`
	Publishingrange string    `form:"publishingrange" json:"publishingrange"`
	Issn            string    `form:"issn" json:"issn"`
	Cn              string    `form:"cn" json:"cn"`
	Authorname      string    `form:"authorname" json:"authorname"`
	Flagapprove     int       `form:"flagapprove" json:"flagapprove"`
	Flagschool      int       `form:"flagschool" json:"flagschool"`
	Approvecomments string    `form:"approvecomments" json:"approvecomments"`
	Systemtype      int       `form:"systemtype" json:"systemtype"`
	Disname         string    `form:"disname" json:"disname"`
	Orgid           int       `form:"orgid" json:"orgid"`
	Deptid          int       `form:"deptid" json:"deptid"`
	Updateid        int       `form:"updateid" json:"updateid"`
	Updatetime      time.Time `form:"updatetime" json:"updatetime"`
}
