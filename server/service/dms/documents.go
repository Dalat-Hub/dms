package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
	dmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/dms/request"
)

type DocumentsService struct {
}

// CreateDocuments create new document
func (documentsService *DocumentsService) CreateDocuments(documents dms.Documents) (err error) {
	err = global.GVA_DB.Create(&documents).Error
	return err
}

// DeleteDocuments delete document by ID
func (documentsService *DocumentsService) DeleteDocuments(documents dms.Documents) (err error) {
	err = global.GVA_DB.Delete(&documents).Error
	return err
}

// DeleteDocumentsByIds bulk delete document by Ids
func (documentsService *DocumentsService) DeleteDocumentsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dms.Documents{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateDocuments update document
func (documentsService *DocumentsService) UpdateDocuments(documents dms.Documents) (err error) {
	err = global.GVA_DB.Save(&documents).Error
	return err
}

// GetDocuments get document by Id
func (documentsService *DocumentsService) GetDocuments(id uint) (documents dms.Documents, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&documents).Error
	return
}

// GetDocumentsInfoList get list of documents
func (documentsService *DocumentsService) GetDocumentsInfoList(info dmsReq.DocumentsSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&dms.Documents{})

	var documentss []dms.Documents

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&documentss).Error
	return documentss, total, err
}
