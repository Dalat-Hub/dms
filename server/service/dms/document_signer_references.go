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
