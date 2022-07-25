package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DocumentFilesRouter struct {
}

// InitDocumentFilesRouter Initialize DocumentFiles routing information
func (s *DocumentFilesRouter) InitDocumentFilesRouter(Router *gin.RouterGroup) {
	documentFilesRouter := Router.Group("documentFiles").Use(middleware.OperationRecord())
	documentFilesRouterWithoutRecord := Router.Group("documentFiles")
	var documentFilesApi = v1.ApiGroupApp.DmsApiGroup.DocumentFilesApi
	{
		documentFilesRouter.POST("createDocumentFiles", documentFilesApi.CreateDocumentFiles)
		documentFilesRouter.DELETE("deleteDocumentFiles", documentFilesApi.DeleteDocumentFiles)
		documentFilesRouter.DELETE("deleteDocumentFilesByIds", documentFilesApi.DeleteDocumentFilesByIds)
		documentFilesRouter.PUT("updateDocumentFiles", documentFilesApi.UpdateDocumentFiles)
	}
	{
		documentFilesRouterWithoutRecord.GET("findDocumentFiles", documentFilesApi.FindDocumentFiles)
		documentFilesRouterWithoutRecord.GET("getDocumentFilesList", documentFilesApi.GetDocumentFilesList)
	}
}
