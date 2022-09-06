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

	var documentCategoriess []*dms.DocumentCategories

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&documentCategoriess).Error
	if err != nil {
		return
	}

	err = documentCategoriesService.attachDocumentCount(documentCategoriess)
	if err != nil {
		return
	}

	return documentCategoriess, total, err
}

// GetDocumentCategoriesInfoListPublic get list of categories
func (documentCategoriesService *DocumentCategoriesService) GetDocumentCategoriesInfoListPublic(info dmsReq.DocumentCategoriesSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&dms.DocumentCategories{})
	db = db.Where("`hidden` = ?", false)
	db = db.Order("`order` desc")

	var documentCategoriess []*dms.DocumentCategories

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	db = db.Limit(limit)
	db = db.Offset(offset)
	db = db.Order("`order` desc").Debug()

	err = db.Find(&documentCategoriess).Error
	if err != nil {
		return
	}

	err = documentCategoriesService.attachDocumentCount(documentCategoriess)
	if err != nil {
		return
	}

	return documentCategoriess, total, err
}

type categoryStats struct {
	CategoryId uint `json:"category_id"`
	Count      uint `json:"count"`
}

func (documentCategoriesService *DocumentCategoriesService) attachDocumentCount(categories []*dms.DocumentCategories) (err error) {
	results := make([]categoryStats, 0)

	err = global.GVA_DB.Model(&dms.Documents{}).
		Select("category_id, count(id) as count").
		Where("type = ? AND status = ?", dms.TYPE_DOCUMENT, dms.STATUS_PUBLISHED).
		Group("category_id").
		Find(&results).Error

	if err != nil {
		return err
	}

	resultsAsMap := map[uint]uint{}
	for _, stat := range results {
		resultsAsMap[stat.CategoryId] = stat.Count
	}

	for _, category := range categories {
		if val, ok := resultsAsMap[category.ID]; ok {
			category.Count = int(val)
		}
	}

	return nil
}
