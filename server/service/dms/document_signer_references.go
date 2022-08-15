package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
	dmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/dms/request"
)

type DocumentSignerReferencesService struct {
}

// CreateDocumentSignerReferences create new pivot
func (documentSignerReferencesService *DocumentSignerReferencesService) CreateDocumentSignerReferences(documentSignerReferences dms.DocumentSignerReferences) (err error) {
	err = global.GVA_DB.Create(&documentSignerReferences).Error
	return err
}

// DeleteDocumentSignerReferences delete pivot by ID
func (documentSignerReferencesService *DocumentSignerReferencesService) DeleteDocumentSignerReferences(documentSignerReferences dms.DocumentSignerReferences) (err error) {
	err = global.GVA_DB.Delete(&documentSignerReferences).Error
	return err
}

// DeleteDocumentSignerReferencesByIds bulk delete pivot by Ids
func (documentSignerReferencesService *DocumentSignerReferencesService) DeleteDocumentSignerReferencesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dms.DocumentSignerReferences{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateDocumentSignerReferences update pivot
func (documentSignerReferencesService *DocumentSignerReferencesService) UpdateDocumentSignerReferences(documentSignerReferences dms.DocumentSignerReferences) (err error) {
	err = global.GVA_DB.Save(&documentSignerReferences).Error
	return err
}

// GetDocumentSignerReferences get pivot by ID
func (documentSignerReferencesService *DocumentSignerReferencesService) GetDocumentSignerReferences(id uint) (documentSignerReferences dms.DocumentSignerReferences, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&documentSignerReferences).Error
	return
}

// GetDocumentSignerReferencesInfoList get list of pivot
func (documentSignerReferencesService *DocumentSignerReferencesService) GetDocumentSignerReferencesInfoList(info dmsReq.DocumentSignerReferencesSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&dms.DocumentSignerReferences{})

	var documentSignerReferencess []dms.DocumentSignerReferences

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&documentSignerReferencess).Error
	return documentSignerReferencess, total, err
}

type signerStats struct {
	SignerId uint `json:"signer_id"`
	Count    uint `json:"count"`
}

func (documentSignerReferencesService *DocumentSignerReferencesService) AttachDocumentCount(signers []*dms.DocumentSigners) (err error) {
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

	results := make([]signerStats, 0)

	err = global.GVA_DB.Model(&dms.DocumentSignerReferences{}).
		Select("signer_id, count(id) as count").
		Where("document_id IN ?", documentIds).
		Group("signer_id").
		Find(&results).Error

	if err != nil {
		return err
	}

	resultsAsMap := map[uint]uint{}
	for _, stat := range results {
		resultsAsMap[stat.SignerId] = stat.Count
	}

	for _, signer := range signers {
		if val, ok := resultsAsMap[signer.ID]; ok {
			signer.Count = int(val)
		}
	}

	return nil
}
