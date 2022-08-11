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

type AuthorityMenuApi struct{}

// GetMenu
// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {object} response.Response{data=systemRes.SysMenusResponse,msg=string} "获取用户动态路由,返回包括系统菜单详情列表"
// @Router /menu/getMenu [post]
func (a *AuthorityMenuApi) GetMenu(c *gin.Context) {
	if menus, err := menuService.GetMenuTree(utils.GetUserAuthorityId(c)); err != nil {
		global.GVA_LOG.Error("fail to get menu", zap.Error(err))
		response.FailWithMessage("fail to get menu", c)
	} else {
		if menus == nil {
			menus = []system.SysMenu{}
		}
		response.OkWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "success", c)
	}
}

// GetBaseMenuTree
// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {object} response.Response{data=systemRes.SysBaseMenusResponse,msg=string} "获取用户动态路由,返回包括系统菜单列表"
// @Router /menu/getBaseMenuTree [post]
func (a *AuthorityMenuApi) GetBaseMenuTree(c *gin.Context) {
	if menus, err := menuService.GetBaseMenuTree(); err != nil {
		global.GVA_LOG.Error("fail to get base menu tree", zap.Error(err))
		response.FailWithMessage("fail to get base menu tree", c)
	} else {
		response.OkWithDetailed(systemRes.SysBaseMenusResponse{Menus: menus}, "success", c)
	}
}

// AddMenuAuthority
// @Tags AuthorityMenu
// @Summary 增加menu和角色关联关系
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.AddMenuAuthorityInfo true "角色ID"
// @Success 200 {object} response.Response{msg=string} "增加menu和角色关联关系"
// @Router /menu/addMenuAuthority [post]
func (a *AuthorityMenuApi) AddMenuAuthority(c *gin.Context) {
	var authorityMenu systemReq.AddMenuAuthorityInfo
	var err error

	err = c.ShouldBindJSON(&authorityMenu)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := utils.Verify(authorityMenu, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := menuService.AddMenuAuthority(authorityMenu.Menus, authorityMenu.AuthorityId); err != nil {
		global.GVA_LOG.Error("fail to add menu authority", zap.Error(err))
		response.FailWithMessage("fail to add menu authority", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// GetMenuAuthority
// @Tags AuthorityMenu
// @Summary 获取指定角色menu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetAuthorityId true "角色ID"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取指定角色menu"
// @Router /menu/getMenuAuthority [post]
func (a *AuthorityMenuApi) GetMenuAuthority(c *gin.Context) {
	var param request.GetAuthorityId
	var err error

	err = c.ShouldBindJSON(&param)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := utils.Verify(param, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if menus, err := menuService.GetMenuAuthority(&param); err != nil {
		global.GVA_LOG.Error("fail to get menu authority", zap.Error(err))
		response.FailWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "fail to get menu authority", c)
	} else {
		response.OkWithDetailed(gin.H{"menus": menus}, "success", c)
	}
}

// AddBaseMenu
// @Tags Menu
// @Summary 新增菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysBaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {object} response.Response{msg=string} "新增菜单"
// @Router /menu/addBaseMenu [post]
func (a *AuthorityMenuApi) AddBaseMenu(c *gin.Context) {
	var menu system.SysBaseMenu
	var err error

	err = c.ShouldBindJSON(&menu)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := utils.Verify(menu, utils.MenuVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(menu.Meta, utils.MenuMetaVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := menuService.AddBaseMenu(menu); err != nil {
		global.GVA_LOG.Error("fail to add base menu", zap.Error(err))

		response.FailWithMessage("fail to add base menu", c)
	} else {
		response.OkWithMessage("fail to add base menu", c)
	}
}

// DeleteBaseMenu
// @Tags Menu
// @Summary 删除菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "菜单id"
// @Success 200 {object} response.Response{msg=string} "删除菜单"
// @Router /menu/deleteBaseMenu [post]
func (a *AuthorityMenuApi) DeleteBaseMenu(c *gin.Context) {
	var menu request.GetById
	var err error

	err = c.ShouldBindJSON(&menu)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := utils.Verify(menu, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := baseMenuService.DeleteBaseMenu(menu.ID); err != nil {
		global.GVA_LOG.Error("fail to delete base menu", zap.Error(err))
		response.FailWithMessage("fail to delete base menu", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// UpdateBaseMenu
// @Tags Menu
// @Summary 更新菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysBaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {object} response.Response{msg=string} "更新菜单"
// @Router /menu/updateBaseMenu [post]
func (a *AuthorityMenuApi) UpdateBaseMenu(c *gin.Context) {
	var menu system.SysBaseMenu
	var err error

	err = c.ShouldBindJSON(&menu)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := utils.Verify(menu, utils.MenuVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(menu.Meta, utils.MenuMetaVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := baseMenuService.UpdateBaseMenu(menu); err != nil {
		global.GVA_LOG.Error("fail to update base menu", zap.Error(err))
		response.FailWithMessage("fail to update base menu", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// GetBaseMenuById
// @Tags Menu
// @Summary 根据id获取菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "菜单id"
// @Success 200 {object} response.Response{data=systemRes.SysBaseMenuResponse,msg=string} "根据id获取菜单,返回包括系统菜单列表"
// @Router /menu/getBaseMenuById [post]
func (a *AuthorityMenuApi) GetBaseMenuById(c *gin.Context) {
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
	if menu, err := baseMenuService.GetBaseMenuById(idInfo.ID); err != nil {
		global.GVA_LOG.Error("fail to get base menu by id", zap.Error(err))
		response.FailWithMessage("fail to get base menu by id", c)
	} else {
		response.OkWithDetailed(systemRes.SysBaseMenuResponse{Menu: menu}, "success", c)
	}
}

// GetMenuList
// @Tags Menu
// @Summary 分页获取基础menu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取基础menu列表,返回包括列表,总数,页码,每页数量"
// @Router /menu/getMenuList [post]
func (a *AuthorityMenuApi) GetMenuList(c *gin.Context) {
	var pageInfo request.PageInfo
	var err error

	err = c.ShouldBindJSON(&pageInfo)

	if err != nil {
		global.GVA_LOG.Error("please provide valid data", zap.Error(err))
		response.FailWithMessage("please provide valid data", c)
		return
	}

	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if menuList, total, err := menuService.GetInfoList(); err != nil {
		global.GVA_LOG.Error("fail to get menu list", zap.Error(err))
		response.FailWithMessage("fail to get menu list", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     menuList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "success", c)
	}
}
