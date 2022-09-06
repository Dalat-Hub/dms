package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
	dmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/dms/request"
)

type DocumentAgenciesService struct {
}

// CreateDocumentAgencies create agency
func (documentAgenciesService *DocumentAgenciesService) CreateDocumentAgencies(documentAgencies *dms.DocumentAgencies) (err error) {
	err = global.GVA_DB.Create(&documentAgencies).Error
	return err
}

// DeleteDocumentAgencies delete agency by ID
func (documentAgenciesService *DocumentAgenciesService) DeleteDocumentAgencies(documentAgencies dms.DocumentAgencies) (err error) {
	err = global.GVA_DB.Delete(&documentAgencies).Error
	return err
}

// DeleteDocumentAgenciesByIds bulk delete agency by IDs
func (documentAgenciesService *DocumentAgenciesService) DeleteDocumentAgenciesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dms.DocumentAgencies{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateDocumentAgencies update agency
func (documentAgenciesService *DocumentAgenciesService) UpdateDocumentAgencies(documentAgencies dms.DocumentAgencies) (err error) {
	err = global.GVA_DB.Save(&documentAgencies).Error
	return err
}

// GetDocumentAgencies find agency by ID
func (documentAgenciesService *DocumentAgenciesService) GetDocumentAgencies(id uint) (documentAgencies dms.DocumentAgencies, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&documentAgencies).Error
	return
}

// GetDocumentAgenciesInfoList get list of agencies
func (documentAgenciesService *DocumentAgenciesService) GetDocumentAgenciesInfoList(info dmsReq.DocumentAgenciesSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// create db object
	db := global.GVA_DB.Model(&dms.DocumentAgencies{})

	var documentAgenciess []*dms.DocumentAgencies

	// If there is a conditional search, the search statement will be automatically created below
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&documentAgenciess).Error
	if err != nil {
		return
	}

	err = documentAgenciesService.attachDocumentCount(documentAgenciess)
	if err != nil {
		return
	}

	return documentAgenciess, total, err
}

// GetDocumentAgenciesInfoListPublic get list of agencies
func (documentAgenciesService *DocumentAgenciesService) GetDocumentAgenciesInfoListPublic(info dmsReq.DocumentAgenciesSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// create db object
	db := global.GVA_DB.Model(&dms.DocumentAgencies{})

	db = db.Where("`hidden` = ?", false)
	db = db.Order("`order` desc")

	var documentAgenciess []*dms.DocumentAgencies

	// If there is a conditional search, the search statement will be automatically created below
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&documentAgenciess).Error
	if err != nil {
		return
	}

	err = documentAgenciesService.attachDocumentCount(documentAgenciess)
	if err != nil {
		return
	}

	return documentAgenciess, total, err
}

type agencyTreeNode struct {
	AgencyId   uint                    `json:"agencyId" gorm:"agency_id"`
	CategoryId uint                    `json:"categoryId" gorm:"category_id"`
	Agency     *dms.DocumentAgencies   `json:"agency" gorm:"-"`
	Category   *dms.DocumentCategories `json:"category" gorm:"-"`
	Count      uint                    `json:"count" gorm:"count"`
}

func (documentAgenciesService *DocumentAgenciesService) GetAgencyTree() (list interface{}, err error) {
	var results []*agencyTreeNode

	err = global.GVA_DB.Model(&dms.Documents{}).Debug().
		Select("document_agency_references.agency_id as agency_id, documents.category_id as category_id, count(documents.ID) as count").
		Where("type = ? AND status = ?", dms.TYPE_DOCUMENT, dms.STATUS_PUBLISHED).
		Joins("JOIN document_agency_references ON documents.id = document_agency_references.document_id").
		Group("document_agency_references.agency_id, documents.category_id").
		Find(&results).Error

	if err != nil {
		return nil, err
	}

	// attach data
	var agencies []*dms.DocumentAgencies
	var categories []*dms.DocumentCategories

	agenciesMap := map[uint]*dms.DocumentAgencies{}
	categoriesMap := map[uint]*dms.DocumentCategories{}

	err = global.GVA_DB.Model(&dms.DocumentAgencies{}).Find(&agencies).Error
	if err != nil {
		return nil, err
	}

	err = global.GVA_DB.Model(&dms.DocumentCategories{}).Find(&categories).Error
	if err != nil {
		return nil, err
	}

	for _, v := range agencies {
		agenciesMap[v.ID] = v
	}

	for _, v := range categories {
		categoriesMap[v.ID] = v
	}

	for _, v := range results {
		v.Agency = agenciesMap[v.AgencyId]
		v.Category = categoriesMap[v.CategoryId]
	}

	return results, nil
}

type agencyFieldTreeNode struct {
	AgencyId uint                  `json:"agencyId" gorm:"agency_id"`
	FieldId  uint                  `json:"fieldId" gorm:"fieldId"`
	Agency   *dms.DocumentAgencies `json:"agency" gorm:"-"`
	Field    *dms.DocumentFields   `json:"field" gorm:"-"`
	Count    uint                  `json:"count" gorm:"count"`
}

func (documentAgenciesService *DocumentAgenciesService) GetAgencyTreeForField(pageInfo request.GetById) (list interface{}, err error) {
	var results []*agencyFieldTreeNode

	err = global.GVA_DB.Model(&dms.Documents{}).
		Select("documents.agency_id as agency_id, document_field_references.field_id as field_id, count(documents.id) as count").
		Where("documents.agency_id = ? AND type = ? AND status = ?", pageInfo.ID, dms.TYPE_DOCUMENT, dms.STATUS_PUBLISHED).
		Joins("JOIN document_field_references ON documents.id = document_field_references.document_id").
		Group("documents.agency_id, document_field_references.field_id").
		Find(&results).Error

	if results == nil {
		return make([]*agencyFieldTreeNode, 0), nil
	}

	// attach data
	var agencies []*dms.DocumentAgencies
	var fields []*dms.DocumentFields

	agenciesMap := map[uint]*dms.DocumentAgencies{}
	fieldsMap := map[uint]*dms.DocumentFields{}

	err = global.GVA_DB.Model(&dms.DocumentAgencies{}).Find(&agencies).Error
	if err != nil {
		return nil, err
	}

	err = global.GVA_DB.Model(&dms.DocumentFields{}).Find(&fields).Error
	if err != nil {
		return nil, err
	}

	for _, v := range agencies {
		agenciesMap[v.ID] = v
	}

	for _, v := range fields {
		fieldsMap[v.ID] = v
	}

	for _, v := range results {
		v.Agency = agenciesMap[v.AgencyId]
		v.Field = fieldsMap[v.FieldId]
	}

	return results, nil
}

type agencyStats struct {
	AgencyId uint `json:"agency_id" gorm:"agency_id"`
	Count    uint `json:"count"`
}

func (documentAgenciesService *DocumentAgenciesService) attachDocumentCount(agencies []*dms.DocumentAgencies) (err error) {
	results := make([]agencyStats, 0)

	err = global.GVA_DB.Model(&dms.DocumentAgencyReferences{}).
		Select("agency_id, count(id) as count").
		Group("agency_id").
		Find(&results).Error

	if err != nil {
		return err
	}

	resultsAsMap := map[uint]uint{}
	for _, stat := range results {
		resultsAsMap[stat.AgencyId] = stat.Count
	}

	for _, agency := range agencies {
		if val, ok := resultsAsMap[agency.ID]; ok {
			agency.Count = int(val)
		}
	}

	return nil
}
