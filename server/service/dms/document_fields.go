package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
	dmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/dms/request"
)

type DocumentFieldsService struct {
}

// CreateDocumentFields create new field
func (documentFieldsService *DocumentFieldsService) CreateDocumentFields(documentFields *dms.DocumentFields) (err error) {
	err = global.GVA_DB.Create(&documentFields).Error
	return err
}

// DeleteDocumentFields delete field by ID
func (documentFieldsService *DocumentFieldsService) DeleteDocumentFields(documentFields dms.DocumentFields) (err error) {
	err = global.GVA_DB.Delete(&documentFields).Error
	return err
}

// DeleteDocumentFieldsByIds bulk delete field by IDs
func (documentFieldsService *DocumentFieldsService) DeleteDocumentFieldsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dms.DocumentFields{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateDocumentFields update field
func (documentFieldsService *DocumentFieldsService) UpdateDocumentFields(documentFields dms.DocumentFields) (err error) {
	err = global.GVA_DB.Save(&documentFields).Error
	return err
}

// GetDocumentFields get field by ID
func (documentFieldsService *DocumentFieldsService) GetDocumentFields(id uint) (documentFields dms.DocumentFields, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&documentFields).Error
	return
}

// GetDocumentFieldsInfoList get list of fields
func (documentFieldsService *DocumentFieldsService) GetDocumentFieldsInfoList(info dmsReq.DocumentFieldsSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&dms.DocumentFields{})

	var documentFieldss []*dms.DocumentFields

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&documentFieldss).Error
	if err != nil {
		return
	}

	service := new(DocumentFieldReferencesService)

	err = service.AttachDocumentCount(documentFieldss)
	if err != nil {
		return
	}

	return documentFieldss, total, err
}
