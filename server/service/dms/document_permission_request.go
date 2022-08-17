package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
	dmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/dms/request"
)

type DocumentPermissionRequestService struct {
}

func (documentPermissionRequestService *DocumentPermissionRequestService) CreateDocumentPermissionRequest(documentPermissionRequest dms.DocumentPermissionRequest) (err error) {
	err = global.GVA_DB.Create(&documentPermissionRequest).Error
	return err
}

func (documentPermissionRequestService *DocumentPermissionRequestService) DeleteDocumentPermissionRequest(documentPermissionRequest dms.DocumentPermissionRequest) (err error) {
	err = global.GVA_DB.Delete(&documentPermissionRequest).Error
	return err
}

func (documentPermissionRequestService *DocumentPermissionRequestService) DeleteDocumentPermissionRequestByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dms.DocumentPermissionRequest{}, "id in ?", ids.Ids).Error
	return err
}

func (documentPermissionRequestService *DocumentPermissionRequestService) UpdateDocumentPermissionRequest(documentPermissionRequest dms.DocumentPermissionRequest) (err error) {
	err = global.GVA_DB.Save(&documentPermissionRequest).Error
	return err
}

func (documentPermissionRequestService *DocumentPermissionRequestService) GetDocumentPermissionRequest(id uint) (documentPermissionRequest dms.DocumentPermissionRequest, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&documentPermissionRequest).Error
	return
}

func (documentPermissionRequestService *DocumentPermissionRequestService) GetDocumentPermissionRequestInfoList(info dmsReq.DocumentPermissionRequestSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&dms.DocumentPermissionRequest{})
	var documentPermissionRequests []dms.DocumentPermissionRequest

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&documentPermissionRequests).Error
	return documentPermissionRequests, total, err
}
