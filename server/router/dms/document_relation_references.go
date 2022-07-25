package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DocumentRelationReferencesRouter struct {
}

// InitDocumentRelationReferencesRouter Initialize DocumentRelationReferences routing information
func (s *DocumentRelationReferencesRouter) InitDocumentRelationReferencesRouter(Router *gin.RouterGroup) {
	documentRelationReferencesRouter := Router.Group("documentRelationReferences").Use(middleware.OperationRecord())
	documentRelationReferencesRouterWithoutRecord := Router.Group("documentRelationReferences")
	var documentRelationReferencesApi = v1.ApiGroupApp.DmsApiGroup.DocumentRelationReferencesApi
	{
		documentRelationReferencesRouter.POST("createDocumentRelationReferences", documentRelationReferencesApi.CreateDocumentRelationReferences)
		documentRelationReferencesRouter.DELETE("deleteDocumentRelationReferences", documentRelationReferencesApi.DeleteDocumentRelationReferences)
		documentRelationReferencesRouter.DELETE("deleteDocumentRelationReferencesByIds", documentRelationReferencesApi.DeleteDocumentRelationReferencesByIds)
		documentRelationReferencesRouter.PUT("updateDocumentRelationReferences", documentRelationReferencesApi.UpdateDocumentRelationReferences)
	}
	{
		documentRelationReferencesRouterWithoutRecord.GET("findDocumentRelationReferences", documentRelationReferencesApi.FindDocumentRelationReferences)
		documentRelationReferencesRouterWithoutRecord.GET("getDocumentRelationReferencesList", documentRelationReferencesApi.GetDocumentRelationReferencesList)
	}
}
