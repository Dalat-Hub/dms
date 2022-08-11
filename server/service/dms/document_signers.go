package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
	dmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/dms/request"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
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
	var documentSignerss []*dms.DocumentSigners
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&documentSignerss).Error
	if err != nil {
		return
	}

	service := new(DocumentSignerReferencesService)

	err = service.AttachDocumentCount(documentSignerss)
	if err != nil {
		return
	}

	err = documentSignersService.attachSignerTitle(documentSignerss)
	if err != nil {
		return
	}

	return documentSignerss, total, err
}

func (documentSignersService *DocumentSignersService) attachSignerTitle(signers []*dms.DocumentSigners) (err error) {
	var sysDictionary sysModel.SysDictionary

	err = global.GVA_DB.Model(&sysModel.SysDictionary{}).
		Where("type = ?", "signerTitles").
		Preload("SysDictionaryDetails").
		First(&sysDictionary).Error

	if err != nil {
		return err
	}

	detailMap := make(map[int]sysModel.SysDictionaryDetail)

	for _, v := range sysDictionary.SysDictionaryDetails {
		detailMap[v.Value] = v
	}

	for _, s := range signers {
		if val, ok := detailMap[s.Title]; ok {
			s.SignerTitle = val
		}
	}

	return nil
}
