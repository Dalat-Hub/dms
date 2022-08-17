package dms

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
	dmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/dms/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	sysService "github.com/flipped-aurora/gin-vue-admin/server/service/system"
	uuid "github.com/satori/go.uuid"
)

type DocumentRulesService struct {
}

// CreateDocumentRules create new rule
func (documentRulesService *DocumentRulesService) CreateDocumentRules(documentRules dms.DocumentRules) (err error) {
	err = global.GVA_DB.Create(&documentRules).Error
	return err
}

// DeleteDocumentRules delete rule by ID
func (documentRulesService *DocumentRulesService) DeleteDocumentRules(documentRules dms.DocumentRules) (err error) {
	err = global.GVA_DB.Delete(&documentRules).Error
	return err
}

// DeleteDocumentRulesByIds bulk delete rule by IDs
func (documentRulesService *DocumentRulesService) DeleteDocumentRulesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dms.DocumentRules{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateDocumentRules update rule
func (documentRulesService *DocumentRulesService) UpdateDocumentRules(documentRules dms.DocumentRules) (err error) {
	err = global.GVA_DB.Save(&documentRules).Error
	return err
}

// GetDocumentRules get rule by ID
func (documentRulesService *DocumentRulesService) GetDocumentRules(id uint) (documentRules dms.DocumentRules, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&documentRules).Error
	return
}

// GetDocumentRulesInfoList get list of rules
func (documentRulesService *DocumentRulesService) GetDocumentRulesInfoList(info dmsReq.DocumentRulesSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&dms.DocumentRules{})
	var documentRuless []dms.DocumentRules

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&documentRuless).Error
	return documentRuless, total, err
}

// CheckPermissionFactory check user permission
func (documentRulesService *DocumentRulesService) CheckPermissionFactory() func(userId uint, userUuid uuid.UUID, documentId uint, permission string) (err error) {
	userAuthority := map[uint]map[uint]string{}
	userInfoMap := map[uint]system.SysUser{}

	return func(userId uint, userUuid uuid.UUID, documentId uint, permission string) (err error) {
		if userAuthority == nil {
			userAuthority = map[uint]map[uint]string{}
		}

		db := global.GVA_DB.Model(&dms.DocumentRules{})
		var documentRules []dms.DocumentRules

		err = db.Where("document_id = ?", documentId).Find(&documentRules).Error
		if err != nil {
			return err
		}

		sysUserService := new(sysService.UserService)

		if _, exist := userInfoMap[userId]; !exist {
			u, authErr := sysUserService.GetUserInfo(userUuid)

			if authErr != nil {
				return authErr
			}

			userInfoMap[userId] = u
		}

		user := userInfoMap[userId]

		if _, exist := userAuthority[userId]; !exist {
			userAuthority[userId] = map[uint]string{}

			authorityService := new(sysService.AuthorityService)

			authorityMap := map[uint]string{}

			for _, authority := range user.Authorities {
				err, auMap := authorityService.GetChildrenAuthorities(authority)
				if err != nil {
					return err
				}

				for k, v := range auMap {
					authorityMap[k] = v
				}
			}

			userAuthority[userId] = authorityMap
		}

		for _, rule := range documentRules {
			if rule.Type == dms.RULE_TYPE_USER {
				if rule.SubjectId == user.ID {
					if rule.Permission == dms.PERMISSION_OWNER || rule.Permission == permission {
						return nil
					}
				}
			} else if rule.Type == dms.RULE_TYPE_GROUP {
				if _, ok := userAuthority[userId][rule.SubjectId]; ok {
					if rule.Permission == dms.PERMISSION_OWNER || rule.Permission == permission {
						return nil
					}
				}
			}
		}

		return errors.New("not found")
	}
}

// CheckPermission check user permission
func (documentRulesService *DocumentRulesService) CheckPermission(userId uint, userUuid uuid.UUID, documentId uint, permission string) (err error) {
	db := global.GVA_DB.Model(&dms.DocumentRules{})
	var documentRules []dms.DocumentRules

	err = db.Where("document_id = ?", documentId).Find(&documentRules).Error
	if err != nil {
		return err
	}

	sysUserService := new(sysService.UserService)
	user, authErr := sysUserService.GetUserInfo(userUuid)
	if authErr != nil {
		return authErr
	}

	authorityIds := make(map[uint]string)
	for _, authority := range user.Authorities {
		authorityIds[authority.AuthorityId] = authority.AuthorityName
	}

	for _, rule := range documentRules {
		if rule.Type == dms.RULE_TYPE_USER {
			if rule.SubjectId == user.ID {
				if rule.Permission == dms.PERMISSION_OWNER || rule.Permission == permission {
					return nil
				}
			}
		} else if rule.Type == dms.RULE_TYPE_GROUP {
			if _, ok := authorityIds[rule.SubjectId]; ok {
				if rule.Permission == dms.PERMISSION_OWNER || rule.Permission == permission {
					return nil
				}
			}
		}
	}

	return errors.New("not found")
}
