package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
	dmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/dms/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DocumentCategoriesApi struct {
}

var documentCategoriesService = service.ServiceGroupApp.DmsServiceGroup.DocumentCategoriesService

// CreateDocumentCategories Create new category
// @Tags DocumentCategories
// @Summary create new category
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentCategories true "create new category"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentCategories/createDocumentCategories [post]
func (documentCategoriesApi *DocumentCategoriesApi) CreateDocumentCategories(c *gin.Context) {
	var documentCategories dms.DocumentCategories
	var err error

	err = c.ShouldBindJSON(&documentCategories)
	if err != nil {
		global.GVA_LOG.Error("provide valid category", zap.Error(err))
		response.FailWithMessage("provide valid category", c)
	}

	if err := documentCategoriesService.CreateDocumentCategories(&documentCategories); err != nil {
		global.GVA_LOG.Error("fail to create new category", zap.Error(err))
		response.FailWithMessage("fail to create new category", c)
	} else {
		response.OkWithData(gin.H{"category": documentCategories}, c)
	}
}

// DeleteDocumentCategories delete category by ID
// @Tags DocumentCategories
// @Summary delete category by ID
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentCategories true "delete category by ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentCategories/deleteDocumentCategories [delete]
func (documentCategoriesApi *DocumentCategoriesApi) DeleteDocumentCategories(c *gin.Context) {
	var documentCategories dms.DocumentCategories
	var err error

	err = c.ShouldBindJSON(&documentCategories)
	if err != nil {
		global.GVA_LOG.Error("provide valid ID", zap.Error(err))
		response.FailWithMessage("provide valid ID", c)
	}

	if err := documentCategoriesService.DeleteDocumentCategories(documentCategories); err != nil {
		global.GVA_LOG.Error("fail to delete category", zap.Error(err))
		response.FailWithMessage("fail to delete category", c)
	} else {
		response.OkWithMessage("delete category successfully", c)
	}
}

// DeleteDocumentCategoriesByIds bulk delete category by IDs
// @Tags DocumentCategories
// @Summary bulk delete category by IDs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "bulk delete category by IDs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentCategories/deleteDocumentCategoriesByIds [delete]
func (documentCategoriesApi *DocumentCategoriesApi) DeleteDocumentCategoriesByIds(c *gin.Context) {
	var IDS request.IdsReq
	var err error

	err = c.ShouldBindJSON(&IDS)
	if err != nil {
		global.GVA_LOG.Error("provide valid ID array", zap.Error(err))
		response.FailWithMessage("provide valid ID array", c)
	}

	if err := documentCategoriesService.DeleteDocumentCategoriesByIds(IDS); err != nil {
		global.GVA_LOG.Error("fail to bulk delete categories", zap.Error(err))
		response.FailWithMessage("fail to bulk delete categories", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// UpdateDocumentCategories update category
// @Tags DocumentCategories
// @Summary update category
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentCategories true "update category"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentCategories/updateDocumentCategories [put]
func (documentCategoriesApi *DocumentCategoriesApi) UpdateDocumentCategories(c *gin.Context) {
	var documentCategories dms.DocumentCategories
	var err error

	err = c.ShouldBindJSON(&documentCategories)
	if err != nil {
		global.GVA_LOG.Error("provide valid updated category", zap.Error(err))
		response.FailWithMessage("provide valid updated category", c)
	}

	if err := documentCategoriesService.UpdateDocumentCategories(documentCategories); err != nil {
		global.GVA_LOG.Error("fail to update category", zap.Error(err))
		response.FailWithMessage("fail to update category", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// FindDocumentCategories find category by ID
// @Tags DocumentCategories
// @Summary find category by ID
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dms.DocumentCategories true "find category by ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentCategories/findDocumentCategories [get]
func (documentCategoriesApi *DocumentCategoriesApi) FindDocumentCategories(c *gin.Context) {
	var documentCategories dms.DocumentCategories
	var err error

	err = c.ShouldBindQuery(&documentCategories)
	if err != nil {
		global.GVA_LOG.Error("provide valid ID to find", zap.Error(err))
		response.FailWithMessage("provide valid ID to find", c)
	}

	if redocumentCategories, err := documentCategoriesService.GetDocumentCategories(documentCategories.ID); err != nil {
		global.GVA_LOG.Error("fail to find category", zap.Error(err))
		response.FailWithMessage("fail to find category", c)
	} else {
		response.OkWithData(gin.H{"redocumentCategories": redocumentCategories}, c)
	}
}

// GetDocumentCategoriesList get categories list
// @Tags DocumentCategories
// @Summary get categories list
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dmsReq.DocumentCategoriesSearch true "get categories list"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentCategories/getDocumentCategoriesList [get]
func (documentCategoriesApi *DocumentCategoriesApi) GetDocumentCategoriesList(c *gin.Context) {
	var pageInfo dmsReq.DocumentCategoriesSearch
	var err error

	err = c.ShouldBindQuery(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("provide valid search information", zap.Error(err))
		response.FailWithMessage("provide valid search information", c)
	}

	if list, total, err := documentCategoriesService.GetDocumentCategoriesInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("fail to get list of categories", zap.Error(err))
		response.FailWithMessage("fail to get list of categories", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "success", c)
	}
}

// GetDocumentCategoriesListPublic get categories list
// @Tags DocumentCategories
// @Summary get categories list
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dmsReq.DocumentCategoriesSearch true "get categories list"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /api/v1/categories [get]
func (documentCategoriesApi *DocumentCategoriesApi) GetDocumentCategoriesListPublic(c *gin.Context) {
	var pageInfo dmsReq.DocumentCategoriesSearch
	var err error

	err = c.ShouldBindQuery(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("provide valid search information", zap.Error(err))
		response.FailWithMessage("provide valid search information", c)
	}

	if list, total, err := documentCategoriesService.GetDocumentCategoriesInfoListPublic(pageInfo); err != nil {
		global.GVA_LOG.Error("fail to get list of categories", zap.Error(err))
		response.FailWithMessage("fail to get list of categories", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "success", c)
	}
}
