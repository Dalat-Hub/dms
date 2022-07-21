import service from '@/utils/request'

// @Tags DocumentRelationReferences
// @Summary 创建DocumentRelationReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentRelationReferences true "创建DocumentRelationReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /documentRelationReferences/createDocumentRelationReferences [post]
export const createDocumentRelationReferences = (data) => {
  return service({
    url: '/documentRelationReferences/createDocumentRelationReferences',
    method: 'post',
    data
  })
}

// @Tags DocumentRelationReferences
// @Summary 删除DocumentRelationReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentRelationReferences true "删除DocumentRelationReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /documentRelationReferences/deleteDocumentRelationReferences [delete]
export const deleteDocumentRelationReferences = (data) => {
  return service({
    url: '/documentRelationReferences/deleteDocumentRelationReferences',
    method: 'delete',
    data
  })
}

// @Tags DocumentRelationReferences
// @Summary 删除DocumentRelationReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DocumentRelationReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /documentRelationReferences/deleteDocumentRelationReferences [delete]
export const deleteDocumentRelationReferencesByIds = (data) => {
  return service({
    url: '/documentRelationReferences/deleteDocumentRelationReferencesByIds',
    method: 'delete',
    data
  })
}

// @Tags DocumentRelationReferences
// @Summary 更新DocumentRelationReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentRelationReferences true "更新DocumentRelationReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /documentRelationReferences/updateDocumentRelationReferences [put]
export const updateDocumentRelationReferences = (data) => {
  return service({
    url: '/documentRelationReferences/updateDocumentRelationReferences',
    method: 'put',
    data
  })
}

// @Tags DocumentRelationReferences
// @Summary 用id查询DocumentRelationReferences
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DocumentRelationReferences true "用id查询DocumentRelationReferences"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /documentRelationReferences/findDocumentRelationReferences [get]
export const findDocumentRelationReferences = (params) => {
  return service({
    url: '/documentRelationReferences/findDocumentRelationReferences',
    method: 'get',
    params
  })
}

// @Tags DocumentRelationReferences
// @Summary 分页获取DocumentRelationReferences列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取DocumentRelationReferences列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /documentRelationReferences/getDocumentRelationReferencesList [get]
export const getDocumentRelationReferencesList = (params) => {
  return service({
    url: '/documentRelationReferences/getDocumentRelationReferencesList',
    method: 'get',
    params
  })
}
