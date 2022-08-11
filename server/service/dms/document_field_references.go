package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
	dmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/dms/request"
)

type DocumentFieldReferencesService struct {
}

// CreateDocumentFieldReferences create new pivot
func (documentFieldReferencesService *DocumentFieldReferencesService) CreateDocumentFieldReferences(documentFieldReferences dms.DocumentFieldReferences) (err error) {
	err = global.GVA_DB.Create(&documentFieldReferences).Error
	return err
}

// DeleteDocumentFieldReferences delete pivot
func (documentFieldReferencesService *DocumentFieldReferencesService) DeleteDocumentFieldReferences(documentFieldReferences dms.DocumentFieldReferences) (err error) {
	err = global.GVA_DB.Delete(&documentFieldReferences).Error
	return err
}

// DeleteDocumentFieldReferencesByIds bulk delete pivot
func (documentFieldReferencesService *DocumentFieldReferencesService) DeleteDocumentFieldReferencesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dms.DocumentFieldReferences{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateDocumentFieldReferences update pivot
func (documentFieldReferencesService *DocumentFieldReferencesService) UpdateDocumentFieldReferences(documentFieldReferences dms.DocumentFieldReferences) (err error) {
	err = global.GVA_DB.Save(&documentFieldReferences).Error
	return err
}

// GetDocumentFieldReferences get pivot by ID
func (documentFieldReferencesService *DocumentFieldReferencesService) GetDocumentFieldReferences(id uint) (documentFieldReferences dms.DocumentFieldReferences, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&documentFieldReferences).Error
	return
}

// GetDocumentFieldReferencesInfoList get list of pivots
func (documentFieldReferencesService *DocumentFieldReferencesService) GetDocumentFieldReferencesInfoList(info dmsReq.DocumentFieldReferencesSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&dms.DocumentFieldReferences{})

	var documentFieldReferencess []dms.DocumentFieldReferences

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&documentFieldReferencess).Error
	return documentFieldReferencess, total, err
}

type fieldStats struct {
	FieldId uint `json:"field_id"`
	Count   uint `json:"count"`
}

func (documentFieldReferencesService *DocumentFieldReferencesService) AttachDocumentCount(fields []*dms.DocumentFields) (err error) {
	documents := make([]dms.Documents, 0)

	err = global.GVA_DB.Model(&dms.Documents{}).
		Where("type = ? AND status = ?", dms.TYPE_DOCUMENT, dms.STATUS_PUBLISHED).
		Find(&documents).Error

	if err != nil {
		return err
	}

	documentIds := make([]uint, 0)
	for _, v := range documents {
		documentIds = append(documentIds, v.ID)
	}

	results := make([]fieldStats, 0)

	err = global.GVA_DB.Model(&dms.DocumentFieldReferences{}).
		Select("field_id, count(id) as count").
		Where("document_id IN ?", documentIds).
		Group("field_id").
		Find(&results).Error

	if err != nil {
		return err
	}

	resultsAsMap := map[uint]uint{}
	for _, stat := range results {
		resultsAsMap[stat.FieldId] = stat.Count
	}

	for _, field := range fields {
		if val, ok := resultsAsMap[field.ID]; ok {
			field.Count = int(val)
		}
	}

	return nil
}
