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

type DocumentRulesApi struct {
}

var documentRulesService = service.ServiceGroupApp.DmsServiceGroup.DocumentRulesService

// CreateDocumentRules create new rule
// @Tags DocumentRules
// @Summary create new rule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentRules true "create new rule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentRules/createDocumentRules [post]
func (documentRulesApi *DocumentRulesApi) CreateDocumentRules(c *gin.Context) {
	var documentRules dms.DocumentRules
	_ = c.ShouldBindJSON(&documentRules)
	if err := documentRulesService.CreateDocumentRules(documentRules); err != nil {
		global.GVA_LOG.Error("fail to create document rule", zap.Error(err))
		response.FailWithMessage("fail to create document rule", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// DeleteDocumentRules delete a rule
// @Tags DocumentRules
// @Summary delete a rule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentRules true "delete a rule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentRules/deleteDocumentRules [delete]
func (documentRulesApi *DocumentRulesApi) DeleteDocumentRules(c *gin.Context) {
	var documentRules dms.DocumentRules
	_ = c.ShouldBindJSON(&documentRules)
	if err := documentRulesService.DeleteDocumentRules(documentRules); err != nil {
		global.GVA_LOG.Error("fail to delete rule", zap.Error(err))
		response.FailWithMessage("fail to delete rule", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// DeleteDocumentRulesByIds bulk delete rules
// @Tags DocumentRules
// @Summary bulk delete rules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "bulk delete rules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentRules/deleteDocumentRulesByIds [delete]
func (documentRulesApi *DocumentRulesApi) DeleteDocumentRulesByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := documentRulesService.DeleteDocumentRulesByIds(IDS); err != nil {
		global.GVA_LOG.Error("fail to bulk delete rules", zap.Error(err))
		response.FailWithMessage("fail to bulk delete rules", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// UpdateDocumentRules update rule
// @Tags DocumentRules
// @Summary update rule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dms.DocumentRules true "update rule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentRules/updateDocumentRules [put]
func (documentRulesApi *DocumentRulesApi) UpdateDocumentRules(c *gin.Context) {
	var documentRules dms.DocumentRules
	_ = c.ShouldBindJSON(&documentRules)
	if err := documentRulesService.UpdateDocumentRules(documentRules); err != nil {
		global.GVA_LOG.Error("fail to update rule", zap.Error(err))
		response.FailWithMessage("fail to update rule", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// FindDocumentRules get rule by ID
// @Tags DocumentRules
// @Summary get rule by ID
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dms.DocumentRules true "get rule by ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentRules/findDocumentRules [get]
func (documentRulesApi *DocumentRulesApi) FindDocumentRules(c *gin.Context) {
	var documentRules dms.DocumentRules
	_ = c.ShouldBindQuery(&documentRules)
	if redocumentRules, err := documentRulesService.GetDocumentRules(documentRules.ID); err != nil {
		global.GVA_LOG.Error("fail to get rule by ID", zap.Error(err))
		response.FailWithMessage("fail to get rule by ID", c)
	} else {
		response.OkWithData(gin.H{"redocumentRules": redocumentRules}, c)
	}
}

// GetDocumentRulesList get list of rules
// @Tags DocumentRules
// @Summary get list of rules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dmsReq.DocumentRulesSearch true "get list of rules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentRules/getDocumentRulesList [get]
func (documentRulesApi *DocumentRulesApi) GetDocumentRulesList(c *gin.Context) {
	var pageInfo dmsReq.DocumentRulesSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := documentRulesService.GetDocumentRulesInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("fail to get list of rules", zap.Error(err))
		response.FailWithMessage("fail to get list of rules", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "success", c)
	}
}
