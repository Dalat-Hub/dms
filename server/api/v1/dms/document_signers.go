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

// CreateDocumentSigners 创建DocumentSigners
// @Tags DocumentSigners
// @Summary 创建DocumentSigners
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentSigners true "创建DocumentSigners"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /documentSigners/createDocumentSigners [post]
func (documentSignersApi *DocumentSignersApi) CreateDocumentSigners(c *gin.Context) {
	var documentSigners dms.DocumentSigners
	_ = c.ShouldBindJSON(&documentSigners)
	if err := documentSignersService.CreateDocumentSigners(documentSigners); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDocumentSigners 删除DocumentSigners
// @Tags DocumentSigners
// @Summary 删除DocumentSigners
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentSigners true "删除DocumentSigners"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /documentSigners/deleteDocumentSigners [delete]
func (documentSignersApi *DocumentSignersApi) DeleteDocumentSigners(c *gin.Context) {
	var documentSigners dms.DocumentSigners
	_ = c.ShouldBindJSON(&documentSigners)
	if err := documentSignersService.DeleteDocumentSigners(documentSigners); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDocumentSignersByIds 批量删除DocumentSigners
// @Tags DocumentSigners
// @Summary 批量删除DocumentSigners
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DocumentSigners"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /documentSigners/deleteDocumentSignersByIds [delete]
func (documentSignersApi *DocumentSignersApi) DeleteDocumentSignersByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := documentSignersService.DeleteDocumentSignersByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDocumentSigners 更新DocumentSigners
// @Tags DocumentSigners
// @Summary 更新DocumentSigners
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentSigners true "更新DocumentSigners"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /documentSigners/updateDocumentSigners [put]
func (documentSignersApi *DocumentSignersApi) UpdateDocumentSigners(c *gin.Context) {
	var documentSigners dms.DocumentSigners
	_ = c.ShouldBindJSON(&documentSigners)
	if err := documentSignersService.UpdateDocumentSigners(documentSigners); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDocumentSigners 用id查询DocumentSigners
// @Tags DocumentSigners
// @Summary 用id查询DocumentSigners
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dms.DocumentSigners true "用id查询DocumentSigners"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /documentSigners/findDocumentSigners [get]
func (documentSignersApi *DocumentSignersApi) FindDocumentSigners(c *gin.Context) {
	var documentSigners dms.DocumentSigners
	_ = c.ShouldBindQuery(&documentSigners)
	if redocumentSigners, err := documentSignersService.GetDocumentSigners(documentSigners.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redocumentSigners": redocumentSigners}, c)
	}
}

// GetDocumentSignersList 分页获取DocumentSigners列表
// @Tags DocumentSigners
// @Summary 分页获取DocumentSigners列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dmsReq.DocumentSignersSearch true "分页获取DocumentSigners列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /documentSigners/getDocumentSignersList [get]
func (documentSignersApi *DocumentSignersApi) GetDocumentSignersList(c *gin.Context) {
	var pageInfo dmsReq.DocumentSignersSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := documentSignersService.GetDocumentSignersInfoList(pageInfo); err != nil {
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
