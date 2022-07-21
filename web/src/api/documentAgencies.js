import service from '@/utils/request'

// @Tags DocumentAgencies
// @Summary create new agency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentAgencies true "create new agency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentAgencies/createDocumentAgencies [post]
export const createDocumentAgencies = (data) => {
  return service({
    url: '/documentAgencies/createDocumentAgencies',
    method: 'post',
    data
  })
}

// @Tags DocumentAgencies
// @Summary delete agency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentAgencies true "delete agency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentAgencies/deleteDocumentAgencies [delete]
export const deleteDocumentAgencies = (data) => {
  return service({
    url: '/documentAgencies/deleteDocumentAgencies',
    method: 'delete',
    data
  })
}

// @Tags DocumentAgencies
// @Summary bulk delete agency by IDs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "bulk delete agency by IDs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentAgencies/deleteDocumentAgencies [delete]
export const deleteDocumentAgenciesByIds = (data) => {
  return service({
    url: '/documentAgencies/deleteDocumentAgenciesByIds',
    method: 'delete',
    data
  })
}

// @Tags DocumentAgencies
// @Summary update agency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentAgencies true "update agency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentAgencies/updateDocumentAgencies [put]
export const updateDocumentAgencies = (data) => {
  return service({
    url: '/documentAgencies/updateDocumentAgencies',
    method: 'put',
    data
  })
}

// @Tags DocumentAgencies
// @Summary get agency by ID
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DocumentAgencies true "get agency by ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentAgencies/findDocumentAgencies [get]
export const findDocumentAgencies = (params) => {
  return service({
    url: '/documentAgencies/findDocumentAgencies',
    method: 'get',
    params
  })
}

// @Tags DocumentAgencies
// @Summary get list of agencies
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "get list of agencies"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentAgencies/getDocumentAgenciesList [get]
export const getDocumentAgenciesList = (params) => {
  return service({
    url: '/documentAgencies/getDocumentAgenciesList',
    method: 'get',
    params
  })
}
