package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
	dmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/dms/request"
	sysMod "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
	"strconv"
	"time"
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
			DateIssued:      null.Time{},
			StillInEffect:   false,
			EffectDate:      null.Time{},
			ExpirationDate:  null.Time{},
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
				Title:      draft.Title,
				ShortTitle: "--",
				SignText:   "--",
				UserId:     userId,
				DoneAt:     null.Time{},
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

// CreateFullDocument create new full document
func (documentsService *DocumentsService) CreateFullDocument(full dmsReq.FullDocument) (doc *dms.Documents, err error) {
	var document dms.Documents

	categoryService := new(DocumentCategoriesService)
	agencyService := new(DocumentAgenciesService)

	category, catErr := categoryService.GetDocumentCategories(full.CategoryId)
	if catErr != nil {
		return nil, catErr
	}

	agency, agencyErr := agencyService.GetDocumentAgencies(full.AgencyId)
	if agencyErr != nil {
		return nil, agencyErr
	}

	shortTitle := category.Name + " " + full.SignText

	maxAuthorityId, authErr := system.AuthorityServiceApp.GetMaxId()
	if authErr != nil {
		return nil, authErr
	}

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// create new document
		document = dms.Documents{
			GVA_MODEL:       global.GVA_MODEL{},
			Title:           full.Title,
			ShortTitle:      shortTitle,
			Expert:          full.Expert,
			Content:         full.Content,
			DateIssued:      full.DateIssued,
			StillInEffect:   true,
			EffectDate:      full.EffectDate,
			ExpirationDate:  full.ExpirationDate,
			SignNumber:      full.SignNumber,
			SignYear:        full.SignYear,
			SignCategory:    category.Code,
			SignAgency:      agency.Code,
			SignText:        full.SignText,
			CategoryId:      full.CategoryId,
			AgencyId:        full.AgencyId,
			CreatedBy:       full.CreatedBy,
			BeResponsibleBy: full.BeResponsibleBy,
			ViewCount:       0,
			DownloadCount:   0,
			Status:          full.Status,
			Type:            dms.TYPE_DOCUMENT,
			Priority:        full.Priority,
			ParentId:        0,
			CurrentId:       0,
			Path:            "",
		}

		err = tx.Model(&dms.Documents{}).Create(&document).Error
		if err != nil {
			return err
		}

		// create signers
		if len(full.ReqSigners) > 0 {
			signers := make([]dms.DocumentSignerReferences, 0)
			for _, userId := range full.ReqSigners {
				signers = append(signers, dms.DocumentSignerReferences{
					GVA_MODEL:  global.GVA_MODEL{},
					DocumentId: document.ID,
					UserId:     userId,
				})
			}

			err = tx.Model(&dms.DocumentSignerReferences{}).Create(&signers).Error
			if err != nil {
				return err
			}
		}

		// create fields
		if len(full.ReqFields) > 0 {
			fields := make([]dms.DocumentFieldReferences, 0)
			for _, fieldId := range full.ReqFields {
				fields = append(fields, dms.DocumentFieldReferences{
					GVA_MODEL:  global.GVA_MODEL{},
					FieldId:    fieldId,
					DocumentId: document.ID,
				})
			}

			err = tx.Model(&dms.DocumentFieldReferences{}).Create(&fields).Error
			if err != nil {
				return err
			}
		}

		// create file
		if full.ReqFileInfo != nil {
			file := dms.DocumentFiles{
				GVA_MODEL:  global.GVA_MODEL{},
				Name:       full.ReqFileInfo.Name,
				Key:        full.ReqFileInfo.Key,
				Url:        full.ReqFileInfo.Url,
				Tag:        full.ReqFileInfo.Tag,
				Size:       full.ReqFileInfo.Size,
				Order:      full.ReqFileInfo.Order,
				DocumentId: document.ID,
				Path:       full.ReqFileInfo.Path,
			}

			err = tx.Model(&dms.DocumentFiles{}).Create(&file).Error
			if err != nil {
				return err
			}
		}

		// create document - user relations
		if len(full.ReqUsersRelated) > 0 {
			users := make([]dms.DocumentUsers, 0)
			for _, userId := range full.ReqUsersRelated {
				doneAt := null.Time{}
				if document.Status == dms.STATUS_PUBLISHED {
					doneAt = null.TimeFrom(time.Now())
				}

				users = append(users, dms.DocumentUsers{
					GVA_MODEL:  global.GVA_MODEL{},
					DocumentId: document.ID,
					Title:      document.Title,
					ShortTitle: document.ShortTitle,
					SignText:   document.SignText,
					UserId:     userId,
					DoneAt:     doneAt,
				})
			}

			err = tx.Model(&dms.DocumentUsers{}).Create(&users).Error
			if err != nil {
				return err
			}
		}

		// create all other document related stuffs
		relations := make([]dms.DocumentRelationReferences, 0)

		if len(full.ReqUsersRelated) > 0 {
			for _, userId := range full.ReqUsersRelated {
				relations = append(relations, dms.DocumentRelationReferences{
					GVA_MODEL:    global.GVA_MODEL{},
					DocumentId:   document.ID,
					DestId:       userId,
					RelationType: dms.RELATION_TYPE_USER,
				})
			}
		}

		if len(full.ReqAgenciesRelated) > 0 {
			for _, agencyId := range full.ReqAgenciesRelated {
				relations = append(relations, dms.DocumentRelationReferences{
					GVA_MODEL:    global.GVA_MODEL{},
					DocumentId:   document.ID,
					DestId:       agencyId,
					RelationType: dms.RELATION_TYPE_AGENCY,
				})
			}
		}

		if len(full.ReqDocumentBaseOns) > 0 {
			for _, baseOnId := range full.ReqDocumentBaseOns {
				relations = append(relations, dms.DocumentRelationReferences{
					GVA_MODEL:    global.GVA_MODEL{},
					DocumentId:   document.ID,
					DestId:       baseOnId,
					RelationType: dms.RELATION_TYPE_DOC_BASE,
				})
			}
		}

		if len(full.ReqDocumentReferences) > 0 {
			for _, relationId := range full.ReqDocumentReferences {
				relations = append(relations, dms.DocumentRelationReferences{
					GVA_MODEL:    global.GVA_MODEL{},
					DocumentId:   document.ID,
					DestId:       relationId,
					RelationType: dms.RELATION_TYPE_DOC_RELATION,
				})
			}
		}

		if len(relations) > 0 {
			err = tx.Model(&dms.DocumentRelationReferences{}).Create(&relations).Error
			if err != nil {
				return err
			}
		}

		// Authorzing
		// view - edit - delete - download
		authorities := make([]sysMod.SysAuthority, 0)
		documentIdAsString := strconv.Itoa(int(document.ID))

		parentId := maxAuthorityId + 1
		viewId := maxAuthorityId + 2
		editId := maxAuthorityId + 3
		deleteId := maxAuthorityId + 4
		downloadId := maxAuthorityId + 5

		authorities = append(authorities, sysMod.SysAuthority{
			CreatedAt:     time.Time{},
			UpdatedAt:     time.Time{},
			DeletedAt:     nil,
			AuthorityId:   parentId,
			AuthorityName: "document-" + documentIdAsString,
			ParentId:      0,
			DefaultRouter: "",
		})

		authorities = append(authorities, sysMod.SysAuthority{
			CreatedAt:     time.Time{},
			UpdatedAt:     time.Time{},
			DeletedAt:     nil,
			AuthorityId:   viewId,
			AuthorityName: "view-" + documentIdAsString,
			ParentId:      parentId,
			DefaultRouter: "",
		})

		authorities = append(authorities, sysMod.SysAuthority{
			CreatedAt:     time.Time{},
			UpdatedAt:     time.Time{},
			DeletedAt:     nil,
			AuthorityId:   editId,
			AuthorityName: "edit-" + documentIdAsString,
			ParentId:      parentId,
			DefaultRouter: "",
		})

		authorities = append(authorities, sysMod.SysAuthority{
			CreatedAt:     time.Time{},
			UpdatedAt:     time.Time{},
			DeletedAt:     nil,
			AuthorityId:   deleteId,
			AuthorityName: "delete-" + documentIdAsString,
			ParentId:      parentId,
			DefaultRouter: "",
		})

		authorities = append(authorities, sysMod.SysAuthority{
			CreatedAt:     time.Time{},
			UpdatedAt:     time.Time{},
			DeletedAt:     nil,
			AuthorityId:   downloadId,
			AuthorityName: "download-" + documentIdAsString,
			ParentId:      parentId,
			DefaultRouter: "",
		})

		err = tx.Model(&sysMod.SysAuthority{}).Create(&authorities).Error
		if err != nil {
			return err
		}

		// Update casbin rule
		viewUrl = "/api/v1/documents"

		return nil
	})

	return &document, nil
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
