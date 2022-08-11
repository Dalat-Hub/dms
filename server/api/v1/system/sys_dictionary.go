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

type DictionaryApi struct{}

// CreateSysDictionary
// @Tags SysDictionary
// @Summary 创建SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictionary true "SysDictionary模型"
// @Success 200 {object} response.Response{msg=string} "创建SysDictionary"
// @Router /sysDictionary/createSysDictionary [post]
func (s *DictionaryApi) CreateSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	var err error

	err = c.ShouldBindJSON(&dictionary)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := dictionaryService.CreateSysDictionary(dictionary); err != nil {
		global.GVA_LOG.Error("fail to create new dictionary", zap.Error(err))
		response.FailWithMessage("fail to create new dictionary", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// DeleteSysDictionary
// @Tags SysDictionary
// @Summary 删除SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictionary true "SysDictionary模型"
// @Success 200 {object} response.Response{msg=string} "删除SysDictionary"
// @Router /sysDictionary/deleteSysDictionary [delete]
func (s *DictionaryApi) DeleteSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	var err error

	err = c.ShouldBindJSON(&dictionary)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := dictionaryService.DeleteSysDictionary(dictionary); err != nil {
		global.GVA_LOG.Error("fail to delete dictionary", zap.Error(err))
		response.FailWithMessage("fail to delete dictionary", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// UpdateSysDictionary
// @Tags SysDictionary
// @Summary 更新SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictionary true "SysDictionary模型"
// @Success 200 {object} response.Response{msg=string} "更新SysDictionary"
// @Router /sysDictionary/updateSysDictionary [put]
func (s *DictionaryApi) UpdateSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	var err error

	err = c.ShouldBindJSON(&dictionary)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := dictionaryService.UpdateSysDictionary(&dictionary); err != nil {
		global.GVA_LOG.Error("fail to update dictionary", zap.Error(err))
		response.FailWithMessage("fail to update dictionary", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// FindSysDictionary
// @Tags SysDictionary
// @Summary 用id查询SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query system.SysDictionary true "ID或字典英名"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "用id查询SysDictionary"
// @Router /sysDictionary/findSysDictionary [get]
func (s *DictionaryApi) FindSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	var err error

	err = c.ShouldBindQuery(&dictionary)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if sysDictionary, err := dictionaryService.GetSysDictionary(dictionary.Type, dictionary.ID, dictionary.Status); err != nil {
		global.GVA_LOG.Error("fail to get dictionary", zap.Error(err))
		response.FailWithMessage("fail to get dictionary", c)
	} else {
		response.OkWithDetailed(gin.H{"resysDictionary": sysDictionary}, "success", c)
	}
}

// GetSysDictionaryList
// @Tags SysDictionary
// @Summary 分页获取SysDictionary列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.SysDictionarySearch true "页码, 每页大小, 搜索条件"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取SysDictionary列表,返回包括列表,总数,页码,每页数量"
// @Router /sysDictionary/getSysDictionaryList [get]
func (s *DictionaryApi) GetSysDictionaryList(c *gin.Context) {
	var pageInfo request.SysDictionarySearch
	var err error

	err = c.ShouldBindQuery(&pageInfo)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := dictionaryService.GetSysDictionaryInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("fail to get list of dictionaries", zap.Error(err))
		response.FailWithMessage("fail to get list of dictionaries", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "success", c)
	}
}
