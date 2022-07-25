package middleware

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// CasbinHandler Interceptor
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		waitUse, _ := utils.GetClaims(c)

		// Get the requested PATh
		obj := c.Request.URL.Path

		// Get the requested METHOD
		act := c.Request.Method

		// Get the user's ROLE
		sub := strconv.Itoa(int(waitUse.AuthorityId))

		e := casbinService.Casbin() // Determine whether the strategy exists
		success, _ := e.Enforce(sub, obj, act)
		if global.GVA_CONFIG.System.Env == "develop" || success {
			c.Next()
		} else {
			response.FailWithDetailed(gin.H{}, "Bạn không có đủ quyền để thực hiện hành động này", c)
			c.Abort()
			return
		}
	}
}
