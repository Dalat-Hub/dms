// 自动生成模板DocumentSigners
package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// DocumentSigners 结构体
type DocumentSigners struct {
	global.GVA_MODEL
	AgencyId int    `json:"agencyId" form:"agencyId" gorm:"column:agency_id;comment:;size:19;"`
	Count    int    `json:"count" form:"count" gorm:"column:count;comment:;size:10;"`
	Fullname string `json:"fullname" form:"fullname" gorm:"column:fullname;comment:;"`
	Title    int    `json:"title" form:"title" gorm:"column:title;comment:;size:10;"`
}

// TableName DocumentSigners 表名
func (DocumentSigners) TableName() string {
	return "document_signers"
}
