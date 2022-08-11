package system

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AutoCodeApi struct{}

var caser = cases.Title(language.English)

// PreviewTemp
// @Tags AutoCode
// @Summary 预览创建后的代码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.AutoCodeStruct true "预览创建代码"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "预览创建后的代码"
// @Router /autoCode/preview [post]
func (autoApi *AutoCodeApi) PreviewTemp(c *gin.Context) {
	var a system.AutoCodeStruct
	var err error
	err = c.ShouldBindJSON(&a)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := utils.Verify(a, utils.AutoCodeVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	a.Pretreatment() // 处理go关键字
	a.PackageT = caser.String(a.Package)
	autoCode, err := autoCodeService.PreviewTemp(a)
	if err != nil {
		global.GVA_LOG.Error("fail to preview", zap.Error(err))
		response.FailWithMessage("fail to preview", c)
	} else {
		response.OkWithDetailed(gin.H{"autoCode": autoCode}, "success", c)
	}
}

// CreateTemp
// @Tags AutoCode
// @Summary 自动代码模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.AutoCodeStruct true "创建自动代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/createTemp [post]
func (autoApi *AutoCodeApi) CreateTemp(c *gin.Context) {
	var a system.AutoCodeStruct
	var err error

	err = c.ShouldBindJSON(&a)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := utils.Verify(a, utils.AutoCodeVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	a.Pretreatment()
	var apiIds []uint
	if a.AutoCreateApiToSql {
		if ids, err := autoCodeService.AutoCreateApi(&a); err != nil {
			global.GVA_LOG.Error("Automatic creation failed! Please clear the junk data yourself!", zap.Error(err))
			c.Writer.Header().Add("success", "false")
			c.Writer.Header().Add("msg", url.QueryEscape("Automatic creation failed! Please clear the junk data yourself!"))
			return
		} else {
			apiIds = ids
		}
	}
	a.PackageT = caser.String(a.Package)
	err = autoCodeService.CreateTemp(a, apiIds...)
	if err != nil {
		if errors.Is(err, system.AutoMoveErr) {
			c.Writer.Header().Add("success", "true")
			c.Writer.Header().Add("msg", url.QueryEscape(err.Error()))
		} else {
			c.Writer.Header().Add("success", "false")
			c.Writer.Header().Add("msg", url.QueryEscape(err.Error()))
			_ = os.Remove("./ginvueadmin.zip")
		}
	} else {
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "ginvueadmin.zip")) // fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.Header().Add("success", "true")
		c.File("./ginvueadmin.zip")
		_ = os.Remove("./ginvueadmin.zip")
	}
}

// GetDB
// @Tags AutoCode
// @Summary 获取当前所有数据库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取当前所有数据库"
// @Router /autoCode/getDatabase [get]
func (autoApi *AutoCodeApi) GetDB(c *gin.Context) {
	dbs, err := autoCodeService.Database().GetDB()
	if err != nil {
		global.GVA_LOG.Error("fail to get db", zap.Error(err))
		response.FailWithMessage("fail to get db", c)
	} else {
		response.OkWithDetailed(gin.H{"dbs": dbs}, "success", c)
	}
}

// GetTables
// @Tags AutoCode
// @Summary 获取当前数据库所有表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取当前数据库所有表"
// @Router /autoCode/getTables [get]
func (autoApi *AutoCodeApi) GetTables(c *gin.Context) {
	dbName := c.DefaultQuery("dbName", global.GVA_CONFIG.Mysql.Dbname)
	tables, err := autoCodeService.Database().GetTables(dbName)
	if err != nil {
		global.GVA_LOG.Error("fail to get table", zap.Error(err))
		response.FailWithMessage("fail to get table", c)
	} else {
		response.OkWithDetailed(gin.H{"tables": tables}, "success", c)
	}
}

// GetColumn
// @Tags AutoCode
// @Summary 获取当前表所有字段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取当前表所有字段"
// @Router /autoCode/getColumn [get]
func (autoApi *AutoCodeApi) GetColumn(c *gin.Context) {
	dbName := c.DefaultQuery("dbName", global.GVA_CONFIG.Mysql.Dbname)
	tableName := c.Query("tableName")
	columns, err := autoCodeService.Database().GetColumn(tableName, dbName)
	if err != nil {
		global.GVA_LOG.Error("fail to get column", zap.Error(err))
		response.FailWithMessage("fail to get column", c)
	} else {
		response.OkWithDetailed(gin.H{"columns": columns}, "success", c)
	}
}

// CreatePackage
// @Tags AutoCode
// @Summary 创建package
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAutoCode true "创建package"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "创建package成功"
// @Router /autoCode/createPackage [post]
func (autoApi *AutoCodeApi) CreatePackage(c *gin.Context) {
	var a system.SysAutoCode
	var err error

	err = c.ShouldBindJSON(&a)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := utils.Verify(a, utils.AutoPackageVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = autoCodeService.CreateAutoCode(&a)
	if err != nil {
		global.GVA_LOG.Error("fail to create new package", zap.Error(err))
		response.FailWithMessage("fail to create new package", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// GetPackage
// @Tags AutoCode
// @Summary 获取package
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "创建package成功"
// @Router /autoCode/getPackage [post]
func (autoApi *AutoCodeApi) GetPackage(c *gin.Context) {
	pkgs, err := autoCodeService.GetPackage()
	if err != nil {
		global.GVA_LOG.Error("fail to get package", zap.Error(err))
		response.FailWithMessage("fail to get package", c)
	} else {
		response.OkWithDetailed(gin.H{"pkgs": pkgs}, "success", c)
	}
}

// DelPackage
// @Tags AutoCode
// @Summary 删除package
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAutoCode true "创建package"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "删除package成功"
// @Router /autoCode/delPackage [post]
func (autoApi *AutoCodeApi) DelPackage(c *gin.Context) {
	var a system.SysAutoCode
	var err error

	err = c.ShouldBindJSON(&a)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	err = autoCodeService.DelPackage(a)
	if err != nil {
		global.GVA_LOG.Error("fail to delete package", zap.Error(err))
		response.FailWithMessage("fail to delete package", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// AutoPlug
// @Tags AutoCode
// @Summary 创建插件模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAutoCode true "创建插件模板"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "创建插件模板成功"
// @Router /autoCode/createPlug [post]
func (autoApi *AutoCodeApi) AutoPlug(c *gin.Context) {
	var a system.AutoPlugReq
	var err error

	err = c.ShouldBindJSON(&a)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	a.Snake = strings.ToLower(a.PlugName)
	a.NeedModel = a.HasRequest || a.HasResponse

	err = autoCodeService.CreatePlug(a)
	if err != nil {
		global.GVA_LOG.Error("fail to create plug", zap.Error(err))
		response.FailWithMessage("fail to create plug", c)
	} else {
		response.Ok(c)
	}
}

func (autoApi *AutoCodeApi) InstallPlugin(c *gin.Context) {
	header, err := c.FormFile("plug")
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	web, server, err := autoCodeService.InstallPlugin(header)
	webStr := "web plugin installed successfully"
	serverStr := "server plugin installed successfully"
	if web == -1 {
		webStr = "the web-side plug-in is not successfully installed, please decompress and install it according to the documentation. If it is a pure back-end plug-in, please ignore this prompt"
	}
	if server == -1 {
		serverStr = "the server-side plug-in is not successfully installed, please decompress and install it according to the documentation. If it is a pure front-end plug-in, please ignore this prompt"
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData([]interface{}{
			gin.H{
				"code": web,
				"msg":  webStr,
			},
			gin.H{
				"code": server,
				"msg":  serverStr,
			}}, c)
	}
}
