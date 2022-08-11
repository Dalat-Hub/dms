package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OperationRecordApi struct{}

// CreateSysOperationRecord
// @Tags SysOperationRecord
// @Summary 创建SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysOperationRecord true "创建SysOperationRecord"
// @Success 200 {object} response.Response{msg=string} "创建SysOperationRecord"
// @Router /sysOperationRecord/createSysOperationRecord [post]
func (s *OperationRecordApi) CreateSysOperationRecord(c *gin.Context) {
	var sysOperationRecord system.SysOperationRecord
	var err error

	err = c.ShouldBindJSON(&sysOperationRecord)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := operationRecordService.CreateSysOperationRecord(sysOperationRecord); err != nil {
		global.GVA_LOG.Error("fail to create operation record", zap.Error(err))
		response.FailWithMessage("fail to create operation record", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// DeleteSysOperationRecord
// @Tags SysOperationRecord
// @Summary 删除SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysOperationRecord true "SysOperationRecord模型"
// @Success 200 {object} response.Response{msg=string} "删除SysOperationRecord"
// @Router /sysOperationRecord/deleteSysOperationRecord [delete]
func (s *OperationRecordApi) DeleteSysOperationRecord(c *gin.Context) {
	var sysOperationRecord system.SysOperationRecord
	var err error

	err = c.ShouldBindJSON(&sysOperationRecord)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := operationRecordService.DeleteSysOperationRecord(sysOperationRecord); err != nil {
		global.GVA_LOG.Error("fail to delete operation record", zap.Error(err))
		response.FailWithMessage("fail to delete operation record", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// DeleteSysOperationRecordByIds
// @Tags SysOperationRecord
// @Summary 批量删除SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysOperationRecord"
// @Success 200 {object} response.Response{msg=string} "批量删除SysOperationRecord"
// @Router /sysOperationRecord/deleteSysOperationRecordByIds [delete]
func (s *OperationRecordApi) DeleteSysOperationRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	var err error

	err = c.ShouldBindJSON(&IDS)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := operationRecordService.DeleteSysOperationRecordByIds(IDS); err != nil {
		global.GVA_LOG.Error("fail to delete operation record", zap.Error(err))
		response.FailWithMessage("fail to delete operation record", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// FindSysOperationRecord
// @Tags SysOperationRecord
// @Summary 用id查询SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query system.SysOperationRecord true "Id"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "用id查询SysOperationRecord"
// @Router /sysOperationRecord/findSysOperationRecord [get]
func (s *OperationRecordApi) FindSysOperationRecord(c *gin.Context) {
	var sysOperationRecord system.SysOperationRecord
	var err error

	err = c.ShouldBindQuery(&sysOperationRecord)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := utils.Verify(sysOperationRecord, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reSysOperationRecord, err := operationRecordService.GetSysOperationRecord(sysOperationRecord.ID); err != nil {
		global.GVA_LOG.Error("fail to get operation record", zap.Error(err))
		response.FailWithMessage("fail to get operation record", c)
	} else {
		response.OkWithDetailed(gin.H{"reSysOperationRecord": reSysOperationRecord}, "success", c)
	}
}

// GetSysOperationRecordList
// @Tags SysOperationRecord
// @Summary 分页获取SysOperationRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.SysOperationRecordSearch true "页码, 每页大小, 搜索条件"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取SysOperationRecord列表,返回包括列表,总数,页码,每页数量"
// @Router /sysOperationRecord/getSysOperationRecordList [get]
func (s *OperationRecordApi) GetSysOperationRecordList(c *gin.Context) {
	var pageInfo systemReq.SysOperationRecordSearch
	var err error

	err = c.ShouldBindQuery(&pageInfo)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if list, total, err := operationRecordService.GetSysOperationRecordInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("fail to get operation record list", zap.Error(err))
		response.FailWithMessage("fail to get operation record list", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "success", c)
	}
}
