package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type DocumentRules struct {
	global.GVA_MODEL
	DocumentId uint   `json:"documentId" form:"documentId" gorm:"column:document_id;comment:;size:19;"`
	Permission string `json:"permission" form:"permission" gorm:"column:permission;comment:;size:255;"`
	SubjectId  uint   `json:"subjectId" form:"subjectId" gorm:"column:subject_id;comment:;size:19;"`
	Type       string `json:"type" form:"type" gorm:"column:type;comment:;size:255;"`
}

func (DocumentRules) TableName() string {
	return "document_rules"
}

var (
	RULE_VIEW_ALL       string = "all"
	RULE_VIEW_LIMIT     string = "limit"
	RULE_DOWNLOAD_ALL   string = "all"
	RULE_DOWNLOAD_LIMIT string = "limit"
	RULE_EDIT_ONLY      string = "only"
	RULE_EDIT_LIMIT     string = "limit"
	RULE_OWNER_ONLY     string = "only"
	RULE_OWNER_LIMIT    string = "limit"

	PERMISSION_OWNER    string = "owner"
	PERMISSION_EDIT     string = "edit"
	PERMISSION_VIEW     string = "view"
	PERMISSION_DOWNLOAD string = "download"

	RULE_TYPE_USER  string = "user"
	RULE_TYPE_GROUP string = "group"
)
