package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AuthorityRouter struct{}

func (s *AuthorityRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	authorityRouter := Router.Group("authority").Use(middleware.OperationRecord())
	authorityRouterWithoutRecord := Router.Group("authority")
	authorityApi := v1.ApiGroupApp.SystemApiGroup.AuthorityApi
	{
		authorityRouter.POST("createAuthority", authorityApi.CreateAuthority)
		authorityRouter.POST("deleteAuthority", authorityApi.DeleteAuthority)
		authorityRouter.PUT("updateAuthority", authorityApi.UpdateAuthority)
		authorityRouter.POST("copyAuthority", authorityApi.CopyAuthority)
		authorityRouter.POST("setDataAuthority", authorityApi.SetDataAuthority)
	}
	{
		authorityRouterWithoutRecord.POST("getAuthorityList", authorityApi.GetAuthorityList)
		authorityRouterWithoutRecord.POST("getAuthorityInfo", authorityApi.GetAuthorityInfo)
	}
}
