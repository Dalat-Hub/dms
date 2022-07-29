package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DocumentUsersRouter struct {
}

func (s *DocumentUsersRouter) InitDocumentUsersRouter(Router *gin.RouterGroup) {
	documentUsersRouter := Router.Group("documentUsers").Use(middleware.OperationRecord())
	documentUsersRouterWithoutRecord := Router.Group("documentUsers")
	var documentUsersApi = v1.ApiGroupApp.DmsApiGroup.DocumentUsersApi
	{
		documentUsersRouter.POST("createDocumentUsers", documentUsersApi.CreateDocumentUsers)
		documentUsersRouter.DELETE("deleteDocumentUsers", documentUsersApi.DeleteDocumentUsers)
		documentUsersRouter.DELETE("deleteDocumentUsersByIds", documentUsersApi.DeleteDocumentUsersByIds)
		documentUsersRouter.PUT("updateDocumentUsers", documentUsersApi.UpdateDocumentUsers)
		documentUsersRouter.POST("distributeTasks", documentUsersApi.DistributeTasks)
	}
	{
		documentUsersRouterWithoutRecord.GET("findDocumentUsers", documentUsersApi.FindDocumentUsers)
		documentUsersRouterWithoutRecord.GET("getDocumentUsersList", documentUsersApi.GetDocumentUsersList)
	}
}
