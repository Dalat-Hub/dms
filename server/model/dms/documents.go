package dms

import (
	"database/sql"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type Documents struct {
	global.GVA_MODEL
	Title           string       `json:"title" form:"title" gorm:"column:title;"`
	ShortTitle      string       `json:"shortTitle" form:"shortTitle" gorm:"column:short_title;"`
	Expert          string       `json:"expert" form:"expert" gorm:"column:expert;"`
	Content         string       `json:"content" form:"content" gorm:"column:content;"`
	DateIssued      sql.NullTime `json:"dateIssued" form:"dateIssued" gorm:"column:date_issued;"`
	StillInEffect   bool         `json:"stillInEffect" form:"stillInEffect" gorm:"column:still_in_effect;"`
	EffectDate      sql.NullTime `json:"effectDate" form:"effectDate" gorm:"column:effect_date;"`
	ExpirationDate  sql.NullTime `json:"expirationDate" form:"expirationDate" gorm:"column:expiration_date;"`
	SignNumber      int          `json:"signNumber" form:"signNumber" gorm:"column:sign_number;size:10;"`
	SignYear        int          `json:"signYear" form:"signYear" gorm:"column:sign_year;comment:;size:10;"`
	SignCategory    string       `json:"signCategory" form:"signCategory" gorm:"column:sign_category;size:255;"`
	SignAgency      string       `json:"signAgency" form:"signAgency" gorm:"column:sign_agency;size:255;"`
	SignText        string       `json:"signText" form:"signText" gorm:"column:sign_text;size:255;"`
	CategoryId      uint         `json:"categoryId" form:"categoryId" gorm:"column:category_id;size:19;"`
	AgencyId        uint         `json:"agencyId" form:"agencyId" gorm:"column:agency_id;size:19;"`
	CreatedBy       uint         `json:"createdBy" form:"createdBy" gorm:"column:created_by;size:19;"`
	BeResponsibleBy uint         `json:"beResponsibleBy" form:"beResponsibleBy" gorm:"column:be_responsible_by;size:19;"`
	ViewCount       uint         `json:"viewCount" form:"viewCount" gorm:"column:view_count;size:19;"`
	DownloadCount   uint         `json:"downloadCount" form:"downloadCount" gorm:"column:download_count;size:19;"`
	Status          int          `json:"status" form:"status" gorm:"column:status;"`
	Type            int          `json:"type" form:"type" gorm:"column:type;"`
	Priority        int          `json:"priority" form:"priority" gorm:"column:priority_id;"`
	ParentId        uint         `json:"parentId" form:"parentId" gorm:"column:parent_id;size:10;"`
	CurrentId       uint         `json:"currentId" form:"currentId" gorm:"column:current_id;size:10;"`
	Path            string       `json:"path" form:"path" gorm:"column:path;"`
}

var (
	TYPE_DOCUMENT    int = 1
	TYPE_REVISION    int = 2
	STATUS_DRAFT     int = 1
	STATUS_PUBLISHED int = 2
	STATUS_INACTIVE  int = 3
	STATUS_TRASH     int = 4
	PRIORITY_NORMAL  int = 1
)

func (Documents) TableName() string {
	return "documents"
}
