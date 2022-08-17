package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
	dmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/dms/request"
	reqDocuments "github.com/flipped-aurora/gin-vue-admin/server/model/dms/request/documents"
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

	user := utils.GetUserInfo(c)

	if doc, err := documentsService.CreateDraftDocument(draft, user.ID); err != nil {
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

// UpdateBasicDocumentInformation update basic document information
// @Tags Documents
// @Summary update basic document information
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.Documents true "update basic document information"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/updateBasicDocumentInformation [put]
func (documentsApi *DocumentsApi) UpdateBasicDocumentInformation(c *gin.Context) {
	var basic reqDocuments.UpdateBasic
	var err error

	err = c.ShouldBindJSON(&basic)
	if err != nil {
		global.GVA_LOG.Error("provide valid updated document", zap.Error(err))
		response.FailWithMessage("provide valid updated document", c)
		return
	}

	if err = documentsService.UpdateBasicDocumentInformation(basic); err != nil {
		global.GVA_LOG.Error("fail to update document", zap.Error(err))
		response.FailWithMessage("fail to update document", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// UpdateRelatedDocumentInformation update related document information
// @Tags Documents
// @Summary update related document information
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.Documents true "update related document information"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/updateRelatedDocumentInformation [put]
func (documentsApi *DocumentsApi) UpdateRelatedDocumentInformation(c *gin.Context) {
	var relation reqDocuments.UpdateRelation
	var err error

	err = c.ShouldBindJSON(&relation)
	if err != nil {
		global.GVA_LOG.Error("provide valid updated document", zap.Error(err))
		response.FailWithMessage("provide valid updated document", c)
		return
	}

	if err = documentsService.UpdateRelationDocumentInformation(relation); err != nil {
		global.GVA_LOG.Error("fail to update document", zap.Error(err))
		response.FailWithMessage("fail to update document", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// UpdateDocumentFiles update document attached files
// @Tags Documents
// @Summary update document attached files
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.Documents true "update attached document files"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/updateDocumentFiles [put]
func (documentsApi *DocumentsApi) UpdateDocumentFiles(c *gin.Context) {
	var file reqDocuments.UpdateFile
	var err error

	err = c.ShouldBindJSON(&file)
	if err != nil {
		global.GVA_LOG.Error("provide valid updated document", zap.Error(err))
		response.FailWithMessage("provide valid updated document", c)
		return
	}

	if err = documentsService.UpdateDocumentAttachedFiles(file); err != nil {
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
		return
	}

	user := utils.GetUserInfo(c)
	if err = documentRulesService.CheckPermission(user.ID, user.UUID, documents.ID, dms.PERMISSION_VIEW); err != nil {
		global.GVA_LOG.Error("you don't have permission to view the desired document", zap.Error(err))
		response.FailWithMessage("you don't have permission to view the desired document", c)
		return
	}

	if redocuments, err := documentsService.GetDocuments(documents, user.ID, user.UUID); err != nil {
		global.GVA_LOG.Error("fail to find document", zap.Error(err))
		response.FailWithMessage("fail to find document", c)
	} else {
		response.OkWithData(gin.H{"document": redocuments}, c)
	}
}

func (documentsApi *DocumentsApi) FindDocumentsPublic(c *gin.Context) {
	var documents dmsReq.DocumentsSearch
	var err error

	err = c.ShouldBindQuery(&documents)
	if err != nil {
		global.GVA_LOG.Error("provide valid ID", zap.Error(err))
		response.FailWithMessage("provide valid ID", c)
		return
	}

	if redocuments, err := documentsService.GetDocumentsPublic(documents); err != nil {
		global.GVA_LOG.Error("fail to find document", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
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
		return
	}

	if list, total, err := documentsService.GetDocumentsInfoList(pageInfo, false); err != nil {
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

func (documentsApi *DocumentsApi) GetDocumentsListPublic(c *gin.Context) {
	var pageInfo reqDocuments.PublicSearch
	var err error

	err = c.ShouldBindQuery(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("provide valid search params", zap.Error(err))
		response.FailWithMessage("provide valid search params", c)
		return
	}

	userInfo := utils.GetUserInfo(c)

	if list, total, err := documentsService.GetDocumentsInfoListPublic(pageInfo, userInfo); err != nil {
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

// GetDocumentRevisions get list of revision of the given document
// @Tags Documents
// @Summary get list of revision of the given document
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.GetByID true "get list of revision of the given document"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/getDocumentRevisions [get]
func (documentsApi *DocumentsApi) GetDocumentRevisions(c *gin.Context) {
	var info request.GetById
	var err error

	err = c.ShouldBindQuery(&info)
	if err != nil {
		global.GVA_LOG.Error("provide valid search params", zap.Error(err))
		response.FailWithMessage("provide valid search params", c)
		return
	}

	if list, err := documentsService.GetDocumentRevisions(uint(info.ID)); err != nil {
		global.GVA_LOG.Error("fail to get list of revisions", zap.Error(err))
		response.FailWithMessage("fail to get list of revisions", c)
	} else {
		response.OkWithData(gin.H{"revisions": list}, c)
	}
}

// GetFileList get list of files attached to the document
// @Tags Documents
// @Summary get list of files attached to the document
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dmsReq.DocumentsSearch true "get list of files attached to the document"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/getFileList [get]
func (documentsApi *DocumentsApi) GetFileList(c *gin.Context) {
	var searchInfo request.GetById
	var err error

	err = c.ShouldBindQuery(&searchInfo)
	if err != nil {
		global.GVA_LOG.Error("provide valid search params", zap.Error(err))
		response.FailWithMessage("provide valid search params", c)
		return
	}

	canDownloadFile := true
	user := utils.GetUserInfo(c)

	if err = documentRulesService.CheckPermission(user.ID, user.UUID, uint(searchInfo.ID), dms.PERMISSION_DOWNLOAD); err != nil {
		canDownloadFile = false
	}

	if files, cDownload, err := documentsService.GetDocumentFiles(searchInfo, canDownloadFile); err != nil {
		global.GVA_LOG.Error("fail to get list of documents", zap.Error(err))
		response.FailWithMessage("fail to get list of documents", c)
	} else {
		response.OkWithData(gin.H{"files": files, "canDownload": cDownload}, c)
	}
}

func (documentsApi *DocumentsApi) GetFileListPublic(c *gin.Context) {
	var searchInfo request.GetById
	var err error

	err = c.ShouldBindQuery(&searchInfo)
	if err != nil {
		global.GVA_LOG.Error("provide valid search params", zap.Error(err))
		response.FailWithMessage("provide valid search params", c)
		return
	}

	if files, cDownload, err := documentsService.GetDocumentFilesPublic(searchInfo); err != nil {
		global.GVA_LOG.Error("fail to get list of documents", zap.Error(err))
		response.FailWithMessage("fail to get list of documents", c)
	} else {
		response.OkWithData(gin.H{"files": files, "canDownload": cDownload}, c)
	}
}

// MakeDuplication create new duplicated version of the given document
// @Tags Documents
// @Summary create new duplicated version of the given document
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.GetByID true "create new duplicated version of the given document"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/makeDuplication [post]
func (documentsApi *DocumentsApi) MakeDuplication(c *gin.Context) {
	var info request.GetById
	var err error

	err = c.ShouldBindQuery(&info)
	if err != nil {
		global.GVA_LOG.Error("provide valid search params", zap.Error(err))
		response.FailWithMessage("provide valid search params", c)
		return
	}

	user := utils.GetUserInfo(c)

	if backup, err := documentsService.Duplicate(uint(info.ID), user.ID, dms.TYPE_DOCUMENT); err != nil {
		global.GVA_LOG.Error("fail to create duplication", zap.Error(err))
		response.FailWithMessage("fail to create duplication", c)
	} else {
		response.OkWithData(gin.H{"duplication": backup}, c)
	}
}
