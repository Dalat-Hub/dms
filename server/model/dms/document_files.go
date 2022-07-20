package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type DocumentFiles struct {
	global.GVA_MODEL
	Name       string  `json:"name" form:"name" gorm:"column:name;"`
	Key        string  `json:"key" form:"key" gorm:"column:key;"`
	Url        string  `json:"url" form:"url" gorm:"column:url;"`
	Tag        string  `json:"tag" form:"tag" gorm:"column:tag;"`
	Size       float64 `json:"size" form:"size" gorm:"column:size;"`
	Order      int     `json:"order" form:"order" gorm:"column:order;"`
	DocumentId uint    `json:"documentId" form:"documentId" gorm:"column:document_id;"`
	Path       string  `json:"path" form:"path" gorm:"column:path;"`
}

func (DocumentFiles) TableName() string {
	return "document_files"
}
