import service from '@/utils/request'

// @Tags DocumentPriorityLevels
// @Summary create new priority level
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentPriorityLevels true "create new priority level"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentPriorityLevels/createDocumentPriorityLevels [post]
export const createDocumentPriorityLevels = (data) => {
  return service({
    url: '/documentPriorityLevels/createDocumentPriorityLevels',
    method: 'post',
    data
  })
}

// @Tags DocumentPriorityLevels
// @Summary delete file by ID
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentPriorityLevels true "delete file by ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentPriorityLevels/deleteDocumentPriorityLevels [delete]
export const deleteDocumentPriorityLevels = (data) => {
  return service({
    url: '/documentPriorityLevels/deleteDocumentPriorityLevels',
    method: 'delete',
    data
  })
}

// @Tags DocumentPriorityLevels
// @Summary bulk delete priority level
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "bulk delete priority level"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentPriorityLevels/deleteDocumentPriorityLevels [delete]
export const deleteDocumentPriorityLevelsByIds = (data) => {
  return service({
    url: '/documentPriorityLevels/deleteDocumentPriorityLevelsByIds',
    method: 'delete',
    data
  })
}

// @Tags DocumentPriorityLevels
// @Summary update priority level
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentPriorityLevels true "update priority level"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentPriorityLevels/updateDocumentPriorityLevels [put]
export const updateDocumentPriorityLevels = (data) => {
  return service({
    url: '/documentPriorityLevels/updateDocumentPriorityLevels',
    method: 'put',
    data
  })
}

// @Tags DocumentPriorityLevels
// @Summary get priority level by ID
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DocumentPriorityLevels true "get priority level by ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentPriorityLevels/findDocumentPriorityLevels [get]
export const findDocumentPriorityLevels = (params) => {
  return service({
    url: '/documentPriorityLevels/findDocumentPriorityLevels',
    method: 'get',
    params
  })
}

// @Tags DocumentPriorityLevels
// @Summary get list of priority levels
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "get list of priority levels"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentPriorityLevels/getDocumentPriorityLevelsList [get]
export const getDocumentPriorityLevelsList = (params) => {
  return service({
    url: '/documentPriorityLevels/getDocumentPriorityLevelsList',
    method: 'get',
    params
  })
}
