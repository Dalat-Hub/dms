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

type agencyStats struct {
	AgencyId uint `json:"agency_id"`
	Count    uint `json:"count"`
}

func (documentAgenciesService *DocumentAgenciesService) attachDocumentCount(agencies []*dms.DocumentAgencies) (err error) {
	results := make([]agencyStats, 0)

	err = global.GVA_DB.Model(&dms.Documents{}).
		Select("agency_id, count(id) as count").
		Where("type = ? AND status = ?", dms.TYPE_DOCUMENT, dms.STATUS_PUBLISHED).
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
