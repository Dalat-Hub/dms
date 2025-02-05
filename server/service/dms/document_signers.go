package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
	dmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/dms/request"
)

type DocumentSignersService struct {
}

// CreateDocumentSigners 创建DocumentSigners记录
// Author [piexlmax](https://github.com/piexlmax)
func (documentSignersService *DocumentSignersService) CreateDocumentSigners(documentSigners *dms.DocumentSigners) (err error) {
	err = global.GVA_DB.Create(&documentSigners).Error
	return err
}

// DeleteDocumentSigners 删除DocumentSigners记录
// Author [piexlmax](https://github.com/piexlmax)
func (documentSignersService *DocumentSignersService) DeleteDocumentSigners(documentSigners dms.DocumentSigners) (err error) {
	err = global.GVA_DB.Delete(&documentSigners).Error
	return err
}

// DeleteDocumentSignersByIds 批量删除DocumentSigners记录
// Author [piexlmax](https://github.com/piexlmax)
func (documentSignersService *DocumentSignersService) DeleteDocumentSignersByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dms.DocumentSigners{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateDocumentSigners 更新DocumentSigners记录
// Author [piexlmax](https://github.com/piexlmax)
func (documentSignersService *DocumentSignersService) UpdateDocumentSigners(documentSigners dms.DocumentSigners) (err error) {
	err = global.GVA_DB.Save(&documentSigners).Error
	return err
}

// GetDocumentSigners 根据id获取DocumentSigners记录
// Author [piexlmax](https://github.com/piexlmax)
func (documentSignersService *DocumentSignersService) GetDocumentSigners(id uint) (documentSigners dms.DocumentSigners, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&documentSigners).Error
	return
}

// GetDocumentSignersInfoList 分页获取DocumentSigners记录
// Author [piexlmax](https://github.com/piexlmax)
func (documentSignersService *DocumentSignersService) GetDocumentSignersInfoList(info dmsReq.DocumentSignersSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&dms.DocumentSigners{})
	var documentSignerss []dms.DocumentSigners
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&documentSignerss).Error
	return documentSignerss, total, err
}
