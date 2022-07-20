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

type DocumentRelationReferencesApi struct {
}

var documentRelationReferencesService = service.ServiceGroupApp.DmsServiceGroup.DocumentRelationReferencesService

// CreateDocumentRelationReferences 创建DocumentRelationReferences
// @Tags DocumentRelationReferences
// @Summary 创建DocumentRelationReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentRelationReferences true "创建DocumentRelationReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /documentRelationReferences/createDocumentRelationReferences [post]
func (documentRelationReferencesApi *DocumentRelationReferencesApi) CreateDocumentRelationReferences(c *gin.Context) {
	var documentRelationReferences dms.DocumentRelationReferences
	_ = c.ShouldBindJSON(&documentRelationReferences)
	if err := documentRelationReferencesService.CreateDocumentRelationReferences(documentRelationReferences); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDocumentRelationReferences 删除DocumentRelationReferences
// @Tags DocumentRelationReferences
// @Summary 删除DocumentRelationReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentRelationReferences true "删除DocumentRelationReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /documentRelationReferences/deleteDocumentRelationReferences [delete]
func (documentRelationReferencesApi *DocumentRelationReferencesApi) DeleteDocumentRelationReferences(c *gin.Context) {
	var documentRelationReferences dms.DocumentRelationReferences
	_ = c.ShouldBindJSON(&documentRelationReferences)
	if err := documentRelationReferencesService.DeleteDocumentRelationReferences(documentRelationReferences); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDocumentRelationReferencesByIds 批量删除DocumentRelationReferences
// @Tags DocumentRelationReferences
// @Summary 批量删除DocumentRelationReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DocumentRelationReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /documentRelationReferences/deleteDocumentRelationReferencesByIds [delete]
func (documentRelationReferencesApi *DocumentRelationReferencesApi) DeleteDocumentRelationReferencesByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := documentRelationReferencesService.DeleteDocumentRelationReferencesByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDocumentRelationReferences 更新DocumentRelationReferences
// @Tags DocumentRelationReferences
// @Summary 更新DocumentRelationReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentRelationReferences true "更新DocumentRelationReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /documentRelationReferences/updateDocumentRelationReferences [put]
func (documentRelationReferencesApi *DocumentRelationReferencesApi) UpdateDocumentRelationReferences(c *gin.Context) {
	var documentRelationReferences dms.DocumentRelationReferences
	_ = c.ShouldBindJSON(&documentRelationReferences)
	if err := documentRelationReferencesService.UpdateDocumentRelationReferences(documentRelationReferences); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDocumentRelationReferences 用id查询DocumentRelationReferences
// @Tags DocumentRelationReferences
// @Summary 用id查询DocumentRelationReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dms.DocumentRelationReferences true "用id查询DocumentRelationReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /documentRelationReferences/findDocumentRelationReferences [get]
func (documentRelationReferencesApi *DocumentRelationReferencesApi) FindDocumentRelationReferences(c *gin.Context) {
	var documentRelationReferences dms.DocumentRelationReferences
	_ = c.ShouldBindQuery(&documentRelationReferences)
	if redocumentRelationReferences, err := documentRelationReferencesService.GetDocumentRelationReferences(documentRelationReferences.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redocumentRelationReferences": redocumentRelationReferences}, c)
	}
}

// GetDocumentRelationReferencesList 分页获取DocumentRelationReferences列表
// @Tags DocumentRelationReferences
// @Summary 分页获取DocumentRelationReferences列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dmsReq.DocumentRelationReferencesSearch true "分页获取DocumentRelationReferences列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /documentRelationReferences/getDocumentRelationReferencesList [get]
func (documentRelationReferencesApi *DocumentRelationReferencesApi) GetDocumentRelationReferencesList(c *gin.Context) {
	var pageInfo dmsReq.DocumentRelationReferencesSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := documentRelationReferencesService.GetDocumentRelationReferencesInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
