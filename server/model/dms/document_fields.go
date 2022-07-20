package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type DocumentFields struct {
	global.GVA_MODEL
	Name  string `json:"name" form:"name" gorm:"column:name;"`
	Count int    `json:"count" form:"count" gorm:"column:count;"`
}

func (DocumentFields) TableName() string {
	return "document_fields"
}
