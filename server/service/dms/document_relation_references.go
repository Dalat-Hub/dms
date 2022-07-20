package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
	dmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/dms/request"
)

type DocumentRelationReferencesService struct {
}

// CreateDocumentRelationReferences create pivot
func (documentRelationReferencesService *DocumentRelationReferencesService) CreateDocumentRelationReferences(documentRelationReferences dms.DocumentRelationReferences) (err error) {
	err = global.GVA_DB.Create(&documentRelationReferences).Error
	return err
}

// DeleteDocumentRelationReferences delete pivot by ID
func (documentRelationReferencesService *DocumentRelationReferencesService) DeleteDocumentRelationReferences(documentRelationReferences dms.DocumentRelationReferences) (err error) {
	err = global.GVA_DB.Delete(&documentRelationReferences).Error
	return err
}

// DeleteDocumentRelationReferencesByIds bulk delete pivot by IDs
func (documentRelationReferencesService *DocumentRelationReferencesService) DeleteDocumentRelationReferencesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dms.DocumentRelationReferences{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateDocumentRelationReferences update pivot
func (documentRelationReferencesService *DocumentRelationReferencesService) UpdateDocumentRelationReferences(documentRelationReferences dms.DocumentRelationReferences) (err error) {
	err = global.GVA_DB.Save(&documentRelationReferences).Error
	return err
}

// GetDocumentRelationReferences get pivot by ID
func (documentRelationReferencesService *DocumentRelationReferencesService) GetDocumentRelationReferences(id uint) (documentRelationReferences dms.DocumentRelationReferences, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&documentRelationReferences).Error
	return
}

// GetDocumentRelationReferencesInfoList get list of pivots
func (documentRelationReferencesService *DocumentRelationReferencesService) GetDocumentRelationReferencesInfoList(info dmsReq.DocumentRelationReferencesSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&dms.DocumentRelationReferences{})

	var documentRelationReferencess []dms.DocumentRelationReferences

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&documentRelationReferencess).Error
	return documentRelationReferencess, total, err
}
