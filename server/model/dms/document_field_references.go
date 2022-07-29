package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type DocumentFieldReferences struct {
	global.GVA_MODEL
	FieldId    uint `json:"fieldId" form:"fieldId" gorm:"column:field_id"`
	DocumentId uint `json:"documentId" form:"documentId" gorm:"column:document_id;"`
}

func (DocumentFieldReferences) TableName() string {
	return "document_field_references"
}
