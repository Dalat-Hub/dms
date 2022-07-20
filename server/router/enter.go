package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/dms"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Dms     dms.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
