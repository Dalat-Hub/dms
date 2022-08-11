package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthorityBtnApi struct{}

// GetAuthorityBtn get button
// @Tags AuthorityBtn
// @Summary 获取权限按钮
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SysAuthorityBtnReq true "菜单id, 角色id, 选中的按钮id"
// @Success 200 {object} response.Response{data=response.SysAuthorityBtnRes,msg=string} "返回列表成功"
// @Router /authorityBtn/getAuthorityBtn [post]
func (a *AuthorityBtnApi) GetAuthorityBtn(c *gin.Context) {
	var req request.SysAuthorityBtnReq
	var err error

	err = c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if res, err := authorityBtnService.GetAuthorityBtn(req); err != nil {
		global.GVA_LOG.Error("fail to get button", zap.Error(err))
		response.FailWithMessage("fail to get button", c)
	} else {
		response.OkWithDetailed(res, "success", c)
	}
}

// SetAuthorityBtn set button
// @Tags AuthorityBtn
// @Summary 设置权限按钮
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SysAuthorityBtnReq true "菜单id, 角色id, 选中的按钮id"
// @Success 200 {object} response.Response{msg=string} "返回列表成功"
// @Router /authorityBtn/setAuthorityBtn [post]
func (a *AuthorityBtnApi) SetAuthorityBtn(c *gin.Context) {
	var req request.SysAuthorityBtnReq
	var err error

	err = c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := authorityBtnService.SetAuthorityBtn(req); err != nil {
		global.GVA_LOG.Error("fail to get authority button", zap.Error(err))
		response.FailWithMessage("fail to get authority button", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// CanRemoveAuthorityBtn
// @Tags AuthorityBtn
// @Summary 设置权限按钮
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /authorityBtn/canRemoveAuthorityBtn [post]
func (a *AuthorityBtnApi) CanRemoveAuthorityBtn(c *gin.Context) {
	id := c.Query("id")
	if err := authorityBtnService.CanRemoveAuthorityBtn(id); err != nil {
		global.GVA_LOG.Error("fail to remove authority button", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("success", c)
	}
}
