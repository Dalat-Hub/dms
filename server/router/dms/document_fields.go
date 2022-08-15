package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DocumentFieldsRouter struct {
}

func (s *DocumentFieldsRouter) InitDocumentFieldsPublicRouter(Router *gin.RouterGroup) {
	documentFieldsRouterWithoutRecord := Router.Group("fields")
	var documentFieldsApi = v1.ApiGroupApp.DmsApiGroup.DocumentFieldsApi
	{
		documentFieldsRouterWithoutRecord.GET("", documentFieldsApi.GetDocumentFieldsList)
	}
}

// InitDocumentFieldsRouter Initialize DocumentFields routing information
func (s *DocumentFieldsRouter) InitDocumentFieldsRouter(Router *gin.RouterGroup) {
	documentFieldsRouter := Router.Group("documentFields").Use(middleware.OperationRecord())
	documentFieldsRouterWithoutRecord := Router.Group("documentFields")
	var documentFieldsApi = v1.ApiGroupApp.DmsApiGroup.DocumentFieldsApi
	{
		documentFieldsRouter.POST("createDocumentFields", documentFieldsApi.CreateDocumentFields)
		documentFieldsRouter.DELETE("deleteDocumentFields", documentFieldsApi.DeleteDocumentFields)
		documentFieldsRouter.DELETE("deleteDocumentFieldsByIds", documentFieldsApi.DeleteDocumentFieldsByIds)
		documentFieldsRouter.PUT("updateDocumentFields", documentFieldsApi.UpdateDocumentFields)
	}
	{
		documentFieldsRouterWithoutRecord.GET("findDocumentFields", documentFieldsApi.FindDocumentFields)
		documentFieldsRouterWithoutRecord.GET("getDocumentFieldsList", documentFieldsApi.GetDocumentFieldsList)
	}
}
