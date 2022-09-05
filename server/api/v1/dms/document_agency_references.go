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

type DocumentAgencyReferencesApi struct {
}

var documentAgencyReferencesService = service.ServiceGroupApp.DmsServiceGroup.DocumentAgencyReferencesService

// CreateDocumentAgencyReferences 创建DocumentAgencyReferences
// @Tags DocumentAgencyReferences
// @Summary 创建DocumentAgencyReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentAgencyReferences true "创建DocumentAgencyReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /documentAgencyReferences/createDocumentAgencyReferences [post]
func (documentAgencyReferencesApi *DocumentAgencyReferencesApi) CreateDocumentAgencyReferences(c *gin.Context) {
	var documentAgencyReferences dms.DocumentAgencyReferences
	_ = c.ShouldBindJSON(&documentAgencyReferences)
	if err := documentAgencyReferencesService.CreateDocumentAgencyReferences(documentAgencyReferences); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDocumentAgencyReferences 删除DocumentAgencyReferences
// @Tags DocumentAgencyReferences
// @Summary 删除DocumentAgencyReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentAgencyReferences true "删除DocumentAgencyReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /documentAgencyReferences/deleteDocumentAgencyReferences [delete]
func (documentAgencyReferencesApi *DocumentAgencyReferencesApi) DeleteDocumentAgencyReferences(c *gin.Context) {
	var documentAgencyReferences dms.DocumentAgencyReferences
	_ = c.ShouldBindJSON(&documentAgencyReferences)
	if err := documentAgencyReferencesService.DeleteDocumentAgencyReferences(documentAgencyReferences); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDocumentAgencyReferencesByIds 批量删除DocumentAgencyReferences
// @Tags DocumentAgencyReferences
// @Summary 批量删除DocumentAgencyReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DocumentAgencyReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /documentAgencyReferences/deleteDocumentAgencyReferencesByIds [delete]
func (documentAgencyReferencesApi *DocumentAgencyReferencesApi) DeleteDocumentAgencyReferencesByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := documentAgencyReferencesService.DeleteDocumentAgencyReferencesByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDocumentAgencyReferences 更新DocumentAgencyReferences
// @Tags DocumentAgencyReferences
// @Summary 更新DocumentAgencyReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentAgencyReferences true "更新DocumentAgencyReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /documentAgencyReferences/updateDocumentAgencyReferences [put]
func (documentAgencyReferencesApi *DocumentAgencyReferencesApi) UpdateDocumentAgencyReferences(c *gin.Context) {
	var documentAgencyReferences dms.DocumentAgencyReferences
	_ = c.ShouldBindJSON(&documentAgencyReferences)
	if err := documentAgencyReferencesService.UpdateDocumentAgencyReferences(documentAgencyReferences); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDocumentAgencyReferences 用id查询DocumentAgencyReferences
// @Tags DocumentAgencyReferences
// @Summary 用id查询DocumentAgencyReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dms.DocumentAgencyReferences true "用id查询DocumentAgencyReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /documentAgencyReferences/findDocumentAgencyReferences [get]
func (documentAgencyReferencesApi *DocumentAgencyReferencesApi) FindDocumentAgencyReferences(c *gin.Context) {
	var documentAgencyReferences dms.DocumentAgencyReferences
	_ = c.ShouldBindQuery(&documentAgencyReferences)
	if redocumentAgencyReferences, err := documentAgencyReferencesService.GetDocumentAgencyReferences(documentAgencyReferences.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redocumentAgencyReferences": redocumentAgencyReferences}, c)
	}
}

// GetDocumentAgencyReferencesList 分页获取DocumentAgencyReferences列表
// @Tags DocumentAgencyReferences
// @Summary 分页获取DocumentAgencyReferences列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dmsReq.DocumentAgencyReferencesSearch true "分页获取DocumentAgencyReferences列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /documentAgencyReferences/getDocumentAgencyReferencesList [get]
func (documentAgencyReferencesApi *DocumentAgencyReferencesApi) GetDocumentAgencyReferencesList(c *gin.Context) {
	var pageInfo dmsReq.DocumentAgencyReferencesSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := documentAgencyReferencesService.GetDocumentAgencyReferencesInfoList(pageInfo); err != nil {
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
