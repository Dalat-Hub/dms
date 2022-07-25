package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gopkg.in/guregu/null.v4"
)

type DocumentUsers struct {
	global.GVA_MODEL
	DocumentId uint      `json:"documentId" form:"documentId" gorm:"column:document_id;size:19;"`
	Title      string    `json:"title" form:"title" gorm:"column:title;"`
	ShortTitle string    `json:"shortTitle" form:"shortTitle" gorm:"column:short_title;"`
	SignText   string    `json:"signText" form:"signText" gorm:"column:sign_text;"`
	UserId     uint      `json:"userId" form:"userId" gorm:"column:user_id;size:19;"`
	DoneAt     null.Time `json:"doneAt" form:"doneAt" gorm:"column:done_at;"`
}

func (DocumentUsers) TableName() string {
	return "document_users"
}
