package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DocumentFieldReferencesRouter struct {
}

// InitDocumentFieldReferencesRouter Initialize DocumentFieldReferences routing information
func (s *DocumentFieldReferencesRouter) InitDocumentFieldReferencesRouter(Router *gin.RouterGroup) {
	documentFieldReferencesRouter := Router.Group("documentFieldReferences").Use(middleware.OperationRecord())
	documentFieldReferencesRouterWithoutRecord := Router.Group("documentFieldReferences")
	var documentFieldReferencesApi = v1.ApiGroupApp.DmsApiGroup.DocumentFieldReferencesApi
	{
		documentFieldReferencesRouter.POST("createDocumentFieldReferences", documentFieldReferencesApi.CreateDocumentFieldReferences)
		documentFieldReferencesRouter.DELETE("deleteDocumentFieldReferences", documentFieldReferencesApi.DeleteDocumentFieldReferences)
		documentFieldReferencesRouter.DELETE("deleteDocumentFieldReferencesByIds", documentFieldReferencesApi.DeleteDocumentFieldReferencesByIds)
		documentFieldReferencesRouter.PUT("updateDocumentFieldReferences", documentFieldReferencesApi.UpdateDocumentFieldReferences)
	}
	{
		documentFieldReferencesRouterWithoutRecord.GET("findDocumentFieldReferences", documentFieldReferencesApi.FindDocumentFieldReferences)
		documentFieldReferencesRouterWithoutRecord.GET("getDocumentFieldReferencesList", documentFieldReferencesApi.GetDocumentFieldReferencesList)
	}
}
