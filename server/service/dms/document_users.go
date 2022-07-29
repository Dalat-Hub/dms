package dms

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
	dmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/dms/request"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
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
	err = db.Preload("User").Limit(limit).Offset(offset).Find(&documentUserss).Error
	return documentUserss, total, err
}

// DistributeTasks distribute tasks
func (documentUsersService *DocumentUsersService) DistributeTasks(distribution dmsReq.DocumentUsersDistribution) (err error) {
	if len(distribution.UserIds) == 0 {
		return errors.New("vui lòng gửi danh sách phân bổ người dùng")
	}

	var task dms.DocumentUsers

	task, err = documentUsersService.GetDocumentUsers(distribution.ID)
	if err != nil {
		return err
	}

	newTasks := make([]dms.DocumentUsers, 0)

	for _, userId := range distribution.UserIds {
		newTasks = append(newTasks, dms.DocumentUsers{
			GVA_MODEL:  global.GVA_MODEL{},
			DocumentId: task.DocumentId,
			Title:      task.Title,
			ShortTitle: task.ShortTitle,
			SignText:   task.SignText,
			UserId:     userId,
			DoneAt:     null.Time{},
		})
	}

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// clear old tasks
		err = tx.Model(&dms.DocumentUsers{}).
			Unscoped().
			Where("document_id = ?", task.DocumentId).
			Delete(&dms.DocumentUsers{}).
			Error

		if err != nil {
			return err
		}

		// create new tasks
		err = tx.Model(&dms.DocumentUsers{}).Create(&newTasks).Error
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
