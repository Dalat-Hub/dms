package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SystemApi struct{}

// GetSystemConfig
// @Tags System
// @Summary 获取配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {object} response.Response{data=systemRes.SysConfigResponse,msg=string} "获取配置文件内容,返回包括系统配置"
// @Router /system/getSystemConfig [post]
func (s *SystemApi) GetSystemConfig(c *gin.Context) {
	if config, err := systemConfigService.GetSystemConfig(); err != nil {
		global.GVA_LOG.Error("fail to get system config", zap.Error(err))
		response.FailWithMessage("fail to get system config", c)
	} else {
		response.OkWithDetailed(systemRes.SysConfigResponse{Config: config}, "success", c)
	}
}

// SetSystemConfig
// @Tags System
// @Summary 设置配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body system.System true "设置配置文件内容"
// @Success 200 {object} response.Response{data=string} "设置配置文件内容"
// @Router /system/setSystemConfig [post]
func (s *SystemApi) SetSystemConfig(c *gin.Context) {
	var sys system.System
	var err error

	err = c.ShouldBindJSON(&sys)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := systemConfigService.SetSystemConfig(sys); err != nil {
		global.GVA_LOG.Error("fail to set system config", zap.Error(err))
		response.FailWithMessage("fail to set system config", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// ReloadSystem
// @Tags System
// @Summary 重启系统
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {object} response.Response{msg=string} "重启系统"
// @Router /system/reloadSystem [post]
func (s *SystemApi) ReloadSystem(c *gin.Context) {
	err := utils.Reload()
	if err != nil {
		global.GVA_LOG.Error("fail to reload system", zap.Error(err))
		response.FailWithMessage("fail to reload system", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// GetServerInfo
// @Tags System
// @Summary 获取服务器信息
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取服务器信息"
// @Router /system/getServerInfo [post]
func (s *SystemApi) GetServerInfo(c *gin.Context) {
	if server, err := systemConfigService.GetServerInfo(); err != nil {
		global.GVA_LOG.Error("fail to get server info", zap.Error(err))
		response.FailWithMessage("fail to get server info", c)
	} else {
		response.OkWithDetailed(gin.H{"server": server}, "success", c)
	}
}
