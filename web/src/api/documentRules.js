import service from '@/utils/request'

// @Tags DocumentRules
// @Summary 创建DocumentRules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentRules true "创建DocumentRules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /documentRules/createDocumentRules [post]
export const createDocumentRules = (data) => {
  return service({
    url: '/documentRules/createDocumentRules',
    method: 'post',
    data
  })
}

// @Tags DocumentRules
// @Summary 删除DocumentRules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentRules true "删除DocumentRules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /documentRules/deleteDocumentRules [delete]
export const deleteDocumentRules = (data) => {
  return service({
    url: '/documentRules/deleteDocumentRules',
    method: 'delete',
    data
  })
}

// @Tags DocumentRules
// @Summary 删除DocumentRules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DocumentRules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /documentRules/deleteDocumentRules [delete]
export const deleteDocumentRulesByIds = (data) => {
  return service({
    url: '/documentRules/deleteDocumentRulesByIds',
    method: 'delete',
    data
  })
}

// @Tags DocumentRules
// @Summary 更新DocumentRules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentRules true "更新DocumentRules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /documentRules/updateDocumentRules [put]
export const updateDocumentRules = (data) => {
  return service({
    url: '/documentRules/updateDocumentRules',
    method: 'put',
    data
  })
}

// @Tags DocumentRules
// @Summary 用id查询DocumentRules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DocumentRules true "用id查询DocumentRules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /documentRules/findDocumentRules [get]
export const findDocumentRules = (params) => {
  return service({
    url: '/documentRules/findDocumentRules',
    method: 'get',
    params
  })
}

// @Tags DocumentRules
// @Summary 分页获取DocumentRules列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取DocumentRules列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /documentRules/getDocumentRulesList [get]
export const getDocumentRulesList = (params) => {
  return service({
    url: '/documentRules/getDocumentRulesList',
    method: 'get',
    params
  })
}
