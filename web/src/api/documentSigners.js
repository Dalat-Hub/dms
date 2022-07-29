import service from '@/utils/request'

// @Tags DocumentSigners
// @Summary 创建DocumentSigners
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentSigners true "创建DocumentSigners"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /documentSigners/createDocumentSigners [post]
export const createDocumentSigners = (data) => {
  return service({
    url: '/documentSigners/createDocumentSigners',
    method: 'post',
    data
  })
}

// @Tags DocumentSigners
// @Summary 删除DocumentSigners
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentSigners true "删除DocumentSigners"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /documentSigners/deleteDocumentSigners [delete]
export const deleteDocumentSigners = (data) => {
  return service({
    url: '/documentSigners/deleteDocumentSigners',
    method: 'delete',
    data
  })
}

// @Tags DocumentSigners
// @Summary 删除DocumentSigners
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DocumentSigners"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /documentSigners/deleteDocumentSigners [delete]
export const deleteDocumentSignersByIds = (data) => {
  return service({
    url: '/documentSigners/deleteDocumentSignersByIds',
    method: 'delete',
    data
  })
}

// @Tags DocumentSigners
// @Summary 更新DocumentSigners
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentSigners true "更新DocumentSigners"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /documentSigners/updateDocumentSigners [put]
export const updateDocumentSigners = (data) => {
  return service({
    url: '/documentSigners/updateDocumentSigners',
    method: 'put',
    data
  })
}

// @Tags DocumentSigners
// @Summary 用id查询DocumentSigners
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DocumentSigners true "用id查询DocumentSigners"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /documentSigners/findDocumentSigners [get]
export const findDocumentSigners = (params) => {
  return service({
    url: '/documentSigners/findDocumentSigners',
    method: 'get',
    params
  })
}

// @Tags DocumentSigners
// @Summary 分页获取DocumentSigners列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取DocumentSigners列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /documentSigners/getDocumentSignersList [get]
export const getDocumentSignersList = (params) => {
  return service({
    url: '/documentSigners/getDocumentSignersList',
    method: 'get',
    params
  })
}
