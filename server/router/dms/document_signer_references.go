package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DocumentSignerReferencesRouter struct {
}

// InitDocumentSignerReferencesRouter Initialize DocumentSignerReferences routing information
func (s *DocumentSignerReferencesRouter) InitDocumentSignerReferencesRouter(Router *gin.RouterGroup) {
	documentSignerReferencesRouter := Router.Group("documentSignerReferences").Use(middleware.OperationRecord())
	documentSignerReferencesRouterWithoutRecord := Router.Group("documentSignerReferences")
	var documentSignerReferencesApi = v1.ApiGroupApp.DmsApiGroup.DocumentSignerReferencesApi
	{
		documentSignerReferencesRouter.POST("createDocumentSignerReferences", documentSignerReferencesApi.CreateDocumentSignerReferences)
		documentSignerReferencesRouter.DELETE("deleteDocumentSignerReferences", documentSignerReferencesApi.DeleteDocumentSignerReferences)
		documentSignerReferencesRouter.DELETE("deleteDocumentSignerReferencesByIds", documentSignerReferencesApi.DeleteDocumentSignerReferencesByIds)
		documentSignerReferencesRouter.PUT("updateDocumentSignerReferences", documentSignerReferencesApi.UpdateDocumentSignerReferences)
	}
	{
		documentSignerReferencesRouterWithoutRecord.GET("findDocumentSignerReferences", documentSignerReferencesApi.FindDocumentSignerReferences)
		documentSignerReferencesRouterWithoutRecord.GET("getDocumentSignerReferencesList", documentSignerReferencesApi.GetDocumentSignerReferencesList)
	}
}
