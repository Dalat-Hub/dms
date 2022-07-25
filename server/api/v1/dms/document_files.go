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

type DocumentFilesApi struct {
}

var documentFilesService = service.ServiceGroupApp.DmsServiceGroup.DocumentFilesService

// CreateDocumentFiles create new file
// @Tags DocumentFiles
// @Summary create new file
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentFiles true "create new file"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentFiles/createDocumentFiles [post]
func (documentFilesApi *DocumentFilesApi) CreateDocumentFiles(c *gin.Context) {
	var documentFiles dms.DocumentFiles
	var err error

	err = c.ShouldBindJSON(&documentFiles)
	if err != nil {
		global.GVA_LOG.Error("provide valid data", zap.Error(err))
		response.FailWithMessage("provide valid data", c)
	}

	if err := documentFilesService.CreateDocumentFiles(documentFiles); err != nil {
		global.GVA_LOG.Error("fail to create new file", zap.Error(err))
		response.FailWithMessage("fail to create new file", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// DeleteDocumentFiles 删除DocumentFiles
// @Tags DocumentFiles
// @Summary 删除DocumentFiles
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentFiles true "删除DocumentFiles"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /documentFiles/deleteDocumentFiles [delete]
func (documentFilesApi *DocumentFilesApi) DeleteDocumentFiles(c *gin.Context) {
	var documentFiles dms.DocumentFiles
	_ = c.ShouldBindJSON(&documentFiles)
	if err := documentFilesService.DeleteDocumentFiles(documentFiles); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDocumentFilesByIds 批量删除DocumentFiles
// @Tags DocumentFiles
// @Summary 批量删除DocumentFiles
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DocumentFiles"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /documentFiles/deleteDocumentFilesByIds [delete]
func (documentFilesApi *DocumentFilesApi) DeleteDocumentFilesByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := documentFilesService.DeleteDocumentFilesByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDocumentFiles 更新DocumentFiles
// @Tags DocumentFiles
// @Summary 更新DocumentFiles
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentFiles true "更新DocumentFiles"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /documentFiles/updateDocumentFiles [put]
func (documentFilesApi *DocumentFilesApi) UpdateDocumentFiles(c *gin.Context) {
	var documentFiles dms.DocumentFiles
	_ = c.ShouldBindJSON(&documentFiles)
	if err := documentFilesService.UpdateDocumentFiles(documentFiles); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDocumentFiles 用id查询DocumentFiles
// @Tags DocumentFiles
// @Summary 用id查询DocumentFiles
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dms.DocumentFiles true "用id查询DocumentFiles"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /documentFiles/findDocumentFiles [get]
func (documentFilesApi *DocumentFilesApi) FindDocumentFiles(c *gin.Context) {
	var documentFiles dms.DocumentFiles
	_ = c.ShouldBindQuery(&documentFiles)
	if redocumentFiles, err := documentFilesService.GetDocumentFiles(documentFiles.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redocumentFiles": redocumentFiles}, c)
	}
}

// GetDocumentFilesList 分页获取DocumentFiles列表
// @Tags DocumentFiles
// @Summary 分页获取DocumentFiles列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dmsReq.DocumentFilesSearch true "分页获取DocumentFiles列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /documentFiles/getDocumentFilesList [get]
func (documentFilesApi *DocumentFilesApi) GetDocumentFilesList(c *gin.Context) {
	var pageInfo dmsReq.DocumentFilesSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := documentFilesService.GetDocumentFilesInfoList(pageInfo); err != nil {
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
