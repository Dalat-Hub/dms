import service from '@/utils/request'

// @Tags DocumentFieldReferences
// @Summary 创建DocumentFieldReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentFieldReferences true "创建DocumentFieldReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /documentFieldReferences/createDocumentFieldReferences [post]
export const createDocumentFieldReferences = (data) => {
  return service({
    url: '/documentFieldReferences/createDocumentFieldReferences',
    method: 'post',
    data
  })
}

// @Tags DocumentFieldReferences
// @Summary 删除DocumentFieldReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentFieldReferences true "删除DocumentFieldReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /documentFieldReferences/deleteDocumentFieldReferences [delete]
export const deleteDocumentFieldReferences = (data) => {
  return service({
    url: '/documentFieldReferences/deleteDocumentFieldReferences',
    method: 'delete',
    data
  })
}

// @Tags DocumentFieldReferences
// @Summary 删除DocumentFieldReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DocumentFieldReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /documentFieldReferences/deleteDocumentFieldReferences [delete]
export const deleteDocumentFieldReferencesByIds = (data) => {
  return service({
    url: '/documentFieldReferences/deleteDocumentFieldReferencesByIds',
    method: 'delete',
    data
  })
}

// @Tags DocumentFieldReferences
// @Summary 更新DocumentFieldReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentFieldReferences true "更新DocumentFieldReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /documentFieldReferences/updateDocumentFieldReferences [put]
export const updateDocumentFieldReferences = (data) => {
  return service({
    url: '/documentFieldReferences/updateDocumentFieldReferences',
    method: 'put',
    data
  })
}

// @Tags DocumentFieldReferences
// @Summary 用id查询DocumentFieldReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DocumentFieldReferences true "用id查询DocumentFieldReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /documentFieldReferences/findDocumentFieldReferences [get]
export const findDocumentFieldReferences = (params) => {
  return service({
    url: '/documentFieldReferences/findDocumentFieldReferences',
    method: 'get',
    params
  })
}

// @Tags DocumentFieldReferences
// @Summary 分页获取DocumentFieldReferences列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取DocumentFieldReferences列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /documentFieldReferences/getDocumentFieldReferencesList [get]
export const getDocumentFieldReferencesList = (params) => {
  return service({
    url: '/documentFieldReferences/getDocumentFieldReferencesList',
    method: 'get',
    params
  })
}
