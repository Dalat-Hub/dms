package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
	dmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/dms/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DocumentsApi struct {
}

var documentsService = service.ServiceGroupApp.DmsServiceGroup.DocumentsService

// CreateDocuments create new document
// @Tags Documents
// @Summary create new document
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.Documents true "create new document"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/createDocuments [post]
func (documentsApi *DocumentsApi) CreateDocuments(c *gin.Context) {
	var documents dms.Documents
	var err error

	err = c.ShouldBindJSON(&documents)
	if err != nil {
		global.GVA_LOG.Error("provide valid document", zap.Error(err))
		response.FailWithMessage("provide valid document", c)
	}

	if err = documentsService.CreateDocuments(documents); err != nil {
		global.GVA_LOG.Error("fail to create new document", zap.Error(err))
		response.FailWithMessage("fail to create new document", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// CreateDraftDocument create new document
// @Tags Documents
// @Summary create new draft document
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dmsReq.DraftDocument true "create new draft document"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/createDocuments [post]
func (documentsApi *DocumentsApi) CreateDraftDocument(c *gin.Context) {
	var draft dmsReq.DraftDocument
	var err error

	err = c.ShouldBindJSON(&draft)
	if err != nil {
		global.GVA_LOG.Error("provide valid document", zap.Error(err))
		response.FailWithMessage("provide valid document", c)
	}

	if doc, err := documentsService.CreateDraftDocument(draft); err != nil {
		global.GVA_LOG.Error("fail to create new draft document", zap.Error(err))
		response.FailWithMessage("fail to create new draft document", c)
	} else {
		response.OkWithData(gin.H{"document": doc}, c)
	}
}

// CreateFullDocument create new full document
// @Tags Documents
// @Summary create new full document
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dmsReq.FullDocument true "create new full document"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/createDocuments [post]
func (documentsApi *DocumentsApi) CreateFullDocument(c *gin.Context) {
	var full dmsReq.FullDocument
	var err error

	err = c.ShouldBindJSON(&full)
	if err != nil {
		global.GVA_LOG.Error("provide valid document", zap.Error(err))
		response.FailWithMessage("provide valid document", c)
	}

	if doc, err := documentsService.CreateFullDocument(full); err != nil {
		global.GVA_LOG.Error("fail to create new full document", zap.Error(err))
		response.FailWithMessage("fail to create new full document", c)
	} else {
		response.OkWithData(gin.H{"document": doc}, c)
	}
}

// DeleteDocuments delete document by ID
// @Tags Documents
// @Summary delete document by ID
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.Documents true "delete document by ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/deleteDocuments [delete]
func (documentsApi *DocumentsApi) DeleteDocuments(c *gin.Context) {
	var documents dms.Documents
	var err error

	err = c.ShouldBindJSON(&documents)
	if err != nil {
		global.GVA_LOG.Error("provide valid ID", zap.Error(err))
		response.FailWithMessage("provide valid ID", c)
	}

	if err := documentsService.DeleteDocuments(documents); err != nil {
		global.GVA_LOG.Error("fail to delete document", zap.Error(err))
		response.FailWithMessage("fail to delete document", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// DeleteDocumentsByIds bulk delete document by IDs
// @Tags Documents
// @Summary bulk delete document by IDs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "bulk delete document by IDs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/deleteDocumentsByIds [delete]
func (documentsApi *DocumentsApi) DeleteDocumentsByIds(c *gin.Context) {
	var IDS request.IdsReq
	var err error

	err = c.ShouldBindJSON(&IDS)
	if err != nil {
		global.GVA_LOG.Error("provide valid document", zap.Error(err))
		response.FailWithMessage("provide valid document", c)
	}

	if err = documentsService.DeleteDocumentsByIds(IDS); err != nil {
		global.GVA_LOG.Error("fail to create", zap.Error(err))
		response.FailWithMessage("fail to create", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// UpdateDocuments update document
// @Tags Documents
// @Summary update document
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.Documents true "update document"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/updateDocuments [put]
func (documentsApi *DocumentsApi) UpdateDocuments(c *gin.Context) {
	var documents dms.Documents
	var err error

	err = c.ShouldBindJSON(&documents)
	if err != nil {
		global.GVA_LOG.Error("provide valid updated document", zap.Error(err))
		response.FailWithMessage("provide valid updated document", c)
	}

	if err = documentsService.UpdateDocuments(documents); err != nil {
		global.GVA_LOG.Error("fail to update document", zap.Error(err))
		response.FailWithMessage("fail to update document", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// FindDocuments find document by ID
// @Tags Documents
// @Summary find document by ID
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dms.Documents true "find document by ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/findDocuments [get]
func (documentsApi *DocumentsApi) FindDocuments(c *gin.Context) {
	var documents dmsReq.DocumentsSearch
	var err error

	err = c.ShouldBindQuery(&documents)
	if err != nil {
		global.GVA_LOG.Error("provide valid ID", zap.Error(err))
		response.FailWithMessage("provide valid ID", c)
	}

	user := utils.GetUserInfo(c)
	if err = documentRulesService.CheckPermission(user.ID, user.UUID, documents.ID, dms.PERMISSION_VIEW); err != nil {
		global.GVA_LOG.Error("forbiden", zap.Error(err))
		response.FailWithMessage("forbiden", c)
		return
	}

	if redocuments, err := documentsService.GetDocuments(documents, user.ID, user.UUID); err != nil {
		global.GVA_LOG.Error("fail to find document", zap.Error(err))
		response.FailWithMessage("fail to find document", c)
	} else {
		response.OkWithData(gin.H{"document": redocuments}, c)
	}
}

// GetDocumentsList get list of documents
// @Tags Documents
// @Summary get list of documents
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dmsReq.DocumentsSearch true "get list of documents"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/getDocumentsList [get]
func (documentsApi *DocumentsApi) GetDocumentsList(c *gin.Context) {
	var pageInfo dmsReq.DocumentsSearch
	var err error

	err = c.ShouldBindQuery(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("provide valid search params", zap.Error(err))
		response.FailWithMessage("provide valid search params", c)
	}

	if list, total, err := documentsService.GetDocumentsInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("fail to get list of documents", zap.Error(err))
		response.FailWithMessage("fail to get list of documents", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "success", c)
	}
}
