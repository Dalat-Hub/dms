package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DocumentAgenciesRouter struct {
}

func (s *DocumentAgenciesRouter) InitDocumentAgenciesPublicRouter(Router *gin.RouterGroup) {
	documentAgenciesRouterWithoutRecord := Router.Group("agencies")
	var documentAgenciesApi = v1.ApiGroupApp.DmsApiGroup.DocumentAgenciesApi
	{
		documentAgenciesRouterWithoutRecord.GET("", documentAgenciesApi.GetDocumentAgenciesListPublic)
		documentAgenciesRouterWithoutRecord.GET("tree", documentAgenciesApi.GetAgencyTree)
		documentAgenciesRouterWithoutRecord.GET("tree/fields", documentAgenciesApi.GetAgencyTreeForField)
	}
}

// InitDocumentAgenciesRouter Initialize DocumentAgencies routing information
func (s *DocumentAgenciesRouter) InitDocumentAgenciesRouter(Router *gin.RouterGroup) {
	documentAgenciesRouter := Router.Group("documentAgencies").Use(middleware.OperationRecord())
	documentAgenciesRouterWithoutRecord := Router.Group("documentAgencies")
	var documentAgenciesApi = v1.ApiGroupApp.DmsApiGroup.DocumentAgenciesApi
	{
		documentAgenciesRouter.POST("createDocumentAgencies", documentAgenciesApi.CreateDocumentAgencies)
		documentAgenciesRouter.DELETE("deleteDocumentAgencies", documentAgenciesApi.DeleteDocumentAgencies)
		documentAgenciesRouter.DELETE("deleteDocumentAgenciesByIds", documentAgenciesApi.DeleteDocumentAgenciesByIds)
		documentAgenciesRouter.PUT("updateDocumentAgencies", documentAgenciesApi.UpdateDocumentAgencies)
	}
	{
		documentAgenciesRouterWithoutRecord.GET("findDocumentAgencies", documentAgenciesApi.FindDocumentAgencies)
		documentAgenciesRouterWithoutRecord.GET("getDocumentAgenciesList", documentAgenciesApi.GetDocumentAgenciesList)
	}
}
