package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DocumentPermissionRequestRouter struct {
}

func (s *DocumentPermissionRequestRouter) InitDocumentPermissionRequestPublicRouter(Router *gin.RouterGroup) {
	documentPermissionRequestRouterWithoutRecord := Router.Group("permission-requests")
	var documentPermissionRequestApi = v1.ApiGroupApp.DmsApiGroup.DocumentPermissionRequestApi
	{
		documentPermissionRequestRouterWithoutRecord.POST("", documentPermissionRequestApi.CreateDocumentPermissionRequestPublic)
	}
}

func (s *DocumentPermissionRequestRouter) InitDocumentPermissionRequestRouter(Router *gin.RouterGroup) {
	documentPermissionRequestRouter := Router.Group("documentPermissionRequest").Use(middleware.OperationRecord())
	documentPermissionRequestRouterWithoutRecord := Router.Group("documentPermissionRequest")
	var documentPermissionRequestApi = v1.ApiGroupApp.DmsApiGroup.DocumentPermissionRequestApi
	{
		documentPermissionRequestRouter.POST("createDocumentPermissionRequest", documentPermissionRequestApi.CreateDocumentPermissionRequest)
		documentPermissionRequestRouter.DELETE("deleteDocumentPermissionRequest", documentPermissionRequestApi.DeleteDocumentPermissionRequest)
		documentPermissionRequestRouter.DELETE("deleteDocumentPermissionRequestByIds", documentPermissionRequestApi.DeleteDocumentPermissionRequestByIds)
		documentPermissionRequestRouter.PUT("updateDocumentPermissionRequest", documentPermissionRequestApi.UpdateDocumentPermissionRequest)
	}
	{
		documentPermissionRequestRouterWithoutRecord.GET("findDocumentPermissionRequest", documentPermissionRequestApi.FindDocumentPermissionRequest)
		documentPermissionRequestRouterWithoutRecord.GET("getDocumentPermissionRequestList", documentPermissionRequestApi.GetDocumentPermissionRequestList)
	}
}
