package dms

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
	dmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/dms/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms/request/documents"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms/response"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
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
func (documentsService *DocumentsService) CreateDraftDocument(draft dmsReq.DraftDocument, loginUserId uint) (doc *dms.Documents, err error) {
	var document dms.Documents

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		document = dms.Documents{
			GVA_MODEL:        global.GVA_MODEL{},
			Title:            draft.Title,
			ShortTitle:       draft.Title,
			Expert:           draft.Title,
			Content:          draft.Title,
			DateIssued:       null.Time{},
			StillInEffect:    false,
			EffectDate:       null.Time{},
			ExpirationDate:   null.Time{},
			SignNumber:       0,
			SignYear:         0,
			SignCategory:     "",
			SignAgency:       "",
			SignText:         uuid.NewV4().String(),
			CategoryId:       0,
			AgencyId:         0,
			CreatedBy:        loginUserId,
			UpdatedBy:        0,
			BeResponsibleBy:  loginUserId,
			ViewCount:        0,
			DownloadCount:    0,
			Status:           dms.STATUS_DRAFT,
			Type:             dms.TYPE_DOCUMENT,
			Priority:         dms.PRIORITY_NORMAL,
			BelongTo:         0,
			ParentId:         0,
			CurrentId:        0,
			Path:             "",
			PublicToView:     false,
			PublicToDownload: false,
		}

		err = tx.Model(&dms.Documents{}).Create(&document).Error
		if err != nil {
			return err
		}

		relatedUsers := make([]dms.DocumentUsers, 0)
		hasUndefinedUser := false

		if len(draft.RelatedUsers) == 0 {
			hasUndefinedUser = true
		}

		for _, userId := range draft.RelatedUsers {
			if userId == 0 {
				hasUndefinedUser = true
				break
			}

			relatedUsers = append(relatedUsers, dms.DocumentUsers{
				GVA_MODEL:  global.GVA_MODEL{},
				DocumentId: document.ID,
				Title:      draft.Title,
				ShortTitle: draft.Title,
				SignText:   "",
				UserId:     userId,
				DoneAt:     null.Time{},
			})
		}

		if hasUndefinedUser {
			relatedUsers = make([]dms.DocumentUsers, 0)

			relatedUsers = append(relatedUsers, dms.DocumentUsers{
				GVA_MODEL:  global.GVA_MODEL{},
				DocumentId: document.ID,
				Title:      draft.Title,
				ShortTitle: draft.Title,
				SignText:   "",
				UserId:     0,
				DoneAt:     null.Time{},
			})
		}

		if len(relatedUsers) > 0 {
			err = tx.Model(dms.DocumentUsers{}).Create(&relatedUsers).Error
			if err != nil {
				return err
			}
		}

		// authorizing
		authorities := make([]dms.DocumentRules, 0)

		authorities = append(authorities, dms.DocumentRules{
			GVA_MODEL:  global.GVA_MODEL{},
			DocumentId: document.ID,
			Permission: dms.PERMISSION_OWNER,
			SubjectId:  loginUserId,
			Type:       dms.RULE_TYPE_USER,
		})

		for _, v := range relatedUsers {
			authorities = append(authorities, dms.DocumentRules{
				GVA_MODEL:  global.GVA_MODEL{},
				DocumentId: document.ID,
				Permission: dms.PERMISSION_VIEW,
				SubjectId:  v.ID,
				Type:       dms.RULE_TYPE_USER,
			})

			authorities = append(authorities, dms.DocumentRules{
				GVA_MODEL:  global.GVA_MODEL{},
				DocumentId: document.ID,
				Permission: dms.PERMISSION_EDIT,
				SubjectId:  v.ID,
				Type:       dms.RULE_TYPE_USER,
			})
		}

		err = tx.Model(&dms.DocumentRules{}).Create(&authorities).Error
		if err != nil {
			return err
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
					SignerId:   userId,
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

		// authorizing
		authorities := make([]dms.DocumentRules, 0)

		authorities = append(authorities, dms.DocumentRules{
			GVA_MODEL:  global.GVA_MODEL{},
			DocumentId: document.ID,
			Permission: dms.PERMISSION_OWNER,
			SubjectId:  document.CreatedBy,
			Type:       dms.RULE_TYPE_USER,
		})

		if full.ReqRuleInfo.View == dms.RULE_VIEW_ALL {
			authorities = append(authorities, dms.DocumentRules{
				GVA_MODEL:  global.GVA_MODEL{},
				DocumentId: document.ID,
				Permission: dms.PERMISSION_VIEW,
				SubjectId:  101,
				Type:       dms.RULE_TYPE_GROUP,
			})
		} else if full.ReqRuleInfo.View == dms.RULE_VIEW_LIMIT {
			if len(full.ReqRuleInfo.ViewUsers) > 0 {
				for _, userId := range full.ReqRuleInfo.ViewUsers {
					authorities = append(authorities, dms.DocumentRules{
						GVA_MODEL:  global.GVA_MODEL{},
						DocumentId: document.ID,
						Permission: dms.PERMISSION_VIEW,
						SubjectId:  userId,
						Type:       dms.RULE_TYPE_USER,
					})
				}
			}

			if len(full.ReqRuleInfo.ViewRoles) > 0 {
				for _, roleId := range full.ReqRuleInfo.ViewRoles {
					authorities = append(authorities, dms.DocumentRules{
						GVA_MODEL:  global.GVA_MODEL{},
						DocumentId: document.ID,
						Permission: dms.PERMISSION_VIEW,
						SubjectId:  roleId,
						Type:       dms.RULE_TYPE_GROUP,
					})
				}
			}
		}

		if full.ReqRuleInfo.Download == dms.RULE_DOWNLOAD_ALL {
			authorities = append(authorities, dms.DocumentRules{
				GVA_MODEL:  global.GVA_MODEL{},
				DocumentId: document.ID,
				Permission: dms.PERMISSION_DOWNLOAD,
				SubjectId:  102,
				Type:       dms.RULE_TYPE_GROUP,
			})
		} else if full.ReqRuleInfo.Download == dms.RULE_DOWNLOAD_LIMIT {
			if len(full.ReqRuleInfo.DownloadUsers) > 0 {
				for _, userId := range full.ReqRuleInfo.DownloadUsers {
					authorities = append(authorities, dms.DocumentRules{
						GVA_MODEL:  global.GVA_MODEL{},
						DocumentId: document.ID,
						Permission: dms.PERMISSION_DOWNLOAD,
						SubjectId:  userId,
						Type:       dms.RULE_TYPE_USER,
					})
				}
			}

			if len(full.ReqRuleInfo.DownloadRoles) > 0 {
				for _, roleId := range full.ReqRuleInfo.DownloadRoles {
					authorities = append(authorities, dms.DocumentRules{
						GVA_MODEL:  global.GVA_MODEL{},
						DocumentId: document.ID,
						Permission: dms.PERMISSION_DOWNLOAD,
						SubjectId:  roleId,
						Type:       dms.RULE_TYPE_GROUP,
					})
				}
			}
		}

		if full.ReqRuleInfo.Edit == dms.RULE_EDIT_LIMIT {
			if len(full.ReqRuleInfo.EditUsers) > 0 {
				for _, userId := range full.ReqRuleInfo.EditUsers {
					authorities = append(authorities, dms.DocumentRules{
						GVA_MODEL:  global.GVA_MODEL{},
						DocumentId: document.ID,
						Permission: dms.PERMISSION_EDIT,
						SubjectId:  userId,
						Type:       dms.RULE_TYPE_USER,
					})
				}
			}

			if len(full.ReqRuleInfo.EditRoles) > 0 {
				for _, roleId := range full.ReqRuleInfo.EditRoles {
					authorities = append(authorities, dms.DocumentRules{
						GVA_MODEL:  global.GVA_MODEL{},
						DocumentId: document.ID,
						Permission: dms.PERMISSION_EDIT,
						SubjectId:  roleId,
						Type:       dms.RULE_TYPE_GROUP,
					})
				}
			}
		}

		if full.ReqRuleInfo.Owner == dms.RULE_OWNER_LIMIT {
			if len(full.ReqRuleInfo.OwnerUsers) > 0 {
				for _, userId := range full.ReqRuleInfo.OwnerUsers {
					authorities = append(authorities, dms.DocumentRules{
						GVA_MODEL:  global.GVA_MODEL{},
						DocumentId: document.ID,
						Permission: dms.PERMISSION_OWNER,
						SubjectId:  userId,
						Type:       dms.RULE_TYPE_USER,
					})
				}
			}

			if len(full.ReqRuleInfo.OwnerRoles) > 0 {
				for _, roleId := range full.ReqRuleInfo.OwnerRoles {
					authorities = append(authorities, dms.DocumentRules{
						GVA_MODEL:  global.GVA_MODEL{},
						DocumentId: document.ID,
						Permission: dms.PERMISSION_OWNER,
						SubjectId:  roleId,
						Type:       dms.RULE_TYPE_GROUP,
					})
				}
			}
		}

		err = tx.Model(&dms.DocumentRules{}).Create(&authorities).Error
		if err != nil {
			return err
		}

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

// UpdateBasicDocumentInformation update basic document information
func (documentsService *DocumentsService) UpdateBasicDocumentInformation(basic documents.UpdateBasic) (err error) {
	var oldDocument dms.Documents
	err = global.GVA_DB.Model(&dms.Documents{}).First(&oldDocument, "id = ?", basic.ID).Error
	if err != nil {
		return err
	}

	var revisionDocument *dms.Documents
	revisionDocument, err = documentsService.Duplicate(oldDocument.ID, basic.UpdatedBy, dms.TYPE_REVISION)
	if err != nil {
		return err
	}

	categoryService := new(DocumentCategoriesService)
	agencyService := new(DocumentAgenciesService)

	category, catErr := categoryService.GetDocumentCategories(basic.CategoryId)
	if catErr != nil {
		return catErr
	}

	_, agencyErr := agencyService.GetDocumentAgencies(basic.AgencyId)
	if agencyErr != nil {
		return agencyErr
	}

	shortTitle := category.Name + " " + basic.SignText

	oldDocument.Title = basic.Title
	oldDocument.Expert = basic.Expert
	oldDocument.ShortTitle = shortTitle
	oldDocument.Content = basic.Content
	oldDocument.DateIssued = basic.DateIssued
	oldDocument.EffectDate = basic.EffectDate
	oldDocument.ExpirationDate = basic.ExpirationDate
	oldDocument.SignNumber = basic.SignNumber
	oldDocument.SignYear = basic.SignYear
	oldDocument.SignCategory = basic.SignCategory
	oldDocument.SignAgency = basic.SignAgency
	oldDocument.SignText = basic.SignText
	oldDocument.CategoryId = basic.CategoryId
	oldDocument.AgencyId = basic.AgencyId
	oldDocument.Priority = basic.Priority
	oldDocument.Status = basic.Status
	oldDocument.UpdatedBy = basic.UpdatedBy
	oldDocument.BeResponsibleBy = basic.BeResponsibleBy

	if oldDocument.Path == "" {
		oldDocument.Path = strconv.Itoa(int(revisionDocument.ID))
	} else {
		oldDocument.Path = strconv.Itoa(int(revisionDocument.ID)) + "/" + oldDocument.Path
	}

	oldDocument.ParentId = revisionDocument.ID

	if revisionDocument.BelongTo == 0 {
		oldDocument.BelongTo = revisionDocument.ID
	}

	newFieldReferences := make([]dms.DocumentFieldReferences, 0)
	newSignerReferences := make([]dms.DocumentSignerReferences, 0)

	for _, fieldId := range basic.Fields {
		newFieldReferences = append(newFieldReferences, dms.DocumentFieldReferences{
			GVA_MODEL:  global.GVA_MODEL{},
			FieldId:    fieldId,
			DocumentId: oldDocument.ID,
		})
	}

	for _, userId := range basic.Signers {
		newSignerReferences = append(newSignerReferences, dms.DocumentSignerReferences{
			GVA_MODEL:  global.GVA_MODEL{},
			DocumentId: oldDocument.ID,
			SignerId:   userId,
		})
	}

	// authorizing
	authorities := make([]dms.DocumentRules, 0)

	authorities = append(authorities, dms.DocumentRules{
		GVA_MODEL:  global.GVA_MODEL{},
		DocumentId: revisionDocument.ID,
		Permission: dms.PERMISSION_OWNER,
		SubjectId:  oldDocument.CreatedBy,
		Type:       dms.RULE_TYPE_USER,
	})

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// update document
		err = tx.Model(&dms.Documents{}).Where("id = ?", oldDocument.ID).Save(&oldDocument).Error
		if err != nil {
			return err
		}

		// clear old fields
		err = tx.Model(&dms.DocumentFieldReferences{}).
			Unscoped().
			Where("document_id = ?", oldDocument.ID).
			Delete(&dms.DocumentFieldReferences{}).
			Error
		if err != nil {
			return err
		}

		// clear old signers
		err = tx.Model(&dms.DocumentSignerReferences{}).
			Unscoped().
			Where("document_id = ?", oldDocument.ID).
			Delete(&dms.DocumentSignerReferences{}).
			Error
		if err != nil {
			return err
		}

		// create new fields

		if len(newFieldReferences) > 0 {
			err = tx.Model(&dms.DocumentFieldReferences{}).Create(&newFieldReferences).Error
			if err != nil {
				return err
			}
		}

		// create new signers
		if len(newSignerReferences) > 0 {
			err = tx.Model(&dms.DocumentSignerReferences{}).Create(&newSignerReferences).Error
			if err != nil {
				return err
			}
		}

		// add authority for the revision document
		err = tx.Model(&dms.DocumentRules{}).Create(&authorities).Error
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

// UpdateRelationDocumentInformation update relation document information
func (documentsService *DocumentsService) UpdateRelationDocumentInformation(relation documents.UpdateRelation) (err error) {
	var oldDocument dms.Documents
	err = global.GVA_DB.Model(&dms.Documents{}).First(&oldDocument, "id = ?", relation.ID).Error
	if err != nil {
		return err
	}

	var revisionDocument *dms.Documents
	revisionDocument, err = documentsService.Duplicate(oldDocument.ID, relation.UpdatedBy, dms.TYPE_REVISION)
	if err != nil {
		return err
	}

	oldDocument.UpdatedBy = relation.UpdatedBy
	oldDocument.BeResponsibleBy = relation.BeResponsibleBy

	if oldDocument.Path == "" {
		oldDocument.Path = strconv.Itoa(int(revisionDocument.ID))
	} else {
		oldDocument.Path = strconv.Itoa(int(revisionDocument.ID)) + "/" + oldDocument.Path
	}

	oldDocument.ParentId = revisionDocument.ID

	if revisionDocument.BelongTo == 0 {
		oldDocument.BelongTo = revisionDocument.ID
	}

	referenceStuffs := make([]dms.DocumentRelationReferences, 0)

	for _, documentId := range relation.BasedDocuments {
		referenceStuffs = append(referenceStuffs, dms.DocumentRelationReferences{
			GVA_MODEL:    global.GVA_MODEL{},
			DocumentId:   oldDocument.ID,
			DestId:       documentId,
			RelationType: dms.RELATION_TYPE_DOC_BASE,
		})
	}

	for _, documentId := range relation.RelatedDocuments {
		referenceStuffs = append(referenceStuffs, dms.DocumentRelationReferences{
			GVA_MODEL:    global.GVA_MODEL{},
			DocumentId:   oldDocument.ID,
			DestId:       documentId,
			RelationType: dms.RELATION_TYPE_DOC_RELATION,
		})
	}

	for _, userId := range relation.RelatedUsers {
		referenceStuffs = append(referenceStuffs, dms.DocumentRelationReferences{
			GVA_MODEL:    global.GVA_MODEL{},
			DocumentId:   oldDocument.ID,
			DestId:       userId,
			RelationType: dms.RELATION_TYPE_USER,
		})
	}

	for _, agencyId := range relation.RelatedAgencies {
		referenceStuffs = append(referenceStuffs, dms.DocumentRelationReferences{
			GVA_MODEL:    global.GVA_MODEL{},
			DocumentId:   oldDocument.ID,
			DestId:       agencyId,
			RelationType: dms.RELATION_TYPE_AGENCY,
		})
	}

	// authorizing
	authorities := make([]dms.DocumentRules, 0)

	authorities = append(authorities, dms.DocumentRules{
		GVA_MODEL:  global.GVA_MODEL{},
		DocumentId: revisionDocument.ID,
		Permission: dms.PERMISSION_OWNER,
		SubjectId:  oldDocument.CreatedBy,
		Type:       dms.RULE_TYPE_USER,
	})

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// update document
		err = tx.Model(&dms.Documents{}).Where("id = ?", oldDocument.ID).Save(&oldDocument).Error
		if err != nil {
			return err
		}

		// clear old relations
		err = tx.Model(&dms.DocumentRelationReferences{}).
			Unscoped().
			Where("document_id = ?", oldDocument.ID).
			Delete(&dms.DocumentRelationReferences{}).
			Error

		if err != nil {
			return err
		}

		// create new relations
		if len(referenceStuffs) > 0 {
			err = tx.Model(&dms.DocumentRelationReferences{}).Create(&referenceStuffs).Error
			if err != nil {
				return err
			}
		}

		// add authority for the revision document
		err = tx.Model(&dms.DocumentRules{}).Create(&authorities).Error
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

// UpdateDocumentAttachedFiles update document attached files
func (documentsService *DocumentsService) UpdateDocumentAttachedFiles(fileInfo documents.UpdateFile) (err error) {
	var oldDocument dms.Documents
	err = global.GVA_DB.Model(&dms.Documents{}).First(&oldDocument, "id = ?", fileInfo.ID).Error
	if err != nil {
		return err
	}

	var revisionDocument *dms.Documents
	revisionDocument, err = documentsService.Duplicate(oldDocument.ID, fileInfo.UpdatedBy, dms.TYPE_REVISION)
	if err != nil {
		return err
	}

	oldDocument.UpdatedBy = fileInfo.UpdatedBy
	oldDocument.BeResponsibleBy = fileInfo.BeResponsibleBy

	if oldDocument.Path == "" {
		oldDocument.Path = strconv.Itoa(int(revisionDocument.ID))
	} else {
		oldDocument.Path = strconv.Itoa(int(revisionDocument.ID)) + "/" + oldDocument.Path
	}

	oldDocument.ParentId = revisionDocument.ID

	if revisionDocument.BelongTo == 0 {
		oldDocument.BelongTo = revisionDocument.ID
	}

	newFile := dms.DocumentFiles{
		GVA_MODEL:  global.GVA_MODEL{},
		Name:       fileInfo.Name,
		Key:        fileInfo.Key,
		Url:        fileInfo.Url,
		Tag:        fileInfo.Tag,
		Size:       fileInfo.Size,
		Order:      fileInfo.Order,
		DocumentId: oldDocument.ID,
		Path:       fileInfo.Path,
	}

	// authorizing
	authorities := make([]dms.DocumentRules, 0)

	authorities = append(authorities, dms.DocumentRules{
		GVA_MODEL:  global.GVA_MODEL{},
		DocumentId: revisionDocument.ID,
		Permission: dms.PERMISSION_OWNER,
		SubjectId:  oldDocument.CreatedBy,
		Type:       dms.RULE_TYPE_USER,
	})

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// update document
		err = tx.Model(&dms.Documents{}).Where("id = ?", oldDocument.ID).Save(&oldDocument).Error
		if err != nil {
			return err
		}

		// clear old file
		err = tx.Model(&dms.DocumentFiles{}).
			Unscoped().
			Where("document_id = ?", oldDocument.ID).
			Delete(&dms.DocumentFiles{}).
			Error

		if err != nil {
			return err
		}

		// create new file
		err = tx.Model(&dms.DocumentFiles{}).Create(&newFile).Error
		if err != nil {
			return err
		}

		// add authority for the revision document
		err = tx.Model(&dms.DocumentRules{}).Create(&authorities).Error
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

// GetDocuments get document by Id
func (documentsService *DocumentsService) GetDocuments(doc dmsReq.DocumentsSearch, userId uint, userUUID uuid.UUID) (documents *dms.Documents, err error) {
	db := global.GVA_DB.Model(&dms.Documents{})
	var document dms.Documents

	if doc.PreloadAgency == 1 {
		db = db.Preload("Agency")
	}

	if doc.PreloadCategory == 1 {
		db = db.Preload("Category")
	}

	if doc.PreloadFields == 1 {
		db = db.Preload("Fields")
	}

	if doc.PreloadSigners == 1 {
		db = db.Preload("Signers")
	}

	err = db.First(&document, "id = ?", doc.ID).Error

	if err != nil {
		return nil, err
	}

	if doc.PreloadBasedDocs == 1 {
		if err, panicErr := documentsService.attachBaseDocuments(&document); panicErr {
			return nil, err
		}
	}

	if doc.PreloadRelatedDocs == 1 {
		if err, panicErr := documentsService.attachReferencesDocuments(&document); panicErr {
			return nil, err
		}
	}

	if doc.PreloadRelatedUsers == 1 {
		if err, panicErr := documentsService.attachRelatedUsers(&document); panicErr {
			return nil, err
		}
	}

	if doc.PreloadRelatedAgencies == 1 {
		if err, panicErr := documentsService.attachRelatedAgencies(&document); panicErr {
			return nil, err
		}
	}

	if doc.PreloadCreatedBy == 1 {
		if err, panicErr := documentsService.attachCreatedUser(&document); panicErr {
			return nil, err
		}
	}

	if doc.PreloadUpdatedBy == 1 {
		if err, panicErr := documentsService.attachUpdatedUser(&document); panicErr {
			return nil, err
		}
	}

	if doc.PreloadBeResponsibleBy == 1 {
		if err, panicErr := documentsService.attachResponsibleUser(&document); panicErr {
			return nil, err
		}
	}

	if doc.PreloadAuthority == 1 {
		service := new(DocumentRulesService)
		if err := service.CheckPermission(userId, userUUID, document.ID, dms.PERMISSION_OWNER); err == nil {
			if err, panicErr := documentsService.attachAuthority(&document); panicErr {
				return nil, err
			}
		}
	}

	return &document, nil
}

// GetDocumentsInfoList get list of documents
func (documentsService *DocumentsService) GetDocumentsInfoList(info dmsReq.DocumentsSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&dms.Documents{})

	var documentss []dms.Documents

	err = db.Where("type = ? AND status = ?", dms.TYPE_DOCUMENT, dms.STATUS_PUBLISHED).Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&documentss).Error
	return documentss, total, err
}

// GetDocumentRevisions get list of revisions of the given documents
func (documentsService *DocumentsService) GetDocumentRevisions(documentId uint) (list interface{}, err error) {
	revisions := make([]*dms.Documents, 0)

	err = global.GVA_DB.Model(&dms.Documents{}).
		Where("type = ? AND belong_to = ?", dms.TYPE_REVISION, documentId).
		Find(&revisions).
		Error

	if err != nil {
		return nil, err
	}

	for _, v := range revisions {
		if err, panicErr := documentsService.attachCreatedUser(v); panicErr {
			return nil, err
		}

		if err, panicErr := documentsService.attachUpdatedUser(v); panicErr {
			return nil, err
		}
	}

	return revisions, nil
}

// GetDocumentFiles get list of attached files
func (documentsService *DocumentsService) GetDocumentFiles(info request.GetById, download bool) (list interface{}, canDownload bool, err error) {
	var document dms.Documents
	err = global.GVA_DB.Model(&dms.Documents{}).First(&document, "id = ?", info.ID).Error
	if err != nil {
		return nil, false, nil
	}

	canDownload = false
	if download {
		canDownload = true
	} else {
		if document.PublicToDownload {
			canDownload = true
		}
	}

	files := make([]dms.DocumentFiles, 0)

	err = global.GVA_DB.Model(&dms.DocumentFiles{}).Where("document_id = ?", document.ID).Find(&files).Error
	if err != nil {
		return nil, false, err
	}

	return files, canDownload, nil
}

// Duplicate create new copy of the document
func (documentsService *DocumentsService) Duplicate(documentId uint, loginUserId uint, docType int) (revision *dms.Documents, err error) {
	// 1. get old document
	var oldDocument dms.Documents
	var newDocument dms.Documents

	err = global.GVA_DB.Model(&dms.Documents{}).First(&oldDocument, "id = ?", documentId).Error
	if err != nil {
		return nil, err
	}

	documentType := dms.TYPE_DOCUMENT
	if docType == dms.TYPE_REVISION {
		documentType = dms.TYPE_REVISION
	}

	var parentId uint = 0
	var newPath = ""
	var newBelongTo uint = 0
	var newViewCount uint = 0
	var newDownloadCount uint = 0
	var newUpdatedBy uint = 0
	var newCreatedBy = loginUserId
	var newResponsibleBy = loginUserId

	if docType != dms.TYPE_DOCUMENT {
		newViewCount = oldDocument.ViewCount
		newDownloadCount = oldDocument.DownloadCount
		newUpdatedBy = oldDocument.UpdatedBy
		newCreatedBy = oldDocument.CreatedBy
		newResponsibleBy = oldDocument.BeResponsibleBy

		newPath = oldDocument.Path
		parentId = oldDocument.ParentId
		newBelongTo = oldDocument.ID
	}

	newDocument = dms.Documents{
		GVA_MODEL:        global.GVA_MODEL{},
		Title:            oldDocument.Title,
		ShortTitle:       oldDocument.ShortTitle,
		Expert:           oldDocument.Expert,
		Content:          oldDocument.Content,
		DateIssued:       oldDocument.DateIssued,
		StillInEffect:    oldDocument.StillInEffect,
		EffectDate:       oldDocument.EffectDate,
		ExpirationDate:   oldDocument.ExpirationDate,
		SignNumber:       oldDocument.SignNumber,
		SignYear:         oldDocument.SignYear,
		SignCategory:     oldDocument.SignCategory,
		SignAgency:       oldDocument.SignAgency,
		SignText:         oldDocument.SignText + "@" + uuid.NewV4().String(),
		CategoryId:       oldDocument.CategoryId,
		AgencyId:         oldDocument.AgencyId,
		CreatedBy:        newCreatedBy,
		UpdatedBy:        newUpdatedBy,
		BeResponsibleBy:  newResponsibleBy,
		ViewCount:        newViewCount,
		DownloadCount:    newDownloadCount,
		Status:           oldDocument.Status,
		Type:             documentType,
		Priority:         oldDocument.Priority,
		ParentId:         parentId,
		BelongTo:         newBelongTo,
		CurrentId:        0,
		Path:             newPath,
		PublicToView:     oldDocument.PublicToView,
		PublicToDownload: oldDocument.PublicToDownload,
	}

	// 2. get list of field references
	oldFieldReferences := make([]*dms.DocumentFieldReferences, 0)

	err = global.GVA_DB.Model(&dms.DocumentFieldReferences{}).Where("document_id = ?", documentId).Find(&oldFieldReferences).Error
	if err != nil {
		return nil, err
	}

	// 3. get list of document relations
	oldDocumentRelations := make([]*dms.DocumentRelationReferences, 0)

	err = global.GVA_DB.Model(&dms.DocumentRelationReferences{}).Where("document_id = ?", documentId).Find(&oldDocumentRelations).Error
	if err != nil {
		return nil, err
	}

	// 4. get list of authorities
	oldAuthorities := make([]*dms.DocumentRules, 0)

	err = global.GVA_DB.Model(&dms.DocumentRules{}).Where("document_id = ?", documentId).Find(&oldAuthorities).Error
	if err != nil {
		return nil, err
	}

	// 5. get list of signers
	oldSigners := make([]*dms.DocumentSignerReferences, 0)

	err = global.GVA_DB.Model(&dms.DocumentSignerReferences{}).Where("document_id = ?", documentId).Find(&oldSigners).Error
	if err != nil {
		return nil, err
	}

	// 6. get list of attached users
	oldUsers := make([]*dms.DocumentUsers, 0)

	err = global.GVA_DB.Model(&dms.DocumentUsers{}).Where("document_id = ?", documentId).Find(&oldUsers).Error
	if err != nil {
		return nil, err
	}

	// 7. get list of attached files
	oldFiles := make([]*dms.DocumentFiles, 0)

	err = global.GVA_DB.Model(&dms.DocumentFiles{}).Where("document_id = ?", documentId).Find(&oldFiles).Error
	if err != nil {
		return nil, err
	}

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {

		// 1. duplicate document
		err = tx.Model(&dms.Documents{}).Create(&newDocument).Error
		if err != nil {
			return err
		}

		// 2. duplicate list of field references
		if len(oldFieldReferences) > 0 {
			for _, v := range oldFieldReferences {
				v.DocumentId = newDocument.ID
				v.ID = 0
				v.CreatedAt = time.Now()
				v.UpdatedAt = time.Now()
			}

			err = tx.Model(&dms.DocumentFieldReferences{}).Create(&oldFieldReferences).Error
			if err != nil {
				return err
			}
		}

		// 3. duplicate list of document relations
		if len(oldDocumentRelations) > 0 {
			for _, v := range oldDocumentRelations {
				v.DocumentId = newDocument.ID
				v.ID = 0
				v.CreatedAt = time.Now()
				v.UpdatedAt = time.Now()
			}

			err = tx.Model(&dms.DocumentRelationReferences{}).Create(&oldDocumentRelations).Error
			if err != nil {
				return err
			}
		}

		// 4. duplicate list of authorities
		if len(oldAuthorities) > 0 {
			if newDocument.Type == dms.TYPE_DOCUMENT {
				for _, v := range oldAuthorities {
					v.DocumentId = newDocument.ID
					v.ID = 0
					v.CreatedAt = time.Now()
					v.UpdatedAt = time.Now()
				}

				err = tx.Model(&dms.DocumentRules{}).Create(&oldAuthorities).Error
				if err != nil {
					return err
				}
			}
		}

		// 5. duplicate list of signers
		if len(oldSigners) > 0 {
			for _, v := range oldSigners {
				v.DocumentId = newDocument.ID
				v.ID = 0
				v.CreatedAt = time.Now()
				v.UpdatedAt = time.Now()
			}

			err = tx.Model(&dms.DocumentSignerReferences{}).Create(&oldSigners).Error
			if err != nil {
				return err
			}
		}

		// 6. duplicate list of attached users
		if len(oldUsers) > 0 {
			if newDocument.Type == dms.TYPE_DOCUMENT {
				for _, v := range oldUsers {
					v.DocumentId = newDocument.ID
					v.ID = 0
					v.CreatedAt = time.Now()
					v.UpdatedAt = time.Now()
				}

				err = tx.Model(&dms.DocumentUsers{}).Create(&oldUsers).Error
				if err != nil {
					return err
				}
			}
		}

		// 7. duplicate list of attached files
		if len(oldFiles) > 0 {
			for _, v := range oldFiles {
				v.DocumentId = newDocument.ID
				v.ID = 0
				v.CreatedAt = time.Now()
				v.UpdatedAt = time.Now()
			}

			err = tx.Model(&dms.DocumentFiles{}).Create(&oldFiles).Error
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &newDocument, nil
}

// ===================== INTERNAL METHODS SECTION ========================

func (documentsService *DocumentsService) attachBaseDocuments(document *dms.Documents) (err error, panicErr bool) {
	var refs []dms.DocumentRelationReferences
	db := global.GVA_DB.Model(&dms.DocumentRelationReferences{})

	err = db.Select("dest_id").Where("document_id = ? AND relation_type = ?", document.ID, "base").Find(&refs).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err, true
		}

		return err, false
	}

	var documentIds []uint
	for _, v := range refs {
		documentIds = append(documentIds, v.DestId)
	}

	db = global.GVA_DB.Model(&dms.Documents{})

	err = db.Find(&document.BasedDocuments, documentIds).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err, true
		}

		return err, false
	}

	return nil, false
}

func (documentsService *DocumentsService) attachReferencesDocuments(document *dms.Documents) (err error, panicErr bool) {
	var refs []dms.DocumentRelationReferences
	db := global.GVA_DB.Model(&dms.DocumentRelationReferences{})

	err = db.Select("dest_id").Where("document_id = ? AND relation_type = ?", document.ID, "relation").Find(&refs).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err, true
		}

		return err, false
	}

	var documentIds []uint
	for _, v := range refs {
		documentIds = append(documentIds, v.DestId)
	}

	db = global.GVA_DB.Model(&dms.Documents{})

	err = db.Find(&document.RelatedDocuments, documentIds).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err, true
		}

		return err, false
	}

	return nil, false
}

