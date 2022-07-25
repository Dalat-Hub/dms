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

type DocumentSignerReferencesApi struct {
}

var documentSignerReferencesService = service.ServiceGroupApp.DmsServiceGroup.DocumentSignerReferencesService

// CreateDocumentSignerReferences 创建DocumentSignerReferences
// @Tags DocumentSignerReferences
// @Summary 创建DocumentSignerReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentSignerReferences true "创建DocumentSignerReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /documentSignerReferences/createDocumentSignerReferences [post]
func (documentSignerReferencesApi *DocumentSignerReferencesApi) CreateDocumentSignerReferences(c *gin.Context) {
	var documentSignerReferences dms.DocumentSignerReferences
	_ = c.ShouldBindJSON(&documentSignerReferences)
	if err := documentSignerReferencesService.CreateDocumentSignerReferences(documentSignerReferences); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDocumentSignerReferences 删除DocumentSignerReferences
// @Tags DocumentSignerReferences
// @Summary 删除DocumentSignerReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentSignerReferences true "删除DocumentSignerReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /documentSignerReferences/deleteDocumentSignerReferences [delete]
func (documentSignerReferencesApi *DocumentSignerReferencesApi) DeleteDocumentSignerReferences(c *gin.Context) {
	var documentSignerReferences dms.DocumentSignerReferences
	_ = c.ShouldBindJSON(&documentSignerReferences)
	if err := documentSignerReferencesService.DeleteDocumentSignerReferences(documentSignerReferences); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDocumentSignerReferencesByIds 批量删除DocumentSignerReferences
// @Tags DocumentSignerReferences
// @Summary 批量删除DocumentSignerReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DocumentSignerReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /documentSignerReferences/deleteDocumentSignerReferencesByIds [delete]
func (documentSignerReferencesApi *DocumentSignerReferencesApi) DeleteDocumentSignerReferencesByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := documentSignerReferencesService.DeleteDocumentSignerReferencesByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDocumentSignerReferences 更新DocumentSignerReferences
// @Tags DocumentSignerReferences
// @Summary 更新DocumentSignerReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentSignerReferences true "更新DocumentSignerReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /documentSignerReferences/updateDocumentSignerReferences [put]
func (documentSignerReferencesApi *DocumentSignerReferencesApi) UpdateDocumentSignerReferences(c *gin.Context) {
	var documentSignerReferences dms.DocumentSignerReferences
	_ = c.ShouldBindJSON(&documentSignerReferences)
	if err := documentSignerReferencesService.UpdateDocumentSignerReferences(documentSignerReferences); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDocumentSignerReferences 用id查询DocumentSignerReferences
// @Tags DocumentSignerReferences
// @Summary 用id查询DocumentSignerReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dms.DocumentSignerReferences true "用id查询DocumentSignerReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /documentSignerReferences/findDocumentSignerReferences [get]
func (documentSignerReferencesApi *DocumentSignerReferencesApi) FindDocumentSignerReferences(c *gin.Context) {
	var documentSignerReferences dms.DocumentSignerReferences
	_ = c.ShouldBindQuery(&documentSignerReferences)
	if redocumentSignerReferences, err := documentSignerReferencesService.GetDocumentSignerReferences(documentSignerReferences.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redocumentSignerReferences": redocumentSignerReferences}, c)
	}
}

// GetDocumentSignerReferencesList 分页获取DocumentSignerReferences列表
// @Tags DocumentSignerReferences
// @Summary 分页获取DocumentSignerReferences列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dmsReq.DocumentSignerReferencesSearch true "分页获取DocumentSignerReferences列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /documentSignerReferences/getDocumentSignerReferencesList [get]
func (documentSignerReferencesApi *DocumentSignerReferencesApi) GetDocumentSignerReferencesList(c *gin.Context) {
	var pageInfo dmsReq.DocumentSignerReferencesSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := documentSignerReferencesService.GetDocumentSignerReferencesInfoList(pageInfo); err != nil {
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
