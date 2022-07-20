package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type DocumentAgencies struct {
	global.GVA_MODEL
	Name  string `json:"name" form:"name" gorm:"column:name;"`
	Code  string `json:"code" form:"code" gorm:"column:code;"`
	Count int    `json:"count" form:"count" gorm:"column:count;"`
	Level string `json:"level" form:"level" gorm:"column:level;"`
}

func (DocumentAgencies) TableName() string {
	return "document_agencies"
}
