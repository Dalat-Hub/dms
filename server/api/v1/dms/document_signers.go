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

type DocumentSignersApi struct {
}

var documentSignersService = service.ServiceGroupApp.DmsServiceGroup.DocumentSignersService

// CreateDocumentSigners create new signer
// @Tags DocumentSigners
// @Summary create new signer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentSigners true "create new signer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentSigners/createDocumentSigners [post]
func (documentSignersApi *DocumentSignersApi) CreateDocumentSigners(c *gin.Context) {
	var documentSigners dms.DocumentSigners
	var err error

	err = c.ShouldBindJSON(&documentSigners)
	if err != nil {
		global.GVA_LOG.Error("provide valid data to create", zap.Error(err))
		response.FailWithMessage("provide valid data to create", c)
		return
	}

	if err = documentSignersService.CreateDocumentSigners(&documentSigners); err != nil {
		global.GVA_LOG.Error("fail to create new signer", zap.Error(err))
		response.FailWithMessage("fail to create new signer", c)
	} else {
		response.OkWithData(gin.H{"signer": documentSigners}, c)
	}
}

// DeleteDocumentSigners delete signer
// @Tags DocumentSigners
// @Summary delete signer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentSigners true "delete signer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentSigners/deleteDocumentSigners [delete]
func (documentSignersApi *DocumentSignersApi) DeleteDocumentSigners(c *gin.Context) {
	var documentSigners dms.DocumentSigners
	var err error

	err = c.ShouldBindJSON(&documentSigners)
	if err != nil {
		global.GVA_LOG.Error("provide valid data", zap.Error(err))
		response.FailWithMessage("provide valid data", c)
		return
	}

	if err = documentSignersService.DeleteDocumentSigners(documentSigners); err != nil {
		global.GVA_LOG.Error("fail to delete signer", zap.Error(err))
		response.FailWithMessage("fail to delete signer", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// DeleteDocumentSignersByIds bulk delete signers by IDs
// @Tags DocumentSigners
// @Summary bulk delete signers by IDs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "bulk delete signers by IDs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentSigners/deleteDocumentSignersByIds [delete]
func (documentSignersApi *DocumentSignersApi) DeleteDocumentSignersByIds(c *gin.Context) {
	var IDS request.IdsReq
	var err error

	err = c.ShouldBindJSON(&IDS)
	if err != nil {
		global.GVA_LOG.Error("provide valid data", zap.Error(err))
		response.FailWithMessage("provide valid data", c)
		return
	}

	if err = documentSignersService.DeleteDocumentSignersByIds(IDS); err != nil {
		global.GVA_LOG.Error("fail to bulk delete", zap.Error(err))
		response.FailWithMessage("fail to bulk delete", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// UpdateDocumentSigners update signer
// @Tags DocumentSigners
// @Summary update signer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentSigners true "update signers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentSigners/updateDocumentSigners [put]
func (documentSignersApi *DocumentSignersApi) UpdateDocumentSigners(c *gin.Context) {
	var documentSigners dms.DocumentSigners
	var err error

	err = c.ShouldBindJSON(&documentSigners)
	if err != nil {
		global.GVA_LOG.Error("provide valid data", zap.Error(err))
		response.FailWithMessage("provide valid data", c)
		return
	}

	if err = documentSignersService.UpdateDocumentSigners(documentSigners); err != nil {
		global.GVA_LOG.Error("fail to update", zap.Error(err))
		response.FailWithMessage("fail to update", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// FindDocumentSigners find signer
// @Tags DocumentSigners
// @Summary find signer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dms.DocumentSigners true "find signer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentSigners/findDocumentSigners [get]
func (documentSignersApi *DocumentSignersApi) FindDocumentSigners(c *gin.Context) {
	var documentSigners dms.DocumentSigners
	var err error

	err = c.ShouldBindQuery(&documentSigners)
	if err != nil {
		global.GVA_LOG.Error("provide valid data", zap.Error(err))
		response.FailWithMessage("provide valid data", c)
		return
	}

	if redocumentSigners, err := documentSignersService.GetDocumentSigners(documentSigners.ID); err != nil {
		global.GVA_LOG.Error("fail to find", zap.Error(err))
		response.FailWithMessage("fail to find", c)
	} else {
		response.OkWithData(gin.H{"redocumentSigners": redocumentSigners}, c)
	}
}

// GetDocumentSignersList get signer list
// @Tags DocumentSigners
// @Summary get signer list
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dmsReq.DocumentSignersSearch true "get signer list"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentSigners/getDocumentSignersList [get]
func (documentSignersApi *DocumentSignersApi) GetDocumentSignersList(c *gin.Context) {
	var pageInfo dmsReq.DocumentSignersSearch
	var err error

	err = c.ShouldBindQuery(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("provide valid data", zap.Error(err))
		response.FailWithMessage("provide valid data", c)
		return
	}

	if list, total, err := documentSignersService.GetDocumentSignersInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("fail to get list", zap.Error(err))
		response.FailWithMessage("fail to get list", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "success", c)
	}
}
