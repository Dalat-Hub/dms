package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
	dmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/dms/request"
)

type DocumentAgencyReferencesService struct {
}

// CreateDocumentAgencyReferences 创建DocumentAgencyReferences记录
// Author [piexlmax](https://github.com/piexlmax)
func (documentAgencyReferencesService *DocumentAgencyReferencesService) CreateDocumentAgencyReferences(documentAgencyReferences dms.DocumentAgencyReferences) (err error) {
	err = global.GVA_DB.Create(&documentAgencyReferences).Error
	return err
}

// DeleteDocumentAgencyReferences 删除DocumentAgencyReferences记录
// Author [piexlmax](https://github.com/piexlmax)
func (documentAgencyReferencesService *DocumentAgencyReferencesService) DeleteDocumentAgencyReferences(documentAgencyReferences dms.DocumentAgencyReferences) (err error) {
	err = global.GVA_DB.Delete(&documentAgencyReferences).Error
	return err
}

// DeleteDocumentAgencyReferencesByIds 批量删除DocumentAgencyReferences记录
// Author [piexlmax](https://github.com/piexlmax)
func (documentAgencyReferencesService *DocumentAgencyReferencesService) DeleteDocumentAgencyReferencesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dms.DocumentAgencyReferences{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateDocumentAgencyReferences 更新DocumentAgencyReferences记录
// Author [piexlmax](https://github.com/piexlmax)
func (documentAgencyReferencesService *DocumentAgencyReferencesService) UpdateDocumentAgencyReferences(documentAgencyReferences dms.DocumentAgencyReferences) (err error) {
	err = global.GVA_DB.Save(&documentAgencyReferences).Error
	return err
}

// GetDocumentAgencyReferences 根据id获取DocumentAgencyReferences记录
// Author [piexlmax](https://github.com/piexlmax)
func (documentAgencyReferencesService *DocumentAgencyReferencesService) GetDocumentAgencyReferences(id uint) (documentAgencyReferences dms.DocumentAgencyReferences, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&documentAgencyReferences).Error
	return
}

// GetDocumentAgencyReferencesInfoList 分页获取DocumentAgencyReferences记录
// Author [piexlmax](https://github.com/piexlmax)
func (documentAgencyReferencesService *DocumentAgencyReferencesService) GetDocumentAgencyReferencesInfoList(info dmsReq.DocumentAgencyReferencesSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&dms.DocumentAgencyReferences{})
	var documentAgencyReferencess []dms.DocumentAgencyReferences
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&documentAgencyReferencess).Error
	return documentAgencyReferencess, total, err
}
