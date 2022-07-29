package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DocumentPriorityLevelsRouter struct {
}

// InitDocumentPriorityLevelsRouter Initialize DocumentPriorityLevels routing information
func (s *DocumentPriorityLevelsRouter) InitDocumentPriorityLevelsRouter(Router *gin.RouterGroup) {
	documentPriorityLevelsRouter := Router.Group("documentPriorityLevels").Use(middleware.OperationRecord())
	documentPriorityLevelsRouterWithoutRecord := Router.Group("documentPriorityLevels")
	var documentPriorityLevelsApi = v1.ApiGroupApp.DmsApiGroup.DocumentPriorityLevelsApi
	{
		documentPriorityLevelsRouter.POST("createDocumentPriorityLevels", documentPriorityLevelsApi.CreateDocumentPriorityLevels)
		documentPriorityLevelsRouter.DELETE("deleteDocumentPriorityLevels", documentPriorityLevelsApi.DeleteDocumentPriorityLevels)
		documentPriorityLevelsRouter.DELETE("deleteDocumentPriorityLevelsByIds", documentPriorityLevelsApi.DeleteDocumentPriorityLevelsByIds)
		documentPriorityLevelsRouter.PUT("updateDocumentPriorityLevels", documentPriorityLevelsApi.UpdateDocumentPriorityLevels)
	}
	{
		documentPriorityLevelsRouterWithoutRecord.GET("findDocumentPriorityLevels", documentPriorityLevelsApi.FindDocumentPriorityLevels)
		documentPriorityLevelsRouterWithoutRecord.GET("getDocumentPriorityLevelsList", documentPriorityLevelsApi.GetDocumentPriorityLevelsList)
	}
}
