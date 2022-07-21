import service from '@/utils/request'

// @Tags DocumentFields
// @Summary create new field
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentFields true "create new field"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentFields/createDocumentFields [post]
export const createDocumentFields = (data) => {
  return service({
    url: '/documentFields/createDocumentFields',
    method: 'post',
    data
  })
}

// @Tags DocumentFields
// @Summary delete field by ID
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentFields true "delete field by ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentFields/deleteDocumentFields [delete]
export const deleteDocumentFields = (data) => {
  return service({
    url: '/documentFields/deleteDocumentFields',
    method: 'delete',
    data
  })
}

// @Tags DocumentFields
// @Summary bulk delete field by IDs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "bulk delete field by IDs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentFields/deleteDocumentFields [delete]
export const deleteDocumentFieldsByIds = (data) => {
  return service({
    url: '/documentFields/deleteDocumentFieldsByIds',
    method: 'delete',
    data
  })
}

// @Tags DocumentFields
// @Summary update field
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentFields true "update field"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentFields/updateDocumentFields [put]
export const updateDocumentFields = (data) => {
  return service({
    url: '/documentFields/updateDocumentFields',
    method: 'put',
    data
  })
}

// @Tags DocumentFields
// @Summary get field by ID
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DocumentFields true "get field by ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentFields/findDocumentFields [get]
export const findDocumentFields = (params) => {
  return service({
    url: '/documentFields/findDocumentFields',
    method: 'get',
    params
  })
}

// @Tags DocumentFields
// @Summary get list of fields
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "get list of fields"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentFields/getDocumentFieldsList [get]
export const getDocumentFieldsList = (params) => {
  return service({
    url: '/documentFields/getDocumentFieldsList',
    method: 'get',
    params
  })
}
