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
	Kind            string    `json:"kind"`
	Publications    string    `json:"publications"`
	Levels          string    `json:"levels"`
	Volumeno        string    `json:"volumeno"`
	Period          string    `json:"period"`
	Publishingdate  time.Time `json:"publishingdate"`
	Wordlength      float32   `json:"wordlength"`
	Knowledgeclass  string    `json:"knowledgeclass"`
	Firstknowledge  string    `json:"firstknowledge"`
	Source          string    `json:"source"`
	Publishingrange string    `json:"publishingrange"`
	Issn            string    `json:"issn"`
	Cn              string    `json:"cn"`
	Orgid           int       `json:"orgid"`
	Authorname      string    `json:"authorname"`
	Flagapprove     int       `json:"flagapprove"`
	Sysemplid       int       `json:"sysemplid"`
	Sysdatetime     time.Time `json:"sysdatetime"`
	Flagschool      int       `json:"flagschool"`
	Approvecomments string    `json:"approvecomments"`
	Systemtype      int       `json:"systemtype"`
	Disname         string    `json:"disname"`
}

// DisquisionEditDto - dto for disquision modification
type DisquisionEditDto struct {
	Id              int       `uri:"id" json:"id" binding:"required,gte=1"`
	Title           string    `form:"title" json:"title" binding:"required"`
	Kind            string    `json:"kind"`
	Publications    string    `json:"publications"`
	Levels          string    `json:"levels"`
	Volumeno        string    `json:"volumeno"`
	Period          string    `json:"period"`
	Publishingdate  time.Time `json:"publishingdate"`
	Wordlength      float32   `json:"wordlength"`
	Knowledgeclass  string    `json:"knowledgeclass"`
	Firstknowledge  string    `json:"firstknowledge"`
	Source          string    `json:"source"`
	Publishingrange string    `json:"publishingrange"`
	Issn            string    `json:"issn"`
	Cn              string    `json:"cn"`
	Orgid           int       `json:"orgid"`
	Authorname      string    `json:"authorname"`
	Flagapprove     int       `json:"flagapprove"`
	Sysemplid       int       `json:"sysemplid"`
	Sysdatetime     time.Time `json:"sysdatetime"`
	Flagschool      int       `json:"flagschool"`
	Approvecomments string    `json:"approvecomments"`
	Systemtype      int       `json:"systemtype"`
	Disname         string    `json:"disname"`
}
