package dms

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type GoogleOAuthRouter struct {
}

// InitGoogleOAuthRouter Initialize Google OAuth routing information
func (s *DocumentsRouter) InitGoogleOAuthRouter(Router *gin.RouterGroup) {
	router := Router.Group("")
	var googleOAuthApi = v1.ApiGroupApp.DmsApiGroup.GoogleOAuthApi
	{
		router.GET("/api/sessions/oauth/google", googleOAuthApi.HandleGoogleOAuth)
	}
}
