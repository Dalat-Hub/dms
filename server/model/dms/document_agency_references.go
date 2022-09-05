package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type DocumentAgencyReferences struct {
	global.GVA_MODEL
	AgencyId   uint `json:"agencyId" form:"agencyId" gorm:"column:agency_id;comment:;size:19;"`
	DocumentId uint `json:"documentId" form:"documentId" gorm:"column:document_id;comment:;size:19;"`
}

func (DocumentAgencyReferences) TableName() string {
	return "document_agency_references"
}