func (documentsService *DocumentsService) attachRelatedUsers(document *dms.Documents) (err error, panicErr bool) {
	var refs []dms.DocumentRelationReferences
	db := global.GVA_DB.Model(&dms.DocumentRelationReferences{})

	err = db.Select("dest_id").Where("document_id = ? AND relation_type = ?", document.ID, "user").Find(&refs).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err, true
		}

		return err, false
	}

	var documentIds []uint
	for _, v := range refs {
		documentIds = append(documentIds, v.DestId)
	}

	db = global.GVA_DB.Model(&sysModel.SysUser{})

	err = db.Find(&document.RelatedUsers, documentIds).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err, true
		}

		return err, false
	}

	return nil, false
}

func (documentsService *DocumentsService) attachRelatedAgencies(document *dms.Documents) (err error, panicErr bool) {
	var refs []dms.DocumentRelationReferences
	db := global.GVA_DB.Model(&dms.DocumentRelationReferences{})

	err = db.Select("dest_id").Where("document_id = ? AND relation_type = ?", document.ID, "agency").Find(&refs).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err, true
		}

		return err, false
	}

	var documentIds []uint
	for _, v := range refs {
		documentIds = append(documentIds, v.DestId)
	}

	db = global.GVA_DB.Model(&dms.DocumentAgencies{})

	err = db.Find(&document.RelatedAgencies, documentIds).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err, true
		}

		return err, false
	}

	return nil, false
}

