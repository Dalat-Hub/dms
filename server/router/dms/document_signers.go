package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DocumentSignersRouter struct {
}

// InitDocumentSignersRouter 初始化 DocumentSigners 路由信息
func (s *DocumentSignersRouter) InitDocumentSignersRouter(Router *gin.RouterGroup) {
	documentSignersRouter := Router.Group("documentSigners").Use(middleware.OperationRecord())
	documentSignersRouterWithoutRecord := Router.Group("documentSigners")
	var documentSignersApi = v1.ApiGroupApp.DmsApiGroup.DocumentSignersApi
	{
		documentSignersRouter.POST("createDocumentSigners", documentSignersApi.CreateDocumentSigners)             // 新建DocumentSigners
		documentSignersRouter.DELETE("deleteDocumentSigners", documentSignersApi.DeleteDocumentSigners)           // 删除DocumentSigners
		documentSignersRouter.DELETE("deleteDocumentSignersByIds", documentSignersApi.DeleteDocumentSignersByIds) // 批量删除DocumentSigners
		documentSignersRouter.PUT("updateDocumentSigners", documentSignersApi.UpdateDocumentSigners)              // 更新DocumentSigners
	}
	{
		documentSignersRouterWithoutRecord.GET("findDocumentSigners", documentSignersApi.FindDocumentSigners)       // 根据ID获取DocumentSigners
		documentSignersRouterWithoutRecord.GET("getDocumentSignersList", documentSignersApi.GetDocumentSignersList) // 获取DocumentSigners列表
	}
}
