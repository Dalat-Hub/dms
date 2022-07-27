package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
	dmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/dms/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms/response"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
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
		err = documentsService.attachBaseDocuments(&document)
		if err != nil {
			return nil, err
		}
	}

	if doc.PreloadRelatedDocs == 1 {
		err = documentsService.attachReferencesDocuments(&document)
		if err != nil {
			return nil, err
		}
	}

	if doc.PreloadRelatedUsers == 1 {
		err = documentsService.attachRelatedUsers(&document)
		if err != nil {
			return nil, err
		}
	}

	if doc.PreloadRelatedAgencies == 1 {
		err = documentsService.attachRelatedAgencies(&document)
		if err != nil {
			return nil, err
		}
	}

	if doc.PreloadCreatedBy == 1 {
		err = documentsService.attachCreatedUser(&document)
		if err != nil {
			return nil, err
		}
	}

	if doc.PreloadUpdatedBy == 1 {
		err = documentsService.attachUpdatedUser(&document)
		if err != nil {
			return nil, err
		}
	}

	if doc.PreloadBeResponsibleBy == 1 {
		err = documentsService.attachResponsibleUser(&document)
		if err != nil {
			return nil, err
		}
	}

	if doc.PreloadAuthority == 1 {
		service := new(DocumentRulesService)
		if err := service.CheckPermission(userId, userUUID, document.ID, dms.PERMISSION_OWNER); err == nil {
			err = documentsService.attachAuthority(&document)
			if err != nil {
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

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&documentss).Error
	return documentss, total, err
}

func (documentsService *DocumentsService) attachBaseDocuments(document *dms.Documents) (err error) {
	var refs []dms.DocumentRelationReferences
	db := global.GVA_DB.Model(&dms.DocumentRelationReferences{})

	err = db.Select("dest_id").Where("document_id = ? AND relation_type = ?", document.ID, "base").Find(&refs).Error
	if err != nil {
		return nil
	}

	var documentIds []uint
	for _, v := range refs {
		documentIds = append(documentIds, v.DestId)
	}

	db = global.GVA_DB.Model(&dms.Documents{})

	err = db.Find(&document.BasedDocuments, documentIds).Error
	if err != nil {
		return nil
	}

	return nil
}

func (documentsService *DocumentsService) attachReferencesDocuments(document *dms.Documents) (err error) {
	var refs []dms.DocumentRelationReferences
	db := global.GVA_DB.Model(&dms.DocumentRelationReferences{})

	err = db.Select("dest_id").Where("document_id = ? AND relation_type = ?", document.ID, "relation").Find(&refs).Error
	if err != nil {
		return nil
	}

	var documentIds []uint
	for _, v := range refs {
		documentIds = append(documentIds, v.DestId)
	}

	db = global.GVA_DB.Model(&dms.Documents{})

	err = db.Find(&document.RelatedDocuments, documentIds).Error
	if err != nil {
		return nil
	}

	return nil
}

func (documentsService *DocumentsService) attachRelatedUsers(document *dms.Documents) (err error) {
	var refs []dms.DocumentRelationReferences
	db := global.GVA_DB.Model(&dms.DocumentRelationReferences{})

	err = db.Select("dest_id").Where("document_id = ? AND relation_type = ?", document.ID, "user").Find(&refs).Error
	if err != nil {
		return nil
	}

	var documentIds []uint
	for _, v := range refs {
		documentIds = append(documentIds, v.DestId)
	}

	db = global.GVA_DB.Model(&sysModel.SysUser{})

	err = db.Find(&document.RelatedUsers, documentIds).Error
	if err != nil {
		return nil
	}

	return nil
}

func (documentsService *DocumentsService) attachRelatedAgencies(document *dms.Documents) (err error) {
	var refs []dms.DocumentRelationReferences
	db := global.GVA_DB.Model(&dms.DocumentRelationReferences{})

	err = db.Select("dest_id").Where("document_id = ? AND relation_type = ?", document.ID, "agency").Find(&refs).Error
	if err != nil {
		return nil
	}

	var documentIds []uint
	for _, v := range refs {
		documentIds = append(documentIds, v.DestId)
	}

	db = global.GVA_DB.Model(&dms.DocumentAgencies{})

	err = db.Find(&document.RelatedAgencies, documentIds).Error
	if err != nil {
		return nil
	}

	return nil
}

func (documentsService *DocumentsService) attachCreatedUser(document *dms.Documents) (err error) {
	db := global.GVA_DB.Model(&sysModel.SysUser{})
	var user sysModel.SysUser

	err = db.First(&user, "id = ?", document.CreatedBy).Error
	if err != nil {
		return err
	}

	document.CreatedUser = &user

	return nil
}

func (documentsService *DocumentsService) attachUpdatedUser(document *dms.Documents) (err error) {
	db := global.GVA_DB.Model(&sysModel.SysUser{})
	var user sysModel.SysUser

	err = db.First(&user, "id = ?", document.UpdatedBy).Error
	if err != nil {
		return err
	}

	document.UpdatedUser = &user

	return nil
}

func (documentsService *DocumentsService) attachResponsibleUser(document *dms.Documents) (err error) {
	db := global.GVA_DB.Model(&sysModel.SysUser{})
	var user sysModel.SysUser

	err = db.First(&user, "id = ?", document.BeResponsibleBy).Error
	if err != nil {
		return err
	}

	document.ResponsibleUser = &user

	return nil
}

func (documentsService *DocumentsService) attachAuthority(document *dms.Documents) (err error) {
	documentRuleDb := global.GVA_DB.Model(&dms.DocumentRules{})

	var documentRules []dms.DocumentRules

	err = documentRuleDb.Where("document_id = ?", document.ID).Find(&documentRules).Error
	if err != nil {
		return err
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
				return err
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
				return err
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

	return nil
}
