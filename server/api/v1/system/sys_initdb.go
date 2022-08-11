package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type DBApi struct{}

// InitDB
// @Tags InitDB
// @Summary 初始化用户数据库
// @Produce  application/json
// @Param data body request.InitDB true "初始化数据库参数"
// @Success 200 {object} response.Response{data=string} "初始化用户数据库"
// @Router /init/initdb [post]
func (i *DBApi) InitDB(c *gin.Context) {
	if global.GVA_DB != nil {
		global.GVA_LOG.Error("database configuration already exists!")
		response.FailWithMessage("database configuration already exists!", c)
		return
	}
	var dbInfo request.InitDB
	if err := c.ShouldBindJSON(&dbInfo); err != nil {
		global.GVA_LOG.Error("parameter verification failed!", zap.Error(err))
		response.FailWithMessage("parameter verification failed!", c)
		return
	}
	if err := initDBService.InitDB(dbInfo); err != nil {
		global.GVA_LOG.Error("automatic database creation failed!", zap.Error(err))
		response.FailWithMessage("failed to create database automatically, please check the background log and initialize after checking", c)
		return
	}
	response.OkWithMessage("automatic database creation succeeded", c)
}

// CheckDB
// @Tags CheckDB
// @Summary 初始化用户数据库
// @Produce  application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "初始化用户数据库"
// @Router /init/checkdb [post]
func (i *DBApi) CheckDB(c *gin.Context) {
	var (
		message  = "go to initialize database"
		needInit = true
	)

	if global.GVA_DB != nil {
		message = "the database does not need to be initialized"
		needInit = false
	}
	global.GVA_LOG.Info(message)
	response.OkWithDetailed(gin.H{"needInit": needInit}, message, c)
	return
}
