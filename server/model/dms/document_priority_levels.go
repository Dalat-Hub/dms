package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type DocumentPriorityLevels struct {
	global.GVA_MODEL
	Level int    `json:"level" form:"level" gorm:"column:level;"`
	Name  string `json:"name" form:"name" gorm:"column:name;"`
}

func (DocumentPriorityLevels) TableName() string {
	return "document_priority_levels"
}
