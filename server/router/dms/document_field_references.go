package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DocumentFieldReferencesRouter struct {
}

// InitDocumentFieldReferencesRouter 初始化 DocumentFieldReferences 路由信息
func (s *DocumentFieldReferencesRouter) InitDocumentFieldReferencesRouter(Router *gin.RouterGroup) {
	documentFieldReferencesRouter := Router.Group("documentFieldReferences").Use(middleware.OperationRecord())
	documentFieldReferencesRouterWithoutRecord := Router.Group("documentFieldReferences")
	var documentFieldReferencesApi = v1.ApiGroupApp.DmsApiGroup.DocumentFieldReferencesApi
	{
		documentFieldReferencesRouter.POST("createDocumentFieldReferences", documentFieldReferencesApi.CreateDocumentFieldReferences)             // 新建DocumentFieldReferences
		documentFieldReferencesRouter.DELETE("deleteDocumentFieldReferences", documentFieldReferencesApi.DeleteDocumentFieldReferences)           // 删除DocumentFieldReferences
		documentFieldReferencesRouter.DELETE("deleteDocumentFieldReferencesByIds", documentFieldReferencesApi.DeleteDocumentFieldReferencesByIds) // 批量删除DocumentFieldReferences
		documentFieldReferencesRouter.PUT("updateDocumentFieldReferences", documentFieldReferencesApi.UpdateDocumentFieldReferences)              // 更新DocumentFieldReferences
	}
	{
		documentFieldReferencesRouterWithoutRecord.GET("findDocumentFieldReferences", documentFieldReferencesApi.FindDocumentFieldReferences)       // 根据ID获取DocumentFieldReferences
		documentFieldReferencesRouterWithoutRecord.GET("getDocumentFieldReferencesList", documentFieldReferencesApi.GetDocumentFieldReferencesList) // 获取DocumentFieldReferences列表
	}
}
