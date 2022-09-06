package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type DocumentCategories struct {
	global.GVA_MODEL
	Name  string `json:"name" form:"name" gorm:"column:name;"`
	Code  string `json:"code" form:"code" gorm:"column:code;"`
	Count int    `json:"count" form:"count" gorm:"column:count;"`

	Order  int   `json:"order" form:"order" gorm:"order"`
	Hidden *bool `json:"hidden" form:"hidden" gorm:"hidden"`
}

func (DocumentCategories) TableName() string {
	return "document_categories"
}
