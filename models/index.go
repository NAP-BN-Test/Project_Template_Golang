package models

type Body_tblAccount struct {
	ID         int    `json:"id" form:"id"`
	Page       int    `json:"page" form:"page"`
	UserName   string `json:"userName" form:"userName"`
	Password   string `json:"password" form:"password"`
	Permission string `json:"permission" form:"permission"`
	Active     bool   `json:"active" form:"active"`
}
