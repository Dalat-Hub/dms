package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
	dmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/dms/request"
)

type DocumentCategoriesService struct {
}

// CreateDocumentCategories create new category
func (documentCategoriesService *DocumentCategoriesService) CreateDocumentCategories(documentCategories *dms.DocumentCategories) (err error) {
	err = global.GVA_DB.Create(&documentCategories).Error
	return err
}

// DeleteDocumentCategories delete category by ID
func (documentCategoriesService *DocumentCategoriesService) DeleteDocumentCategories(documentCategories dms.DocumentCategories) (err error) {
	err = global.GVA_DB.Delete(&documentCategories).Error
	return err
}

// DeleteDocumentCategoriesByIds bulk delete category by IDs
func (documentCategoriesService *DocumentCategoriesService) DeleteDocumentCategoriesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dms.DocumentCategories{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateDocumentCategories update category
func (documentCategoriesService *DocumentCategoriesService) UpdateDocumentCategories(documentCategories dms.DocumentCategories) (err error) {
	err = global.GVA_DB.Save(&documentCategories).Error
	return err
}

// GetDocumentCategories get category by ID
func (documentCategoriesService *DocumentCategoriesService) GetDocumentCategories(id uint) (documentCategories dms.DocumentCategories, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&documentCategories).Error
	return
}

// GetDocumentCategoriesInfoList get list of categories
func (documentCategoriesService *DocumentCategoriesService) GetDocumentCategoriesInfoList(info dmsReq.DocumentCategoriesSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&dms.DocumentCategories{})

	var documentCategoriess []dms.DocumentCategories

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&documentCategoriess).Error
	return documentCategoriess, total, err
}
