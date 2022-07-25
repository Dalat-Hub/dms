package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
	dmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/dms/request"
)

type DocumentPriorityLevelsService struct {
}

// CreateDocumentPriorityLevels create new priority level
func (documentPriorityLevelsService *DocumentPriorityLevelsService) CreateDocumentPriorityLevels(documentPriorityLevels dms.DocumentPriorityLevels) (err error) {
	err = global.GVA_DB.Create(&documentPriorityLevels).Error
	return err
}

// DeleteDocumentPriorityLevels delete priority level by ID
func (documentPriorityLevelsService *DocumentPriorityLevelsService) DeleteDocumentPriorityLevels(documentPriorityLevels dms.DocumentPriorityLevels) (err error) {
	err = global.GVA_DB.Delete(&documentPriorityLevels).Error
	return err
}

// DeleteDocumentPriorityLevelsByIds bulk delete priority level by IDs
func (documentPriorityLevelsService *DocumentPriorityLevelsService) DeleteDocumentPriorityLevelsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dms.DocumentPriorityLevels{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateDocumentPriorityLevels update priority level
func (documentPriorityLevelsService *DocumentPriorityLevelsService) UpdateDocumentPriorityLevels(documentPriorityLevels dms.DocumentPriorityLevels) (err error) {
	err = global.GVA_DB.Save(&documentPriorityLevels).Error
	return err
}

// GetDocumentPriorityLevels get priority level by ID
func (documentPriorityLevelsService *DocumentPriorityLevelsService) GetDocumentPriorityLevels(id uint) (documentPriorityLevels dms.DocumentPriorityLevels, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&documentPriorityLevels).Error
	return
}

// GetDocumentPriorityLevelsInfoList get list of priority levels
func (documentPriorityLevelsService *DocumentPriorityLevelsService) GetDocumentPriorityLevelsInfoList(info dmsReq.DocumentPriorityLevelsSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&dms.DocumentPriorityLevels{})

	var documentPriorityLevelss []dms.DocumentPriorityLevels

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&documentPriorityLevelss).Error
	return documentPriorityLevelss, total, err
}
