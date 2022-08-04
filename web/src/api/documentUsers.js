import service from '@/utils/request'

// @Tags DocumentUsers
// @Summary 创建DocumentUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentUsers true "创建DocumentUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /documentUsers/createDocumentUsers [post]
export const createDocumentUsers = (data) => {
  return service({
    url: '/documentUsers/createDocumentUsers',
    method: 'post',
    data
  })
}

// @Tags DocumentUsers
// @Summary 删除DocumentUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentUsers true "删除DocumentUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /documentUsers/deleteDocumentUsers [delete]
export const deleteDocumentUsers = (data) => {
  return service({
    url: '/documentUsers/deleteDocumentUsers',
    method: 'delete',
    data
  })
}

// @Tags DocumentUsers
// @Summary 删除DocumentUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DocumentUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /documentUsers/deleteDocumentUsers [delete]
export const deleteDocumentUsersByIds = (data) => {
  return service({
    url: '/documentUsers/deleteDocumentUsersByIds',
    method: 'delete',
    data
  })
}

// @Tags DocumentUsers
// @Summary 更新DocumentUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentUsers true "更新DocumentUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /documentUsers/updateDocumentUsers [put]
export const updateDocumentUsers = (data) => {
  return service({
    url: '/documentUsers/updateDocumentUsers',
    method: 'put',
    data
  })
}

// @Tags DocumentUsers
// @Summary 用id查询DocumentUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DocumentUsers true "用id查询DocumentUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /documentUsers/findDocumentUsers [get]
export const findDocumentUsers = (params) => {
  return service({
    url: '/documentUsers/findDocumentUsers',
    method: 'get',
    params
  })
}

// @Tags DocumentUsers
// @Summary 分页获取DocumentUsers列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取DocumentUsers列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /documentUsers/getDocumentUsersList [get]
export const getDocumentUsersList = (params) => {
  return service({
    url: '/documentUsers/getDocumentUsersList',
    method: 'get',
    params
  })
}

// @Tags DocumentUsers
// @Summary distribute tasks
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "distribute tasks"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentUsers/distributeTasks [get]
export const distributeTasks = (data) => {
  return service({
    url: '/documentUsers/distributeTasks',
    method: 'post',
    data
  })
}
