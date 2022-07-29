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

type DocumentFieldReferencesApi struct {
}

var documentFieldReferencesService = service.ServiceGroupApp.DmsServiceGroup.DocumentFieldReferencesService

// CreateDocumentFieldReferences create new pivot between document & field
// @Tags DocumentFieldReferences
// @Summary create new pivot between document & field
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentFieldReferences true "create new pivot between document & field"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentFieldReferences/createDocumentFieldReferences [post]
func (documentFieldReferencesApi *DocumentFieldReferencesApi) CreateDocumentFieldReferences(c *gin.Context) {
	var documentFieldReferences dms.DocumentFieldReferences
	var err error

	err = c.ShouldBindJSON(&documentFieldReferences)
	if err != nil {
		global.GVA_LOG.Error("provide valid data", zap.Error(err))
		response.FailWithMessage("provide valid data", c)
	}

	if err = documentFieldReferencesService.CreateDocumentFieldReferences(documentFieldReferences); err != nil {
		global.GVA_LOG.Error("fail to create pivot!", zap.Error(err))
		response.FailWithMessage("fail to create pivot", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// DeleteDocumentFieldReferences 删除DocumentFieldReferences
// @Tags DocumentFieldReferences
// @Summary 删除DocumentFieldReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentFieldReferences true "删除DocumentFieldReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /documentFieldReferences/deleteDocumentFieldReferences [delete]
func (documentFieldReferencesApi *DocumentFieldReferencesApi) DeleteDocumentFieldReferences(c *gin.Context) {
	var documentFieldReferences dms.DocumentFieldReferences
	_ = c.ShouldBindJSON(&documentFieldReferences)
	if err := documentFieldReferencesService.DeleteDocumentFieldReferences(documentFieldReferences); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDocumentFieldReferencesByIds 批量删除DocumentFieldReferences
// @Tags DocumentFieldReferences
// @Summary 批量删除DocumentFieldReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DocumentFieldReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /documentFieldReferences/deleteDocumentFieldReferencesByIds [delete]
func (documentFieldReferencesApi *DocumentFieldReferencesApi) DeleteDocumentFieldReferencesByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := documentFieldReferencesService.DeleteDocumentFieldReferencesByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDocumentFieldReferences 更新DocumentFieldReferences
// @Tags DocumentFieldReferences
// @Summary 更新DocumentFieldReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentFieldReferences true "更新DocumentFieldReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /documentFieldReferences/updateDocumentFieldReferences [put]
func (documentFieldReferencesApi *DocumentFieldReferencesApi) UpdateDocumentFieldReferences(c *gin.Context) {
	var documentFieldReferences dms.DocumentFieldReferences
	_ = c.ShouldBindJSON(&documentFieldReferences)
	if err := documentFieldReferencesService.UpdateDocumentFieldReferences(documentFieldReferences); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDocumentFieldReferences 用id查询DocumentFieldReferences
// @Tags DocumentFieldReferences
// @Summary 用id查询DocumentFieldReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dms.DocumentFieldReferences true "用id查询DocumentFieldReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /documentFieldReferences/findDocumentFieldReferences [get]
func (documentFieldReferencesApi *DocumentFieldReferencesApi) FindDocumentFieldReferences(c *gin.Context) {
	var documentFieldReferences dms.DocumentFieldReferences
	_ = c.ShouldBindQuery(&documentFieldReferences)
	if redocumentFieldReferences, err := documentFieldReferencesService.GetDocumentFieldReferences(documentFieldReferences.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redocumentFieldReferences": redocumentFieldReferences}, c)
	}
}

// GetDocumentFieldReferencesList 分页获取DocumentFieldReferences列表
// @Tags DocumentFieldReferences
// @Summary 分页获取DocumentFieldReferences列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dmsReq.DocumentFieldReferencesSearch true "分页获取DocumentFieldReferences列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /documentFieldReferences/getDocumentFieldReferencesList [get]
func (documentFieldReferencesApi *DocumentFieldReferencesApi) GetDocumentFieldReferencesList(c *gin.Context) {
	var pageInfo dmsReq.DocumentFieldReferencesSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := documentFieldReferencesService.GetDocumentFieldReferencesInfoList(pageInfo); err != nil {
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
