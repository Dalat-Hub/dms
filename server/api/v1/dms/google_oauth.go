package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GoogleOAuthApi struct {
}

var googleOAuthService = service.ServiceGroupApp.DmsServiceGroup.GoogleOAuthService

func (googleOAuthApi *GoogleOAuthApi) HandleGoogleOAuth(ctx *gin.Context) {
	code := ctx.Query("code")

	if code == "" {
		ctx.Redirect(http.StatusTemporaryRedirect, global.GVA_CONFIG.GoogleOAuth.ClientRedirectURL+"1")
		return
	}

	token, err := googleOAuthService.HandleGoogleOAuth(code)
	if err != nil {
		ctx.Redirect(http.StatusTemporaryRedirect, global.GVA_CONFIG.GoogleOAuth.ClientRedirectURL+"1")
		return
	}

	ctx.Redirect(http.StatusTemporaryRedirect, global.GVA_CONFIG.GoogleOAuth.ClientRedirectURL+token)
}
