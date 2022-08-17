package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms/response"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gopkg.in/guregu/null.v4"
)

type Documents struct {
	global.GVA_MODEL
	Title            string    `json:"title" form:"title" gorm:"column:title;"`
	ShortTitle       string    `json:"shortTitle" form:"shortTitle" gorm:"column:short_title;"`
	Expert           string    `json:"expert" form:"expert" gorm:"column:expert;"`
	Content          string    `json:"content" form:"content" gorm:"column:content;"`
	DateIssued       null.Time `json:"dateIssued" form:"dateIssued" gorm:"column:date_issued;"`
	StillInEffect    bool      `json:"stillInEffect" form:"stillInEffect" gorm:"column:still_in_effect;"`
	EffectDate       null.Time `json:"effectDate" form:"effectDate" gorm:"column:effect_date;"`
	ExpirationDate   null.Time `json:"expirationDate" form:"expirationDate" gorm:"column:expiration_date;"`
	SignNumber       int       `json:"signNumber" form:"signNumber" gorm:"column:sign_number;size:10;"`
	SignYear         int       `json:"signYear" form:"signYear" gorm:"column:sign_year;comment:;size:10;"`
	SignCategory     string    `json:"signCategory" form:"signCategory" gorm:"column:sign_category;size:255;"`
	SignAgency       string    `json:"signAgency" form:"signAgency" gorm:"column:sign_agency;size:255;"`
	SignText         string    `json:"signText" form:"signText" gorm:"column:sign_text;size:255;"`
	CategoryId       uint      `json:"categoryId" form:"categoryId" gorm:"column:category_id;size:19;"`
	AgencyId         uint      `json:"agencyId" form:"agencyId" gorm:"column:agency_id;size:19;"`
	CreatedBy        uint      `json:"createdBy" form:"createdBy" gorm:"column:created_by;size:19;"`
	UpdatedBy        uint      `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by"`
	BeResponsibleBy  uint      `json:"beResponsibleBy" form:"beResponsibleBy" gorm:"column:be_responsible_by;size:19;"`
	ViewCount        uint      `json:"viewCount" form:"viewCount" gorm:"column:view_count;size:19;"`
	DownloadCount    uint      `json:"downloadCount" form:"downloadCount" gorm:"column:download_count;size:19;"`
	Status           int       `json:"status" form:"status" gorm:"column:status;"`
	Type             int       `json:"type" form:"type" gorm:"column:type;"`
	Priority         int       `json:"priority" form:"priority" gorm:"column:priority;"`
	BelongTo         uint      `json:"belong_to" form:"belong_to" gorm:"column:belong_to;size:10;"`
	ParentId         uint      `json:"parentId" form:"parentId" gorm:"column:parent_id;size:10;"`
	CurrentId        uint      `json:"currentId" form:"currentId" gorm:"column:current_id;size:10;"`
	Path             string    `json:"path" form:"path" gorm:"column:path;"`
	PublicToView     bool      `json:"publicToView" form:"publicToView" gorm:"column:public_to_view;"`
	PublicToDownload bool      `json:"publicToDownload" form:"publicToDownload" gorm:"column:public_to_download"`

	Agency   *DocumentAgencies   `json:"agency" gorm:"foreignKey:AgencyId"`
	Category *DocumentCategories `json:"category" gorm:"foreignKey:CategoryId"`
	Fields   []*DocumentFields   `json:"fields" gorm:"many2many:document_field_references;foreignKey:ID;joinForeignKey:DocumentId;References:ID;joinReferences:FieldId;"`
	Signers  []*DocumentSigners  `json:"signers" gorm:"many2many:document_signer_references;foreignKey:ID;joinForeignKey:DocumentId;References:ID;joinReferences:SignerId;"`

	CreatedUser     *sysModel.SysUser `json:"createdUser" gorm:"-"`
	UpdatedUser     *sysModel.SysUser `json:"updatedUser" gorm:"-"`
	ResponsibleUser *sysModel.SysUser `json:"responsibleUser" gorm:"-"`

	BasedDocuments   []Documents        `json:"basedDocuments" gorm:"-"`
	RelatedDocuments []Documents        `json:"relatedDocuments" gorm:"-"`
	RelatedUsers     []sysModel.SysUser `json:"relatedUsers" gorm:"-"`
	RelatedAgencies  []DocumentAgencies `json:"relatedAgencies" gorm:"-"`

	Authority *response.DocumentAuthority `json:"authority" gorm:"-"`

	NeedLogin         bool `json:"needLogin" gorm:"-"`
	NeedAuthorization bool `json:"needAuthorization" gorm:"-"`
}

var (
	TYPE_DOCUMENT              int    = 1
	TYPE_REVISION              int    = 2
	STATUS_DRAFT               int    = 1
	STATUS_PUBLISHED           int    = 2
	STATUS_INACTIVE            int    = 3
	STATUS_TRASH               int    = 4
	PRIORITY_NORMAL            int    = 1
	RELATION_TYPE_USER         string = "user"
	RELATION_TYPE_AGENCY       string = "agency"
	RELATION_TYPE_DOC_BASE     string = "base"
	RELATION_TYPE_DOC_RELATION string = "relation"
)

func (Documents) TableName() string {
	return "documents"
}
