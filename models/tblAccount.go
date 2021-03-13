package models

// bắt buộc phải viết hoa
type TBLAccount struct {
	ID         int64  `gorm:"primary_key;auto_increment;not_null;column:ID" json:"id" form:"id"`
	UserName   string `gorm:"column:UserName" json:"userName" form:"userName"`
	Password   string `gorm:"column:Password" json:"password" form:"password"`
	Permission string `gorm:"column:Permission" json:"permission" form:"permission"`
	Active     bool   `gorm:"column:Active" json:"active" form:"active"`
}

type TBLAccounts []TBLAccount

func (TBLAccount) TableName() string {
	return "tblAccount"
}
