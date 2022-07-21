import service from '@/utils/request'

// @Tags DocumentFiles
// @Summary create new file
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentFiles true "create new file"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentFiles/createDocumentFiles [post]
export const createDocumentFiles = (data) => {
  return service({
    url: '/documentFiles/createDocumentFiles',
    method: 'post',
    data
  })
}

// @Tags DocumentFiles
// @Summary delete file by ID
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentFiles true "delete file by ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentFiles/deleteDocumentFiles [delete]
export const deleteDocumentFiles = (data) => {
  return service({
    url: '/documentFiles/deleteDocumentFiles',
    method: 'delete',
    data
  })
}

// @Tags DocumentFiles
// @Summary bulk delete file by IDs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "bulk delete file by IDs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentFiles/deleteDocumentFiles [delete]
export const deleteDocumentFilesByIds = (data) => {
  return service({
    url: '/documentFiles/deleteDocumentFilesByIds',
    method: 'delete',
    data
  })
}

// @Tags DocumentFiles
// @Summary update file
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentFiles true "update file"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentFiles/updateDocumentFiles [put]
export const updateDocumentFiles = (data) => {
  return service({
    url: '/documentFiles/updateDocumentFiles',
    method: 'put',
    data
  })
}

// @Tags DocumentFiles
// @Summary get file by ID
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DocumentFiles true "get file by ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentFiles/findDocumentFiles [get]
export const findDocumentFiles = (params) => {
  return service({
    url: '/documentFiles/findDocumentFiles',
    method: 'get',
    params
  })
}

// @Tags DocumentFiles
// @Summary get list of files
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "get list of files"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentFiles/getDocumentFilesList [get]
export const getDocumentFilesList = (params) => {
  return service({
    url: '/documentFiles/getDocumentFilesList',
    method: 'get',
    params
  })
}
