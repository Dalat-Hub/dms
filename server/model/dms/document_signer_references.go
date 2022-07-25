package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type DocumentSignerReferences struct {
	global.GVA_MODEL
	DocumentId uint `json:"documentId" form:"documentId" gorm:"column:document_id;"`
	UserId     uint `json:"userId" form:"userId" gorm:"column:user_id;"`
}

func (DocumentSignerReferences) TableName() string {
	return "document_signer_references"
}
