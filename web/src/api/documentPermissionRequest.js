import service from '@/utils/request'

// @Tags DocumentPermissionRequest
// @Summary 创建DocumentPermissionRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentPermissionRequest true "创建DocumentPermissionRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /documentPermissionRequest/createDocumentPermissionRequest [post]
export const createDocumentPermissionRequest = (data) => {
  return service({
    url: '/documentPermissionRequest/createDocumentPermissionRequest',
    method: 'post',
    data
  })
}

// @Tags DocumentPermissionRequest
// @Summary approve permissioin request
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentPermissionRequest true "approve permission request"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentPermissionRequest/approvePermissionRequest [post]
export const approvePermissionRequest = (data) => {
  return service({
    url: '/documentPermissionRequest/approvePermissionRequest',
    method: 'post',
    data
  })
}

// @Tags DocumentPermissionRequest
// @Summary 删除DocumentPermissionRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentPermissionRequest true "删除DocumentPermissionRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /documentPermissionRequest/deleteDocumentPermissionRequest [delete]
export const deleteDocumentPermissionRequest = (data) => {
  return service({
    url: '/documentPermissionRequest/deleteDocumentPermissionRequest',
    method: 'delete',
    data
  })
}

// @Tags DocumentPermissionRequest
// @Summary 删除DocumentPermissionRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DocumentPermissionRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /documentPermissionRequest/deleteDocumentPermissionRequest [delete]
export const deleteDocumentPermissionRequestByIds = (data) => {
  return service({
    url: '/documentPermissionRequest/deleteDocumentPermissionRequestByIds',
    method: 'delete',
    data
  })
}

// @Tags DocumentPermissionRequest
// @Summary 更新DocumentPermissionRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentPermissionRequest true "更新DocumentPermissionRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /documentPermissionRequest/updateDocumentPermissionRequest [put]
export const updateDocumentPermissionRequest = (data) => {
  return service({
    url: '/documentPermissionRequest/updateDocumentPermissionRequest',
    method: 'put',
    data
  })
}

// @Tags DocumentPermissionRequest
// @Summary 用id查询DocumentPermissionRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DocumentPermissionRequest true "用id查询DocumentPermissionRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /documentPermissionRequest/findDocumentPermissionRequest [get]
export const findDocumentPermissionRequest = (params) => {
  return service({
    url: '/documentPermissionRequest/findDocumentPermissionRequest',
    method: 'get',
    params
  })
}

// @Tags DocumentPermissionRequest
// @Summary 分页获取DocumentPermissionRequest列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取DocumentPermissionRequest列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /documentPermissionRequest/getDocumentPermissionRequestList [get]
export const getDocumentPermissionRequestList = (params) => {
  return service({
    url: '/documentPermissionRequest/getDocumentPermissionRequestList',
    method: 'get',
    params
  })
}
