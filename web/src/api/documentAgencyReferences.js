import service from '@/utils/request'

// @Tags DocumentAgencyReferences
// @Summary 创建DocumentAgencyReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentAgencyReferences true "创建DocumentAgencyReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /documentAgencyReferences/createDocumentAgencyReferences [post]
export const createDocumentAgencyReferences = (data) => {
  return service({
    url: '/documentAgencyReferences/createDocumentAgencyReferences',
    method: 'post',
    data
  })
}

// @Tags DocumentAgencyReferences
// @Summary 删除DocumentAgencyReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentAgencyReferences true "删除DocumentAgencyReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /documentAgencyReferences/deleteDocumentAgencyReferences [delete]
export const deleteDocumentAgencyReferences = (data) => {
  return service({
    url: '/documentAgencyReferences/deleteDocumentAgencyReferences',
    method: 'delete',
    data
  })
}

// @Tags DocumentAgencyReferences
// @Summary 删除DocumentAgencyReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DocumentAgencyReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /documentAgencyReferences/deleteDocumentAgencyReferences [delete]
export const deleteDocumentAgencyReferencesByIds = (data) => {
  return service({
    url: '/documentAgencyReferences/deleteDocumentAgencyReferencesByIds',
    method: 'delete',
    data
  })
}

// @Tags DocumentAgencyReferences
// @Summary 更新DocumentAgencyReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentAgencyReferences true "更新DocumentAgencyReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /documentAgencyReferences/updateDocumentAgencyReferences [put]
export const updateDocumentAgencyReferences = (data) => {
  return service({
    url: '/documentAgencyReferences/updateDocumentAgencyReferences',
    method: 'put',
    data
  })
}

// @Tags DocumentAgencyReferences
// @Summary 用id查询DocumentAgencyReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DocumentAgencyReferences true "用id查询DocumentAgencyReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /documentAgencyReferences/findDocumentAgencyReferences [get]
export const findDocumentAgencyReferences = (params) => {
  return service({
    url: '/documentAgencyReferences/findDocumentAgencyReferences',
    method: 'get',
    params
  })
}

// @Tags DocumentAgencyReferences
// @Summary 分页获取DocumentAgencyReferences列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取DocumentAgencyReferences列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /documentAgencyReferences/getDocumentAgencyReferencesList [get]
export const getDocumentAgencyReferencesList = (params) => {
  return service({
    url: '/documentAgencyReferences/getDocumentAgencyReferencesList',
    method: 'get',
    params
  })
}
