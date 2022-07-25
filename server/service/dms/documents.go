package dms

import (
	"database/sql"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
	dmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/dms/request"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type DocumentsService struct {
}

// CreateDocuments create new document
func (documentsService *DocumentsService) CreateDocuments(documents dms.Documents) (err error) {
	err = global.GVA_DB.Create(&documents).Error
	return err
}

// CreateDraftDocument create new draft document
func (documentsService *DocumentsService) CreateDraftDocument(draft dmsReq.DraftDocument) (doc *dms.Documents, err error) {
	var document dms.Documents

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		document = dms.Documents{
			GVA_MODEL:       global.GVA_MODEL{},
			Title:           draft.Title,
			ShortTitle:      "",
			Expert:          "",
			Content:         "",
			DateIssued:      sql.NullTime{},
			StillInEffect:   false,
			EffectDate:      sql.NullTime{},
			ExpirationDate:  sql.NullTime{},
			SignNumber:      0,
			SignYear:        0,
			SignCategory:    "",
			SignAgency:      "",
			SignText:        uuid.NewV4().String(),
			CategoryId:      0,
			AgencyId:        0,
			CreatedBy:       0,
			BeResponsibleBy: 0,
			ViewCount:       0,
			DownloadCount:   0,
			Status:          dms.STATUS_DRAFT,
			Type:            dms.TYPE_DOCUMENT,
			Priority:        dms.PRIORITY_NORMAL,
			ParentId:        0,
			CurrentId:       0,
			Path:            "",
		}

		err = tx.Model(&dms.Documents{}).Create(&document).Error
		if err != nil {
			return err
		}

		relatedUsers := make([]dms.DocumentUsers, 0)
		for _, userId := range draft.RelatedUsers {
			relatedUsers = append(relatedUsers, dms.DocumentUsers{
				GVA_MODEL:  global.GVA_MODEL{},
				DocumentId: document.ID,
				UserId:     userId,
				DoneAt:     sql.NullTime{},
			})
		}

		if len(relatedUsers) > 0 {
			err = tx.Model(dms.DocumentUsers{}).Create(&relatedUsers).Error
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &document, err
}

// DeleteDocuments delete document by ID
func (documentsService *DocumentsService) DeleteDocuments(documents dms.Documents) (err error) {
	err = global.GVA_DB.Delete(&documents).Error
	return err
}

// DeleteDocumentsByIds bulk delete document by Ids
func (documentsService *DocumentsService) DeleteDocumentsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dms.Documents{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateDocuments update document
func (documentsService *DocumentsService) UpdateDocuments(documents dms.Documents) (err error) {
	err = global.GVA_DB.Save(&documents).Error
	return err
}

// GetDocuments get document by Id
func (documentsService *DocumentsService) GetDocuments(id uint) (documents dms.Documents, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&documents).Error
	return
}

// GetDocumentsInfoList get list of documents
func (documentsService *DocumentsService) GetDocumentsInfoList(info dmsReq.DocumentsSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&dms.Documents{})

	var documentss []dms.Documents

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&documentss).Error
	return documentss, total, err
}
