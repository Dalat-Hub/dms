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

type DocumentFieldsApi struct {
}

var documentFieldsService = service.ServiceGroupApp.DmsServiceGroup.DocumentFieldsService

// CreateDocumentFields create field
// @Tags DocumentFields
// @Summary create field
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentFields true "create field"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentFields/createDocumentFields [post]
func (documentFieldsApi *DocumentFieldsApi) CreateDocumentFields(c *gin.Context) {
	var documentFields dms.DocumentFields
	var err error

	err = c.ShouldBindJSON(&documentFields)
	if err != nil {
		global.GVA_LOG.Error("provide valid data", zap.Error(err))
		response.FailWithMessage("provide valid data", c)
	}

	if err := documentFieldsService.CreateDocumentFields(documentFields); err != nil {
		global.GVA_LOG.Error("fail to create new field", zap.Error(err))
		response.FailWithMessage("fail to create new field", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// DeleteDocumentFields delete field by ID
// @Tags DocumentFields
// @Summary delete field by ID
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentFields true "delete field"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentFields/deleteDocumentFields [delete]
func (documentFieldsApi *DocumentFieldsApi) DeleteDocumentFields(c *gin.Context) {
	var documentFields dms.DocumentFields
	var err error

	err = c.ShouldBindJSON(&documentFields)
	if err != nil {
		global.GVA_LOG.Error("provide valid ID to delete", zap.Error(err))
		response.FailWithMessage("provide valid ID to delete", c)
	}

	if err = documentFieldsService.DeleteDocumentFields(documentFields); err != nil {
		global.GVA_LOG.Error("fail to delete field", zap.Error(err))
		response.FailWithMessage("fail to delete field", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// DeleteDocumentFieldsByIds bulk delete field by IDs
// @Tags DocumentFields
// @Summary bulk delete field by IDs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "bulk delete field by IDs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentFields/deleteDocumentFieldsByIds [delete]
func (documentFieldsApi *DocumentFieldsApi) DeleteDocumentFieldsByIds(c *gin.Context) {
	var IDS request.IdsReq
	var err error

	_ = c.ShouldBindJSON(&IDS)
	if err != nil {
		global.GVA_LOG.Error("provide valid Id array", zap.Error(err))
		response.FailWithMessage("provide valid Id array", c)
	}

	if err = documentFieldsService.DeleteDocumentFieldsByIds(IDS); err != nil {
		global.GVA_LOG.Error("fail to bulk delete", zap.Error(err))
		response.FailWithMessage("fail to bulk delete", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// UpdateDocumentFields update field
// @Tags DocumentFields
// @Summary update field
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentFields true "update field"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentFields/updateDocumentFields [put]
func (documentFieldsApi *DocumentFieldsApi) UpdateDocumentFields(c *gin.Context) {
	var documentFields dms.DocumentFields
	var err error

	err = c.ShouldBindJSON(&documentFields)
	if err != nil {
		global.GVA_LOG.Error("provide valid updated field", zap.Error(err))
		response.FailWithMessage("provide valid updated field", c)
	}

	if err := documentFieldsService.UpdateDocumentFields(documentFields); err != nil {
		global.GVA_LOG.Error("fail to update", zap.Error(err))
		response.FailWithMessage("fail to update", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// FindDocumentFields find field by ID
// @Tags DocumentFields
// @Summary find field by ID
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dms.DocumentFields true "find field by ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentFields/findDocumentFields [get]
func (documentFieldsApi *DocumentFieldsApi) FindDocumentFields(c *gin.Context) {
	var documentFields dms.DocumentFields
	var err error

	err = c.ShouldBindQuery(&documentFields)
	if err != nil {
		global.GVA_LOG.Error("provide valid field ID", zap.Error(err))
		response.FailWithMessage("provide valid field ID", c)
	}

	if redocumentFields, err := documentFieldsService.GetDocumentFields(documentFields.ID); err != nil {
		global.GVA_LOG.Error("fail to find", zap.Error(err))
		response.FailWithMessage("fail to find", c)
	} else {
		response.OkWithData(gin.H{"redocumentFields": redocumentFields}, c)
	}
}

// GetDocumentFieldsList get list of fields
// @Tags DocumentFields
// @Summary get list of fields
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dmsReq.DocumentFieldsSearch true "get list of fields"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentFields/getDocumentFieldsList [get]
func (documentFieldsApi *DocumentFieldsApi) GetDocumentFieldsList(c *gin.Context) {
	var pageInfo dmsReq.DocumentFieldsSearch
	var err error

	err = c.ShouldBindQuery(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("provide valid search information", zap.Error(err))
		response.FailWithMessage("provide valid search information", c)
	}

	if list, total, err := documentFieldsService.GetDocumentFieldsInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("find to get list of fields", zap.Error(err))
		response.FailWithMessage("find to get list of fields", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "success", c)
	}
}
