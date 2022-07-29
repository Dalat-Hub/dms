package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type DocumentSignerReferences struct {
	global.GVA_MODEL
	DocumentId uint `json:"documentId" form:"documentId" gorm:"column:document_id;"`
	SignerId   uint `json:"signerId" form:"signerId" gorm:"column:signer_id;"`
}

func (DocumentSignerReferences) TableName() string {
	return "document_signer_references"
}
