package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
	dmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/dms/request"
	request2 "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
	"time"
)

type DocumentPermissionRequestService struct {
}

func (documentPermissionRequestService *DocumentPermissionRequestService) CreateDocumentPermissionRequest(documentPermissionRequest *dms.DocumentPermissionRequest) (err error) {
	err = global.GVA_DB.Create(&documentPermissionRequest).Error
	return err
}

func (documentPermissionRequestService *DocumentPermissionRequestService) ApprovePermissionRequest(request *dms.DocumentPermissionRequest, userInfo *request2.CustomClaims) (err error) {
	var documentPermissionRequest dms.DocumentPermissionRequest

	err = global.GVA_DB.Model(&dms.DocumentPermissionRequest{}).Where("id = ?", request.ID).First(&documentPermissionRequest).Error
	if err != nil {
		return err
	}

	documentPermissionRequest.AcceptAt = null.TimeFrom(time.Now())
	documentPermissionRequest.AcceptUserId = userInfo.ID

	authority := dms.DocumentRules{
		GVA_MODEL:  global.GVA_MODEL{},
		DocumentId: request.DocumentId,
		Permission: request.RequestPermission,
		SubjectId:  request.RequestUserId,
		Type:       dms.RULE_TYPE_USER,
	}

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Model(&dms.DocumentPermissionRequest{}).Where("id = ?", request.ID).Save(&documentPermissionRequest).Error
		if err != nil {
			return err
		}

		err = tx.Model(&dms.DocumentRules{}).Create(&authority).Error
		if err != nil {
			return err
		}

		return nil
	})

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

	db = db.Limit(limit).Offset(offset)

	// preloading
	db = db.Preload("Document")
	db = db.Preload("RequestUser")
	db = db.Preload("AcceptUser")

	err = db.Find(&documentPermissionRequests).Error
	if err != nil {
		return
	}

	return documentPermissionRequests, total, err
}
