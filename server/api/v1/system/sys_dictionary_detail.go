package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DictionaryDetailApi struct{}

// CreateSysDictionaryDetail
// @Tags SysDictionaryDetail
// @Summary 创建SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictionaryDetail true "SysDictionaryDetail模型"
// @Success 200 {object} response.Response{msg=string} "创建SysDictionaryDetail"
// @Router /sysDictionaryDetail/createSysDictionaryDetail [post]
func (s *DictionaryDetailApi) CreateSysDictionaryDetail(c *gin.Context) {
	var detail system.SysDictionaryDetail
	var err error

	err = c.ShouldBindJSON(&detail)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := dictionaryDetailService.CreateSysDictionaryDetail(detail); err != nil {
		global.GVA_LOG.Error("fail to create dictionary detail", zap.Error(err))
		response.FailWithMessage("fail to create dictionary detail", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// DeleteSysDictionaryDetail
// @Tags SysDictionaryDetail
// @Summary 删除SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictionaryDetail true "SysDictionaryDetail模型"
// @Success 200 {object} response.Response{msg=string} "删除SysDictionaryDetail"
// @Router /sysDictionaryDetail/deleteSysDictionaryDetail [delete]
func (s *DictionaryDetailApi) DeleteSysDictionaryDetail(c *gin.Context) {
	var detail system.SysDictionaryDetail
	var err error

	err = c.ShouldBindJSON(&detail)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := dictionaryDetailService.DeleteSysDictionaryDetail(detail); err != nil {
		global.GVA_LOG.Error("fail to delete dictionary detail", zap.Error(err))
		response.FailWithMessage("fail to delete dictionary detail", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// UpdateSysDictionaryDetail
// @Tags SysDictionaryDetail
// @Summary 更新SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictionaryDetail true "更新SysDictionaryDetail"
// @Success 200 {object} response.Response{msg=string} "更新SysDictionaryDetail"
// @Router /sysDictionaryDetail/updateSysDictionaryDetail [put]
func (s *DictionaryDetailApi) UpdateSysDictionaryDetail(c *gin.Context) {
	var detail system.SysDictionaryDetail
	var err error

	err = c.ShouldBindJSON(&detail)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := dictionaryDetailService.UpdateSysDictionaryDetail(&detail); err != nil {
		global.GVA_LOG.Error("fail to update dictionary detail", zap.Error(err))
		response.FailWithMessage("fail to update dictionary detail", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// FindSysDictionaryDetail
// @Tags SysDictionaryDetail
// @Summary 用id查询SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query system.SysDictionaryDetail true "用id查询SysDictionaryDetail"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "用id查询SysDictionaryDetail"
// @Router /sysDictionaryDetail/findSysDictionaryDetail [get]
func (s *DictionaryDetailApi) FindSysDictionaryDetail(c *gin.Context) {
	var detail system.SysDictionaryDetail
	var err error

	err = c.ShouldBindQuery(&detail)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := utils.Verify(detail, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reSysDictionaryDetail, err := dictionaryDetailService.GetSysDictionaryDetail(detail.ID); err != nil {
		global.GVA_LOG.Error("fail to get dictionary detail", zap.Error(err))
		response.FailWithMessage("fail to get dictionary detail", c)
	} else {
		response.OkWithDetailed(gin.H{"reSysDictionaryDetail": reSysDictionaryDetail}, "success", c)
	}
}

// GetSysDictionaryDetailList
// @Tags SysDictionaryDetail
// @Summary 分页获取SysDictionaryDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.SysDictionaryDetailSearch true "页码, 每页大小, 搜索条件"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取SysDictionaryDetail列表,返回包括列表,总数,页码,每页数量"
// @Router /sysDictionaryDetail/getSysDictionaryDetailList [get]
func (s *DictionaryDetailApi) GetSysDictionaryDetailList(c *gin.Context) {
	var pageInfo request.SysDictionaryDetailSearch
	var err error

	err = c.ShouldBindQuery(&pageInfo)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if list, total, err := dictionaryDetailService.GetSysDictionaryDetailInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("fail to get list of dictionary details", zap.Error(err))
		response.FailWithMessage("fail to get list of dictionary details", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "success", c)
	}
}
