package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gopkg.in/guregu/null.v4"
)

type DocumentPermissionRequest struct {
	global.GVA_MODEL
	AcceptAt          null.Time `json:"acceptAt" form:"acceptAt" gorm:"column:accept_at;comment:;"`
	AcceptUserId      uint      `json:"acceptUserId" form:"acceptUserId" gorm:"column:accept_user_id;comment:;size:19;"`
	DocumentId        uint      `json:"documentId" form:"documentId" gorm:"column:document_id;comment:;size:19;"`
	RequestPermission string    `json:"requestPermission" form:"requestPermission" gorm:"column:request_permission;comment:;size:255;"`
	RequestUserId     uint      `json:"requestUserId" form:"requestUserId" gorm:"column:request_user_id;comment:;size:19;"`

	Document    *Documents      `json:"document" gorm:"foreignKey:DocumentId"`
	RequestUser *system.SysUser `json:"requestUser" gorm:"foreignKey:RequestUserId"`
	AcceptUser  *system.SysUser `json:"acceptUser" gorm:"foreignKey:AcceptUserId"`
}

func (DocumentPermissionRequest) TableName() string {
	return "document_permission_request"
}
