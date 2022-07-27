package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DocumentsRouter struct {
}

// InitDocumentsRouter Initialize DocumentSignerReferences routing information
func (s *DocumentsRouter) InitDocumentsRouter(Router *gin.RouterGroup) {
	documentsRouter := Router.Group("documents").Use(middleware.OperationRecord())
	documentsRouterWithoutRecord := Router.Group("documents")
	var documentsApi = v1.ApiGroupApp.DmsApiGroup.DocumentsApi
	{
		documentsRouter.POST("createDocuments", documentsApi.CreateDocuments)
		documentsRouter.POST("createDraftDocuments", documentsApi.CreateDraftDocument)
		documentsRouter.POST("createFullDocuments", documentsApi.CreateFullDocument)
		documentsRouter.DELETE("deleteDocuments", documentsApi.DeleteDocuments)
		documentsRouter.DELETE("deleteDocumentsByIds", documentsApi.DeleteDocumentsByIds)
		documentsRouter.PUT("updateDocuments", documentsApi.UpdateDocuments)
		documentsRouter.PUT("updateBasicDocuments", documentsApi.UpdateBasicDocumentInformation)
		documentsRouter.POST("makeDuplication", documentsApi.MakeDuplication)
	}
	{
		documentsRouterWithoutRecord.GET("findDocuments", documentsApi.FindDocuments)
		documentsRouterWithoutRecord.GET("getDocumentsList", documentsApi.GetDocumentsList)
		documentsRouterWithoutRecord.GET("getDocumentFiles", documentsApi.GetFileList)
	}
}
