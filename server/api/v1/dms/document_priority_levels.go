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

type DocumentPriorityLevelsApi struct {
}

var documentPriorityLevelsService = service.ServiceGroupApp.DmsServiceGroup.DocumentPriorityLevelsService

// CreateDocumentPriorityLevels create new document priority level
// @Tags DocumentPriorityLevels
// @Summary create new document priority level
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentPriorityLevels true "create new document priority level"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentPriorityLevels/createDocumentPriorityLevels [post]
func (documentPriorityLevelsApi *DocumentPriorityLevelsApi) CreateDocumentPriorityLevels(c *gin.Context) {
	var documentPriorityLevels dms.DocumentPriorityLevels
	var err error

	err = c.ShouldBindJSON(&documentPriorityLevels)
	if err != nil {
		global.GVA_LOG.Error("provide valid data", zap.Error(err))
		response.FailWithMessage("provide valid data", c)
	}

	if err = documentPriorityLevelsService.CreateDocumentPriorityLevels(documentPriorityLevels); err != nil {
		global.GVA_LOG.Error("fail to create", zap.Error(err))
		response.FailWithMessage("fail to create", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// DeleteDocumentPriorityLevels delete priority level by ID
// @Tags DocumentPriorityLevels
// @Summary delete priority level by ID
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentPriorityLevels true "delete priority level by ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentPriorityLevels/deleteDocumentPriorityLevels [delete]
func (documentPriorityLevelsApi *DocumentPriorityLevelsApi) DeleteDocumentPriorityLevels(c *gin.Context) {
	var documentPriorityLevels dms.DocumentPriorityLevels
	var err error

	err = c.ShouldBindJSON(&documentPriorityLevels)
	if err != nil {
		global.GVA_LOG.Error("provide valid ID", zap.Error(err))
		response.FailWithMessage("provide valid ID", c)
	}

	if err = documentPriorityLevelsService.DeleteDocumentPriorityLevels(documentPriorityLevels); err != nil {
		global.GVA_LOG.Error("fail to delete", zap.Error(err))
		response.FailWithMessage("fail to delete", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// DeleteDocumentPriorityLevelsByIds bulk delete priority level
// @Tags DocumentPriorityLevels
// @Summary bulk delete priority level
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "bulk delete priority level"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentPriorityLevels/deleteDocumentPriorityLevelsByIds [delete]
func (documentPriorityLevelsApi *DocumentPriorityLevelsApi) DeleteDocumentPriorityLevelsByIds(c *gin.Context) {
	var IDS request.IdsReq
	var err error

	err = c.ShouldBindJSON(&IDS)
	if err != nil {
		global.GVA_LOG.Error("provide valid ID array", zap.Error(err))
		response.FailWithMessage("provide valid ID array", c)
	}

	if err = documentPriorityLevelsService.DeleteDocumentPriorityLevelsByIds(IDS); err != nil {
		global.GVA_LOG.Error("fail to bulk delete", zap.Error(err))
		response.FailWithMessage("fail to bulk delete", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// UpdateDocumentPriorityLevels update priority level
// @Tags DocumentPriorityLevels
// @Summary update priority level
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentPriorityLevels true "update priority level"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentPriorityLevels/updateDocumentPriorityLevels [put]
func (documentPriorityLevelsApi *DocumentPriorityLevelsApi) UpdateDocumentPriorityLevels(c *gin.Context) {
	var documentPriorityLevels dms.DocumentPriorityLevels
	var err error

	err = c.ShouldBindJSON(&documentPriorityLevels)
	if err != nil {
		global.GVA_LOG.Error("provide valid updated priority level", zap.Error(err))
		response.FailWithMessage("provide valid updated priority level", c)
	}

	if err = documentPriorityLevelsService.UpdateDocumentPriorityLevels(documentPriorityLevels); err != nil {
		global.GVA_LOG.Error("fail to update", zap.Error(err))
		response.FailWithMessage("fail to update", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// FindDocumentPriorityLevels find priority level by ID
// @Tags DocumentPriorityLevels
// @Summary find priority level by ID
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dms.DocumentPriorityLevels true "fail priority level by ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentPriorityLevels/findDocumentPriorityLevels [get]
func (documentPriorityLevelsApi *DocumentPriorityLevelsApi) FindDocumentPriorityLevels(c *gin.Context) {
	var documentPriorityLevels dms.DocumentPriorityLevels
	var err error

	err = c.ShouldBindQuery(&documentPriorityLevels)
	if err != nil {
		global.GVA_LOG.Error("provide valid ID", zap.Error(err))
		response.FailWithMessage("provide valid ID", c)
	}

	if redocumentPriorityLevels, err := documentPriorityLevelsService.GetDocumentPriorityLevels(documentPriorityLevels.ID); err != nil {
		global.GVA_LOG.Error("fail to find priority", zap.Error(err))
		response.FailWithMessage("fail to find priority", c)
	} else {
		response.OkWithData(gin.H{"redocumentPriorityLevels": redocumentPriorityLevels}, c)
	}
}

// GetDocumentPriorityLevelsList get list of priorities
// @Tags DocumentPriorityLevels
// @Summary get list of priorities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dmsReq.DocumentPriorityLevelsSearch true "get list of priorities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentPriorityLevels/getDocumentPriorityLevelsList [get]
func (documentPriorityLevelsApi *DocumentPriorityLevelsApi) GetDocumentPriorityLevelsList(c *gin.Context) {
	var pageInfo dmsReq.DocumentPriorityLevelsSearch
	var err error

	err = c.ShouldBindQuery(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("provide valid search params", zap.Error(err))
		response.FailWithMessage("provide valid search params", c)
	}

	if list, total, err := documentPriorityLevelsService.GetDocumentPriorityLevelsInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("fail to get list of priorities", zap.Error(err))
		response.FailWithMessage("fail to get list of priorities", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "success", c)
	}
}
