package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type DocumentRelationReferences struct {
	global.GVA_MODEL
	DocumentId   uint   `json:"documentId" form:"documentId" gorm:"column:document_id"`
	DestId       uint   `json:"destId" form:"destId" gorm:"column:dest_id"`
	RelationType string `json:"relationType" form:"relationType" gorm:"column:relation_type"`
}

func (DocumentRelationReferences) TableName() string {
	return "document_relation_references"
}
