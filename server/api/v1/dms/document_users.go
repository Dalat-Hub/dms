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

type DocumentUsersApi struct {
}

var documentUsersService = service.ServiceGroupApp.DmsServiceGroup.DocumentUsersService

// CreateDocumentUsers create new pivot
// @Tags DocumentUsers
// @Summary create new relation between user and document
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentUsers true "create new relation between user and document"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentUsers/createDocumentUsers [post]
func (documentUsersApi *DocumentUsersApi) CreateDocumentUsers(c *gin.Context) {
	var documentUsers dms.DocumentUsers
	var err error

	err = c.ShouldBindJSON(&documentUsers)
	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err = documentUsersService.CreateDocumentUsers(documentUsers); err != nil {
		global.GVA_LOG.Error("error occurs", zap.Error(err))
		response.FailWithMessage("error occurs", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// DeleteDocumentUsers delete relation
// @Tags DocumentUsers
// @Summary delete relation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentUsers true "delete relation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentUsers/deleteDocumentUsers [delete]
func (documentUsersApi *DocumentUsersApi) DeleteDocumentUsers(c *gin.Context) {
	var documentUsers dms.DocumentUsers
	var err error

	err = c.ShouldBindJSON(&documentUsers)
	if err != nil {
		global.GVA_LOG.Error("provide valid data", zap.Error(err))
		response.FailWithMessage("provide valid data", c)
		return
	}

	if err = documentUsersService.DeleteDocumentUsers(documentUsers); err != nil {
		global.GVA_LOG.Error("error occurs", zap.Error(err))
		response.FailWithMessage("error occurs", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// DeleteDocumentUsersByIds bulk delete by IDs
// @Tags DocumentUsers
// @Summary bulk delete by IDs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "bulk delete by IDs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentUsers/deleteDocumentUsersByIds [delete]
func (documentUsersApi *DocumentUsersApi) DeleteDocumentUsersByIds(c *gin.Context) {
	var IDS request.IdsReq
	var err error

	err = c.ShouldBindJSON(&IDS)
	if err != nil {
		global.GVA_LOG.Error("provide valid data", zap.Error(err))
		response.FailWithMessage("provide valid data", c)
		return
	}

	if err = documentUsersService.DeleteDocumentUsersByIds(IDS); err != nil {
		global.GVA_LOG.Error("error occurs", zap.Error(err))
		response.FailWithMessage("error occurs", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// UpdateDocumentUsers update relations
// @Tags DocumentUsers
// @Summary update relations
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentUsers true "udpate relations"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentUsers/updateDocumentUsers [put]
func (documentUsersApi *DocumentUsersApi) UpdateDocumentUsers(c *gin.Context) {
	var documentUsers dms.DocumentUsers
	var err error

	err = c.ShouldBindJSON(&documentUsers)
	if err != nil {
		global.GVA_LOG.Error("provide valid data", zap.Error(err))
		response.FailWithMessage("provide valid data", c)
		return
	}

	if err = documentUsersService.UpdateDocumentUsers(documentUsers); err != nil {
		global.GVA_LOG.Error("error occurs", zap.Error(err))
		response.FailWithMessage("error occurs", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// FindDocumentUsers get relation by ID
// @Tags DocumentUsers
// @Summary get relation by ID
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dms.DocumentUsers true "get relation by ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentUsers/findDocumentUsers [get]
func (documentUsersApi *DocumentUsersApi) FindDocumentUsers(c *gin.Context) {
	var documentUsers dms.DocumentUsers
	var err error

	err = c.ShouldBindQuery(&documentUsers)
	if err != nil {
		global.GVA_LOG.Error("provide valid data", zap.Error(err))
		response.FailWithMessage("provide valid data", c)
		return
	}

	if redocumentUsers, err := documentUsersService.GetDocumentUsers(documentUsers.ID); err != nil {
		global.GVA_LOG.Error("error occurs", zap.Error(err))
		response.FailWithMessage("error occurs", c)
	} else {
		response.OkWithData(gin.H{"redocumentUsers": redocumentUsers}, c)
	}
}

// GetDocumentUsersList get relation list
// @Tags DocumentUsers
// @Summary get relation list
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dmsReq.DocumentUsersSearch true "get relation list"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"succeess"}"
// @Router /documentUsers/getDocumentUsersList [get]
func (documentUsersApi *DocumentUsersApi) GetDocumentUsersList(c *gin.Context) {
	var pageInfo dmsReq.DocumentUsersSearch
	var err error

	err = c.ShouldBindQuery(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("provide valid data", zap.Error(err))
		response.FailWithMessage("provide valid data", c)
		return
	}

	if list, total, err := documentUsersService.GetDocumentUsersInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("error occurs", zap.Error(err))
		response.FailWithMessage("error occurs", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "success", c)
	}
}

// DistributeTasks redistribute tasks
// @Tags DocumentUsers
// @Summary redistribute tasks
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentUsersDistribution true "redistribute tasks"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentUsers/distributeTasks [post]
func (documentUsersApi *DocumentUsersApi) DistributeTasks(c *gin.Context) {
	var task dmsReq.DocumentUsersDistribution
	var err error

	err = c.ShouldBindJSON(&task)
	if err != nil {
		global.GVA_LOG.Error("provide valid data", zap.Error(err))
		response.FailWithMessage("provide valid data", c)
		return
	}

	if err = documentUsersService.DistributeTasks(task); err != nil {
		global.GVA_LOG.Error("error occurs", zap.Error(err))
		response.FailWithMessage("error occurs", c)
	} else {
		response.OkWithMessage("success", c)
	}
}
