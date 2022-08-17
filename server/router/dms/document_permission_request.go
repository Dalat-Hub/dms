package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DocumentPermissionRequestRouter struct {
}

func (s *DocumentPermissionRequestRouter) InitDocumentPermissionRequestRouter(Router *gin.RouterGroup) {
	documentPermissionRequestRouter := Router.Group("documentPermissionRequest").Use(middleware.OperationRecord())
	documentPermissionRequestRouterWithoutRecord := Router.Group("documentPermissionRequest")
	var documentPermissionRequestApi = v1.ApiGroupApp.DmsApiGroup.DocumentPermissionRequestApi
	{
		documentPermissionRequestRouter.POST("createDocumentPermissionRequest", documentPermissionRequestApi.CreateDocumentPermissionRequest)             // 新建DocumentPermissionRequest
		documentPermissionRequestRouter.DELETE("deleteDocumentPermissionRequest", documentPermissionRequestApi.DeleteDocumentPermissionRequest)           // 删除DocumentPermissionRequest
		documentPermissionRequestRouter.DELETE("deleteDocumentPermissionRequestByIds", documentPermissionRequestApi.DeleteDocumentPermissionRequestByIds) // 批量删除DocumentPermissionRequest
		documentPermissionRequestRouter.PUT("updateDocumentPermissionRequest", documentPermissionRequestApi.UpdateDocumentPermissionRequest)              // 更新DocumentPermissionRequest
	}
	{
		documentPermissionRequestRouterWithoutRecord.GET("findDocumentPermissionRequest", documentPermissionRequestApi.FindDocumentPermissionRequest)       // 根据ID获取DocumentPermissionRequest
		documentPermissionRequestRouterWithoutRecord.GET("getDocumentPermissionRequestList", documentPermissionRequestApi.GetDocumentPermissionRequestList) // 获取DocumentPermissionRequest列表
	}
}