func (documentsService *DocumentsService) attachCreatedUser(document *dms.Documents) (err error, panicErr bool) {
	db := global.GVA_DB.Model(&sysModel.SysUser{})
	var user sysModel.SysUser

	err = db.First(&user, "id = ?", document.CreatedBy).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err, true
		}

		return err, false
	}

	document.CreatedUser = &user

	return nil, false
}

func (documentsService *DocumentsService) attachUpdatedUser(document *dms.Documents) (err error, panicErr bool) {
	db := global.GVA_DB.Model(&sysModel.SysUser{})
	var user sysModel.SysUser

	err = db.First(&user, "id = ?", document.UpdatedBy).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, true
		}

		return err, false
	}

	document.UpdatedUser = &user

	return nil, false
}

func (documentsService *DocumentsService) attachResponsibleUser(document *dms.Documents) (err error, panicErr bool) {
	db := global.GVA_DB.Model(&sysModel.SysUser{})
	var user sysModel.SysUser

	err = db.First(&user, "id = ?", document.BeResponsibleBy).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, true
		}

		return err, false
	}

	document.ResponsibleUser = &user

	return nil, false
}

func (documentsService *DocumentsService) attachAuthority(document *dms.Documents) (err error, panicErr bool) {
	documentRuleDb := global.GVA_DB.Model(&dms.DocumentRules{})

	var documentRules []dms.DocumentRules

	err = documentRuleDb.Where("document_id = ?", document.ID).Find(&documentRules).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err, true
		}

		return err, false
	}

	documentAuthority := new(response.DocumentAuthority)
	documentAuthority.PublicToView = document.PublicToView
	documentAuthority.PublicToDownload = document.PublicToDownload

	viewLimitUsers := make([]sysModel.SysUser, 0)
	viewLimitRoles := make([]sysModel.SysAuthority, 0)
	downloadLimitUsers := make([]sysModel.SysUser, 0)
	downloadLimitRoles := make([]sysModel.SysAuthority, 0)
	updateLimitUsers := make([]sysModel.SysUser, 0)
	updateLimitRoles := make([]sysModel.SysAuthority, 0)
	ownerLimitUsers := make([]sysModel.SysUser, 0)
	ownerLimitRoles := make([]sysModel.SysAuthority, 0)

	for _, d := range documentRules {
		if d.Type == dms.RULE_TYPE_USER {
			var user sysModel.SysUser
			err = global.GVA_DB.Model(&sysModel.SysUser{}).First(&user, "id = ?", d.SubjectId).Error
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return err, true
				}

				return err, false
			}

			if d.Permission == dms.PERMISSION_VIEW {
				viewLimitUsers = append(viewLimitUsers, user)
			} else if d.Permission == dms.PERMISSION_DOWNLOAD {
				downloadLimitUsers = append(downloadLimitUsers, user)
			} else if d.Permission == dms.PERMISSION_EDIT {
				updateLimitUsers = append(updateLimitUsers, user)
			} else if d.Permission == dms.PERMISSION_OWNER {
				ownerLimitUsers = append(ownerLimitUsers, user)
			}
		} else if d.Type == dms.RULE_TYPE_GROUP {
			var role sysModel.SysAuthority
			err = global.GVA_DB.Model(sysModel.SysAuthority{}).First(&role, "authority_id = ?", d.SubjectId).Error
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return err, true
				}

				return err, false
			}

			if d.Permission == dms.PERMISSION_VIEW {
				viewLimitRoles = append(viewLimitRoles, role)
			} else if d.Permission == dms.PERMISSION_DOWNLOAD {
				downloadLimitRoles = append(downloadLimitRoles, role)
			} else if d.Permission == dms.PERMISSION_EDIT {
				updateLimitRoles = append(updateLimitRoles, role)
			} else if d.Permission == dms.PERMISSION_OWNER {
				ownerLimitRoles = append(ownerLimitRoles, role)
			}
		}
	}

	documentAuthority.ViewLimitUsers = viewLimitUsers
	documentAuthority.ViewLimitRoles = viewLimitRoles
	documentAuthority.DownloadLimitUsers = downloadLimitUsers
	documentAuthority.DownloadLimitRoles = downloadLimitRoles
	documentAuthority.UpdateLimitUsers = updateLimitUsers
	documentAuthority.UpdateLimitRoles = updateLimitRoles
	documentAuthority.OwnerLimitUsers = ownerLimitUsers
	documentAuthority.OwnerLimitRoles = ownerLimitRoles

	if len(viewLimitUsers)+len(viewLimitRoles) == 0 {
		documentAuthority.OnlyAdminCanView = true
	}

	if len(downloadLimitUsers)+len(downloadLimitRoles) == 0 {
		documentAuthority.OnlyAdminCanDownload = true
	}

	if len(updateLimitUsers)+len(updateLimitRoles) == 0 {
		documentAuthority.OnlyMeCanUpdate = true
	}

	if len(ownerLimitUsers)+len(ownerLimitRoles) == 1 {
		documentAuthority.OnlyMeBeOwner = true
	}

	document.Authority = documentAuthority

	return nil, false
}

// ===================== END OF INTERNAL METHODS SECTION ========================
