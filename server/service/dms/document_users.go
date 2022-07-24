package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
	dmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/dms/request"
)

type DocumentUsersService struct {
}

// CreateDocumentUsers create new relation
func (documentUsersService *DocumentUsersService) CreateDocumentUsers(documentUsers dms.DocumentUsers) (err error) {
	err = global.GVA_DB.Create(&documentUsers).Error
	return err
}

// DeleteDocumentUsers delete relation by ID
func (documentUsersService *DocumentUsersService) DeleteDocumentUsers(documentUsers dms.DocumentUsers) (err error) {
	err = global.GVA_DB.Delete(&documentUsers).Error
	return err
}

// DeleteDocumentUsersByIds bulk delete relation by IDs
func (documentUsersService *DocumentUsersService) DeleteDocumentUsersByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dms.DocumentUsers{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateDocumentUsers update relation
func (documentUsersService *DocumentUsersService) UpdateDocumentUsers(documentUsers dms.DocumentUsers) (err error) {
	err = global.GVA_DB.Save(&documentUsers).Error
	return err
}

// GetDocumentUsers get relation by ID
func (documentUsersService *DocumentUsersService) GetDocumentUsers(id uint) (documentUsers dms.DocumentUsers, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&documentUsers).Error
	return
}

// GetDocumentUsersInfoList get list of relations
func (documentUsersService *DocumentUsersService) GetDocumentUsersInfoList(info dmsReq.DocumentUsersSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&dms.DocumentUsers{})
	var documentUserss []dms.DocumentUsers
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&documentUserss).Error
	return documentUserss, total, err
}
