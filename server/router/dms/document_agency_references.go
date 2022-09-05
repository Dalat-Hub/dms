package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DocumentAgencyReferencesRouter struct {
}

// InitDocumentAgencyReferencesRouter 初始化 DocumentAgencyReferences 路由信息
func (s *DocumentAgencyReferencesRouter) InitDocumentAgencyReferencesRouter(Router *gin.RouterGroup) {
	documentAgencyReferencesRouter := Router.Group("documentAgencyReferences").Use(middleware.OperationRecord())
	documentAgencyReferencesRouterWithoutRecord := Router.Group("documentAgencyReferences")
	var documentAgencyReferencesApi = v1.ApiGroupApp.DmsApiGroup.DocumentAgencyReferencesApi
	{
		documentAgencyReferencesRouter.POST("createDocumentAgencyReferences", documentAgencyReferencesApi.CreateDocumentAgencyReferences)             // 新建DocumentAgencyReferences
		documentAgencyReferencesRouter.DELETE("deleteDocumentAgencyReferences", documentAgencyReferencesApi.DeleteDocumentAgencyReferences)           // 删除DocumentAgencyReferences
		documentAgencyReferencesRouter.DELETE("deleteDocumentAgencyReferencesByIds", documentAgencyReferencesApi.DeleteDocumentAgencyReferencesByIds) // 批量删除DocumentAgencyReferences
		documentAgencyReferencesRouter.PUT("updateDocumentAgencyReferences", documentAgencyReferencesApi.UpdateDocumentAgencyReferences)              // 更新DocumentAgencyReferences
	}
	{
		documentAgencyReferencesRouterWithoutRecord.GET("findDocumentAgencyReferences", documentAgencyReferencesApi.FindDocumentAgencyReferences)       // 根据ID获取DocumentAgencyReferences
		documentAgencyReferencesRouterWithoutRecord.GET("getDocumentAgencyReferencesList", documentAgencyReferencesApi.GetDocumentAgencyReferencesList) // 获取DocumentAgencyReferences列表
	}
}
