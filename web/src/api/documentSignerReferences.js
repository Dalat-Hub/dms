import service from '@/utils/request'

// @Tags DocumentSignerReferences
// @Summary 创建DocumentSignerReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentSignerReferences true "创建DocumentSignerReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /documentSignerReferences/createDocumentSignerReferences [post]
export const createDocumentSignerReferences = (data) => {
  return service({
    url: '/documentSignerReferences/createDocumentSignerReferences',
    method: 'post',
    data
  })
}

// @Tags DocumentSignerReferences
// @Summary 删除DocumentSignerReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentSignerReferences true "删除DocumentSignerReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /documentSignerReferences/deleteDocumentSignerReferences [delete]
export const deleteDocumentSignerReferences = (data) => {
  return service({
    url: '/documentSignerReferences/deleteDocumentSignerReferences',
    method: 'delete',
    data
  })
}

// @Tags DocumentSignerReferences
// @Summary 删除DocumentSignerReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DocumentSignerReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /documentSignerReferences/deleteDocumentSignerReferences [delete]
export const deleteDocumentSignerReferencesByIds = (data) => {
  return service({
    url: '/documentSignerReferences/deleteDocumentSignerReferencesByIds',
    method: 'delete',
    data
  })
}

// @Tags DocumentSignerReferences
// @Summary 更新DocumentSignerReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentSignerReferences true "更新DocumentSignerReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /documentSignerReferences/updateDocumentSignerReferences [put]
export const updateDocumentSignerReferences = (data) => {
  return service({
    url: '/documentSignerReferences/updateDocumentSignerReferences',
    method: 'put',
    data
  })
}

// @Tags DocumentSignerReferences
// @Summary 用id查询DocumentSignerReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DocumentSignerReferences true "用id查询DocumentSignerReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /documentSignerReferences/findDocumentSignerReferences [get]
export const findDocumentSignerReferences = (params) => {
  return service({
    url: '/documentSignerReferences/findDocumentSignerReferences',
    method: 'get',
    params
  })
}

// @Tags DocumentSignerReferences
// @Summary 分页获取DocumentSignerReferences列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取DocumentSignerReferences列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /documentSignerReferences/getDocumentSignerReferencesList [get]
export const getDocumentSignerReferencesList = (params) => {
  return service({
    url: '/documentSignerReferences/getDocumentSignerReferencesList',
    method: 'get',
    params
  })
}
