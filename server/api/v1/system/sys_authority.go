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

type AuthorityApi struct{}

// CreateAuthority create new authority
// @Tags Authority
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAuthority true "权限id, 权限名, 父角色id"
// @Success 200 {object} response.Response{data=systemRes.SysAuthorityResponse,msg=string} "创建角色,返回包括系统角色详情"
// @Router /authority/createAuthority [post]
func (a *AuthorityApi) CreateAuthority(c *gin.Context) {
	var authority system.SysAuthority
	var err error

	err = c.ShouldBindJSON(&authority)
	if err != nil {
		global.GVA_LOG.Error("provide valid data", zap.Error(err))
		response.FailWithMessage("provide valid data"+err.Error(), c)
		return
	}

	if err := utils.Verify(authority, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if authBack, err := authorityService.CreateAuthority(authority); err != nil {
		global.GVA_LOG.Error("fail to create new authority", zap.Error(err))
		response.FailWithMessage("fail to create new authority"+err.Error(), c)
	} else {
		_ = menuService.AddMenuAuthority(systemReq.DefaultMenu(), authority.AuthorityId)
		_ = casbinService.UpdateCasbin(authority.AuthorityId, systemReq.DefaultCasbin())
		response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authBack}, "success", c)
	}
}

// CopyAuthority copy authority
// @Tags Authority
// @Summary 拷贝角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body response.SysAuthorityCopyResponse true "旧角色id, 新权限id, 新权限名, 新父角色id"
// @Success 200 {object} response.Response{data=systemRes.SysAuthorityResponse,msg=string} "拷贝角色,返回包括系统角色详情"
// @Router /authority/copyAuthority [post]
func (a *AuthorityApi) CopyAuthority(c *gin.Context) {
	var copyInfo systemRes.SysAuthorityCopyResponse
	var err error

	err = c.ShouldBindJSON(&copyInfo)
	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data"+err.Error(), c)
		return
	}

	if err := utils.Verify(copyInfo, utils.OldAuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(copyInfo.Authority, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if authBack, err := authorityService.CopyAuthority(copyInfo); err != nil {
		global.GVA_LOG.Error("fail to copy authority", zap.Error(err))
		response.FailWithMessage("fail to copy authority"+err.Error(), c)
	} else {
		response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authBack}, "success", c)
	}
}

// DeleteAuthority delete authority
// @Tags Authority
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAuthority true "删除角色"
// @Success 200 {object} response.Response{msg=string} "删除角色"
// @Router /authority/deleteAuthority [post]
func (a *AuthorityApi) DeleteAuthority(c *gin.Context) {
	var authority system.SysAuthority
	var err error

	err = c.ShouldBindJSON(&authority)
	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data"+err.Error(), c)
		return
	}

	if err := utils.Verify(authority, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := authorityService.DeleteAuthority(&authority); err != nil { // 删除角色之前需要判断是否有用户正在使用此角色
		global.GVA_LOG.Error("fail to delete authority", zap.Error(err))
		response.FailWithMessage("fail to delete authority"+err.Error(), c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// UpdateAuthority update authority
// @Tags Authority
// @Summary 更新角色信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAuthority true "权限id, 权限名, 父角色id"
// @Success 200 {object} response.Response{data=systemRes.SysAuthorityResponse,msg=string} "更新角色信息,返回包括系统角色详情"
// @Router /authority/updateAuthority [post]
func (a *AuthorityApi) UpdateAuthority(c *gin.Context) {
	var auth system.SysAuthority
	var err error

	err = c.ShouldBindJSON(&auth)
	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data"+err.Error(), c)
		return
	}

	if err := utils.Verify(auth, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if authority, err := authorityService.UpdateAuthority(auth); err != nil {
		global.GVA_LOG.Error("fail to update authority", zap.Error(err))
		response.FailWithMessage("fail to update authority"+err.Error(), c)
	} else {
		response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authority}, "success", c)
	}
}

// GetAuthorityList get list of authorities
// @Tags Authority
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取角色列表,返回包括列表,总数,页码,每页数量"
// @Router /authority/getAuthorityList [post]
func (a *AuthorityApi) GetAuthorityList(c *gin.Context) {
	var pageInfo request.PageInfo
	var err error

	err = c.ShouldBindJSON(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data"+err.Error(), c)
		return
	}

	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := authorityService.GetAuthorityInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("fail to get list of authorities", zap.Error(err))
		response.FailWithMessage("fail to get list of authorities"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "success", c)
	}
}

func (a *AuthorityApi) GetAuthorityInfo(c *gin.Context) {
	var auth system.SysAuthority
	var err error

	err = c.ShouldBindJSON(&auth)
	if err != nil {
		global.GVA_LOG.Error("provide valid data", zap.Error(err))
		response.FailWithMessage("provide valid data"+err.Error(), c)
		return
	}

	if res, err := authorityService.GetAuthorityInfo(auth); err != nil {
		global.GVA_LOG.Error("fail to get authority", zap.Error(err))
		response.FailWithMessage("fail to get authority"+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"authority": res}, c)
	}
}

// SetDataAuthority set authority
// @Tags Authority
// @Summary 设置角色资源权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAuthority true "设置角色资源权限"
// @Success 200 {object} response.Response{msg=string} "设置角色资源权限"
// @Router /authority/setDataAuthority [post]
func (a *AuthorityApi) SetDataAuthority(c *gin.Context) {
	var auth system.SysAuthority
	var err error

	err = c.ShouldBindJSON(&auth)
	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data"+err.Error(), c)
		return
	}

	if err := utils.Verify(auth, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := authorityService.SetDataAuthority(auth); err != nil {
		global.GVA_LOG.Error("fail to set authority", zap.Error(err))
		response.FailWithMessage("fail to set authority"+err.Error(), c)
	} else {
		response.OkWithMessage("success", c)
	}
}
