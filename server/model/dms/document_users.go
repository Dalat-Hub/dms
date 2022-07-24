package dms

import (
	"database/sql"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type DocumentUsers struct {
	global.GVA_MODEL
	DocumentId uint         `json:"documentId" form:"documentId" gorm:"column:document_id;size:19;"`
	UserId     uint         `json:"userId" form:"userId" gorm:"column:user_id;size:19;"`
	DoneAt     sql.NullTime `json:"doneAt" form:"doneAt" gorm:"column:done_at;"`
}

func (DocumentUsers) TableName() string {
	return "document_users"
}
