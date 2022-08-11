package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SystemApiApi struct{}

// CreateApi create new API
// @Tags SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysApi true "api路径, api中文描述, api组, 方法"
// @Success 200 {object} response.Response{msg=string} "创建基础api"
// @Router /api/createApi [post]
func (s *SystemApiApi) CreateApi(c *gin.Context) {
	var api system.SysApi
	var err error

	err = c.ShouldBindJSON(&api)
	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := utils.Verify(api, utils.ApiVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := apiService.CreateApi(api); err != nil {
		global.GVA_LOG.Error("fail to create new api record", zap.Error(err))
		response.FailWithMessage("fail to create new api record", c)
	} else {
		response.OkWithMessage("create new api record success", c)
	}
}

// DeleteApi delete api
// @Tags SysApi
// @Summary 删除api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysApi true "ID"
// @Success 200 {object} response.Response{msg=string} "删除api"
// @Router /api/deleteApi [post]
func (s *SystemApiApi) DeleteApi(c *gin.Context) {
	var api system.SysApi
	var err error

	err = c.ShouldBindJSON(&api)
	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := utils.Verify(api.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := apiService.DeleteApi(api); err != nil {
		global.GVA_LOG.Error("fail to delete api record", zap.Error(err))
		response.FailWithMessage("fail to delete api record", c)
	} else {
		response.OkWithMessage("delete api record successfully", c)
	}
}

// GetApiList get api list
// @Tags SysApi
// @Summary 分页获取API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.SearchApiParams true "分页获取API列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取API列表,返回包括列表,总数,页码,每页数量"
// @Router /api/getApiList [post]
func (s *SystemApiApi) GetApiList(c *gin.Context) {
	var pageInfo systemReq.SearchApiParams
	var err error

	err = c.ShouldBindJSON(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := apiService.GetAPIInfoList(pageInfo.SysApi, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc); err != nil {
		global.GVA_LOG.Error("fail to get api list", zap.Error(err))
		response.FailWithMessage("fail to get api list", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "get api list successfully", c)
	}
}

// GetApiById get api by ID
// @Tags SysApi
// @Summary 根据id获取api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取api"
// @Success 200 {object} response.Response{data=systemRes.SysAPIResponse} "根据id获取api,返回包括api详情"
// @Router /api/getApiById [post]
func (s *SystemApiApi) GetApiById(c *gin.Context) {
	var idInfo request.GetById
	var err error

	err = c.ShouldBindJSON(&idInfo)
	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := utils.Verify(idInfo, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	api, err := apiService.GetApiById(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("fail to get api by ID", zap.Error(err))
		response.FailWithMessage("fail to get api by ID", c)
	} else {
		response.OkWithDetailed(systemRes.SysAPIResponse{Api: api}, "success", c)
	}
}

// UpdateApi update api
// @Tags SysApi
// @Summary 修改基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysApi true "api路径, api中文描述, api组, 方法"
// @Success 200 {object} response.Response{msg=string} "修改基础api"
// @Router /api/updateApi [post]
func (s *SystemApiApi) UpdateApi(c *gin.Context) {
	var api system.SysApi
	var err error

	err = c.ShouldBindJSON(&api)
	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := utils.Verify(api, utils.ApiVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := apiService.UpdateApi(api); err != nil {
		global.GVA_LOG.Error("fail to update api", zap.Error(err))
		response.FailWithMessage("fail to update api", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// GetAllApis get all apis
// @Tags SysApi
// @Summary 获取所有的Api 不分页
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=systemRes.SysAPIListResponse,msg=string} "获取所有的Api 不分页,返回包括api列表"
// @Router /api/getAllApis [post]
func (s *SystemApiApi) GetAllApis(c *gin.Context) {
	if apis, err := apiService.GetAllApis(); err != nil {
		global.GVA_LOG.Error("fail to get all api", zap.Error(err))
		response.FailWithMessage("fail to get all api", c)
	} else {
		response.OkWithDetailed(systemRes.SysAPIListResponse{Apis: apis}, "success", c)
	}
}

// DeleteApisByIds bulk delete API
// @Tags SysApi
// @Summary 删除选中Api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "ID"
// @Success 200 {object} response.Response{msg=string} "删除选中Api"
// @Router /api/deleteApisByIds [delete]
func (s *SystemApiApi) DeleteApisByIds(c *gin.Context) {
	var ids request.IdsReq
	var err error

	err = c.ShouldBindJSON(&ids)
	if err != nil {
		global.GVA_LOG.Error("provide valid data", zap.Error(err))
		response.FailWithMessage("provide valid data", c)
		return
	}

	if err := apiService.DeleteApisByIds(ids); err != nil {
		global.GVA_LOG.Error("fail to bulk delete api by IDs", zap.Error(err))
		response.FailWithMessage("fail to bulk delete api by IDs", c)
	} else {
		response.OkWithMessage("success", c)
	}
}
