package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DocumentRulesRouter struct {
}

func (s *DocumentRulesRouter) InitDocumentRulesRouter(Router *gin.RouterGroup) {
	documentRulesRouter := Router.Group("documentRules").Use(middleware.OperationRecord())
	documentRulesRouterWithoutRecord := Router.Group("documentRules")
	var documentRulesApi = v1.ApiGroupApp.DmsApiGroup.DocumentRulesApi
	{
		documentRulesRouter.POST("createDocumentRules", documentRulesApi.CreateDocumentRules)
		documentRulesRouter.DELETE("deleteDocumentRules", documentRulesApi.DeleteDocumentRules)
		documentRulesRouter.DELETE("deleteDocumentRulesByIds", documentRulesApi.DeleteDocumentRulesByIds)
		documentRulesRouter.PUT("updateDocumentRules", documentRulesApi.UpdateDocumentRules)
	}
	{
		documentRulesRouterWithoutRecord.GET("findDocumentRules", documentRulesApi.FindDocumentRules)
		documentRulesRouterWithoutRecord.GET("getDocumentRulesList", documentRulesApi.GetDocumentRulesList)
	}
}
