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

type DocumentPermissionRequestApi struct {
}

var documentPermissionRequestService = service.ServiceGroupApp.DmsServiceGroup.DocumentPermissionRequestService

// CreateDocumentPermissionRequest create new permission request
// @Tags DocumentPermissionRequest
// @Summary create new permission request
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentPermissionRequest true "create new permission request"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentPermissionRequest/createDocumentPermissionRequest [post]
func (documentPermissionRequestApi *DocumentPermissionRequestApi) CreateDocumentPermissionRequest(c *gin.Context) {
	var documentPermissionRequest dms.DocumentPermissionRequest
	var err error

	err = c.ShouldBindJSON(&documentPermissionRequest)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err = documentPermissionRequestService.CreateDocumentPermissionRequest(&documentPermissionRequest); err != nil {
		global.GVA_LOG.Error("fail to create new permission request", zap.Error(err))
		response.FailWithMessage("fail to create new permission request", c)
	} else {
		response.OkWithData(gin.H{"request": documentPermissionRequest}, c)
	}
}

// ApproveDocumentPermissionRequest approve permission request for document
// @Tags DocumentPermissionRequest
// @Summary approve permission request for document
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentPermissionRequest true "approve permission request for document"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentPermissionRequest/approvePermissionRequest [post]
func (documentPermissionRequestApi *DocumentPermissionRequestApi) ApproveDocumentPermissionRequest(c *gin.Context) {
	var documentPermissionRequest dms.DocumentPermissionRequest
	var err error

	err = c.ShouldBindJSON(&documentPermissionRequest)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	userInfo := utils.GetUserInfo(c)
	if userInfo == nil {
		response.FailWithMessage("please login to process", c)
		return
	}

	if err = documentPermissionRequestService.ApprovePermissionRequest(&documentPermissionRequest, userInfo); err != nil {
		global.GVA_LOG.Error("fail to approve permission request", zap.Error(err))
		response.FailWithMessage("fail to approve permission request", c)
	} else {
		response.OkWithData(gin.H{"request": documentPermissionRequest}, c)
	}
}

func (documentPermissionRequestApi *DocumentPermissionRequestApi) CreateDocumentPermissionRequestPublic(c *gin.Context) {
	var documentPermissionRequest dms.DocumentPermissionRequest
	var err error

	err = c.ShouldBindJSON(&documentPermissionRequest)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if documentPermissionRequest.DocumentId <= 0 {
		response.FailWithMessage("văn bản yêu cầu không hợp lệ", c)
		return
	}

	userInfo := utils.GetUserInfo(c)
	if userInfo == nil {
		response.FailWithMessage("bạn phải đăng nhập để gửi yêu cầu cấp quyền", c)
		return
	}

	documentPermissionRequest.RequestPermission = dms.PERMISSION_VIEW
	documentPermissionRequest.RequestUserId = userInfo.ID

	if err = documentPermissionRequestService.CreateDocumentPermissionRequest(&documentPermissionRequest); err != nil {
		global.GVA_LOG.Error("fail to create new permission request", zap.Error(err))
		response.FailWithMessage("fail to create new permission request", c)
	} else {
		response.OkWithData(gin.H{"request": documentPermissionRequest}, c)
	}
}

// DeleteDocumentPermissionRequest delete permission request
// @Tags DocumentPermissionRequest
// @Summary delete permission request
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentPermissionRequest true "delete permission request"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentPermissionRequest/deleteDocumentPermissionRequest [delete]
func (documentPermissionRequestApi *DocumentPermissionRequestApi) DeleteDocumentPermissionRequest(c *gin.Context) {
	var documentPermissionRequest dms.DocumentPermissionRequest
	var err error

	err = c.ShouldBindJSON(&documentPermissionRequest)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err = documentPermissionRequestService.DeleteDocumentPermissionRequest(documentPermissionRequest); err != nil {
		global.GVA_LOG.Error("fail to delete permission request", zap.Error(err))
		response.FailWithMessage("fail to delete permission request", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// DeleteDocumentPermissionRequestByIds bulk delete permission requests
// @Tags DocumentPermissionRequest
// @Summary bulk delete permission requests
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "bulk delete permission request"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentPermissionRequest/deleteDocumentPermissionRequestByIds [delete]
func (documentPermissionRequestApi *DocumentPermissionRequestApi) DeleteDocumentPermissionRequestByIds(c *gin.Context) {
	var IDS request.IdsReq
	var err error

	err = c.ShouldBindJSON(&IDS)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err = documentPermissionRequestService.DeleteDocumentPermissionRequestByIds(IDS); err != nil {
		global.GVA_LOG.Error("fail to bulk delete permission requests", zap.Error(err))
		response.FailWithMessage("fail to bulk delete permission requests", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// UpdateDocumentPermissionRequest update permission request
// @Tags DocumentPermissionRequest
// @Summary update permission request
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentPermissionRequest true "update permission request"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentPermissionRequest/updateDocumentPermissionRequest [put]
func (documentPermissionRequestApi *DocumentPermissionRequestApi) UpdateDocumentPermissionRequest(c *gin.Context) {
	var documentPermissionRequest dms.DocumentPermissionRequest
	var err error

	err = c.ShouldBindJSON(&documentPermissionRequest)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err = documentPermissionRequestService.UpdateDocumentPermissionRequest(documentPermissionRequest); err != nil {
		global.GVA_LOG.Error("fail to update permission request", zap.Error(err))
		response.FailWithMessage("fail to update permission request", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// FindDocumentPermissionRequest get request by id
// @Tags DocumentPermissionRequest
// @Summary get request by id
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dms.DocumentPermissionRequest true "get request by id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentPermissionRequest/findDocumentPermissionRequest [get]
func (documentPermissionRequestApi *DocumentPermissionRequestApi) FindDocumentPermissionRequest(c *gin.Context) {
	var documentPermissionRequest dms.DocumentPermissionRequest
	var err error

	err = c.ShouldBindQuery(&documentPermissionRequest)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if redocumentPermissionRequest, err := documentPermissionRequestService.GetDocumentPermissionRequest(documentPermissionRequest.ID); err != nil {
		global.GVA_LOG.Error("fail to get request by id", zap.Error(err))
		response.FailWithMessage("fail to get request by id", c)
	} else {
		response.OkWithData(gin.H{"redocumentPermissionRequest": redocumentPermissionRequest}, c)
	}
}

// GetDocumentPermissionRequestList get list of requests
// @Tags DocumentPermissionRequest
// @Summary get list of requests
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dmsReq.DocumentPermissionRequestSearch true "get list of requests"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentPermissionRequest/getDocumentPermissionRequestList [get]
func (documentPermissionRequestApi *DocumentPermissionRequestApi) GetDocumentPermissionRequestList(c *gin.Context) {
	var pageInfo dmsReq.DocumentPermissionRequestSearch
	var err error

	err = c.ShouldBindQuery(&pageInfo)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if list, total, err := documentPermissionRequestService.GetDocumentPermissionRequestInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("fail to get list of requests", zap.Error(err))
		response.FailWithMessage("fail to get list of requests", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "success", c)
	}
}
