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

type DocumentAgenciesApi struct {
}

var documentAgenciesService = service.ServiceGroupApp.DmsServiceGroup.DocumentAgenciesService

// CreateDocumentAgencies Create new agency
// @Tags DocumentAgencies
// @Summary create new agency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentAgencies true "create new agency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"create success"}"
// @Router /documentAgencies/createDocumentAgencies [post]
func (documentAgenciesApi *DocumentAgenciesApi) CreateDocumentAgencies(c *gin.Context) {
	var documentAgencies dms.DocumentAgencies
	var err error
	err = c.ShouldBindJSON(&documentAgencies)

	if err != nil {
		global.GVA_LOG.Error("fail to create new agency! check body data", zap.Error(err))
		response.FailWithMessage("fail to create new agency", c)
	}

	if err = documentAgenciesService.CreateDocumentAgencies(&documentAgencies); err != nil {
		global.GVA_LOG.Error("fail to create new agency!", zap.Error(err))
		response.FailWithMessage("fail to create new agency", c)
	} else {
		response.OkWithData(gin.H{"agency": documentAgencies}, c)
	}
}

// DeleteDocumentAgencies Delete agency by ID
// @Tags DocumentAgencies
// @Summary Delete agency by ID
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentAgencies true "Delete agency by ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"delete successfully"}"
// @Router /documentAgencies/deleteDocumentAgencies [delete]
func (documentAgenciesApi *DocumentAgenciesApi) DeleteDocumentAgencies(c *gin.Context) {
	var documentAgencies dms.DocumentAgencies
	var err error
	err = c.ShouldBindJSON(&documentAgencies)

	if err != nil {
		global.GVA_LOG.Error("please provide agency ID", zap.Error(err))
		response.FailWithMessage("please provide agency ID", c)
	}

	if err = documentAgenciesService.DeleteDocumentAgencies(documentAgencies); err != nil {
		global.GVA_LOG.Error("fail to delete agency", zap.Error(err))
		response.FailWithMessage("fail to delete agency", c)
	} else {
		response.OkWithMessage("delete agency successfully", c)
	}
}

// DeleteDocumentAgenciesByIds Bulk delete document agencies
// @Tags DocumentAgencies
// @Summary Bulk delete document agencies
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "bulk delete document agencies"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"delete successfully"}"
// @Router /documentAgencies/deleteDocumentAgenciesByIds [delete]
func (documentAgenciesApi *DocumentAgenciesApi) DeleteDocumentAgenciesByIds(c *gin.Context) {
	var IDS request.IdsReq
	var err error

	err = c.ShouldBindJSON(&IDS)
	if err != nil {
		global.GVA_LOG.Error("invalid data, please provide an array of ID to delete", zap.Error(err))
		response.FailWithMessage("invalid data, please provide an array of ID to delete", c)
	}

	if err = documentAgenciesService.DeleteDocumentAgenciesByIds(IDS); err != nil {
		global.GVA_LOG.Error("fail to bulk delete agencies", zap.Error(err))
		response.FailWithMessage("fail to bulk delete agencies", c)
	} else {
		response.OkWithMessage("bulk delete agencies successfully", c)
	}
}

// UpdateDocumentAgencies Update document agency infomation
// @Tags DocumentAgencies
// @Summary update document agency information
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentAgencies true "update document agency information"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update document agency information"}"
// @Router /documentAgencies/updateDocumentAgencies [put]
func (documentAgenciesApi *DocumentAgenciesApi) UpdateDocumentAgencies(c *gin.Context) {
	var documentAgencies dms.DocumentAgencies
	var err error

	err = c.ShouldBindJSON(&documentAgencies)
	if err != nil {
		global.GVA_LOG.Error("please provide valid data to update", zap.Error(err))
		response.FailWithMessage("please provide valid data to update", c)
	}

	if err := documentAgenciesService.UpdateDocumentAgencies(documentAgencies); err != nil {
		global.GVA_LOG.Error("fail to update agency", zap.Error(err))
		response.FailWithMessage("fail to update agency", c)
	} else {
		response.OkWithMessage("update agency successfully", c)
	}
}

// FindDocumentAgencies find agency by ID
// @Tags DocumentAgencies
// @Summary find agency by ID
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dms.DocumentAgencies true "find agency by ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"agency"}"
// @Router /documentAgencies/findDocumentAgencies [get]
func (documentAgenciesApi *DocumentAgenciesApi) FindDocumentAgencies(c *gin.Context) {
	var documentAgencies dms.DocumentAgencies
	var err error

	err = c.ShouldBindQuery(&documentAgencies)
	if err != nil {
		global.GVA_LOG.Error("please provide valid agency ID to find", zap.Error(err))
		response.FailWithMessage("please provide valid agency ID to find", c)
	}

	if redocumentAgencies, err := documentAgenciesService.GetDocumentAgencies(documentAgencies.ID); err != nil {
		global.GVA_LOG.Error("fail to find agency", zap.Error(err))
		response.FailWithMessage("fail to find agency", c)
	} else {
		response.OkWithData(gin.H{"redocumentAgencies": redocumentAgencies}, c)
	}
}

// GetDocumentAgenciesList paginate to get list of agencies
// @Tags DocumentAgencies
// @Summary paginate to get list of agencies
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dmsReq.DocumentAgenciesSearch true "paginate to get list of agencies"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"list of agencies"}"
// @Router /documentAgencies/getDocumentAgenciesList [get]
func (documentAgenciesApi *DocumentAgenciesApi) GetDocumentAgenciesList(c *gin.Context) {
	var pageInfo dmsReq.DocumentAgenciesSearch
	var err error

	_ = c.ShouldBindQuery(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("please provide valid data to get list of agencies", zap.Error(err))
		response.FailWithMessage("please provide valid data to get list of agencies", c)
	}

	if list, total, err := documentAgenciesService.GetDocumentAgenciesInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("fail to get list of agencies", zap.Error(err))
		response.FailWithMessage("fail to get list of agencies", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "list of agencies", c)
	}
}
