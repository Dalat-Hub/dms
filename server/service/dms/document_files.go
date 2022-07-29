package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
	dmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/dms/request"
)

type DocumentFilesService struct {
}

// CreateDocumentFiles create new file
func (documentFilesService *DocumentFilesService) CreateDocumentFiles(documentFiles dms.DocumentFiles) (err error) {
	err = global.GVA_DB.Create(&documentFiles).Error
	return err
}

// DeleteDocumentFiles delete file by ID
func (documentFilesService *DocumentFilesService) DeleteDocumentFiles(documentFiles dms.DocumentFiles) (err error) {
	err = global.GVA_DB.Delete(&documentFiles).Error
	return err
}

// DeleteDocumentFilesByIds bulk delete file by IDs
func (documentFilesService *DocumentFilesService) DeleteDocumentFilesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dms.DocumentFiles{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateDocumentFiles update file
func (documentFilesService *DocumentFilesService) UpdateDocumentFiles(documentFiles dms.DocumentFiles) (err error) {
	err = global.GVA_DB.Save(&documentFiles).Error
	return err
}

// GetDocumentFiles get file by ID
func (documentFilesService *DocumentFilesService) GetDocumentFiles(id uint) (documentFiles dms.DocumentFiles, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&documentFiles).Error
	return
}

// GetDocumentFilesInfoList get list of files
func (documentFilesService *DocumentFilesService) GetDocumentFilesInfoList(info dmsReq.DocumentFilesSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&dms.DocumentFiles{})

	var documentFiless []dms.DocumentFiles

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&documentFiless).Error
	return documentFiless, total, err
}
