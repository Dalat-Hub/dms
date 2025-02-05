import service from '@/utils/request'

// @Tags Documents
// @Summary create new document
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Documents true "create new document"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/createDocuments [post]
export const createDocuments = (data) => {
  return service({
    url: '/documents/createDocuments',
    method: 'post',
    data,
  })
}

// @Tags Documents
// @Summary create new draft document
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Documents true "create new draft document"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/createDraftDocuments [post]
export const createDraftDocument = (data) => {
  return service({
    url: '/documents/createDraftDocuments',
    method: 'post',
    data,
  })
}

// @Tags Documents
// @Summary create new full document
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Documents true "create new full document"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/createFullDocuments [post]
export const createFullDocument = (data) => {
  return service({
    url: '/documents/createFullDocuments',
    method: 'post',
    data,
  })
}

// @Tags Documents
// @Summary delete document by ID
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Documents true "delete document by ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/deleteDocuments [delete]
export const deleteDocuments = (data) => {
  return service({
    url: '/documents/deleteDocuments',
    method: 'delete',
    data,
  })
}

// @Tags Documents
// @Summary bulk delete document by IDs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "bulk delete document by IDs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/deleteDocuments [delete]
export const deleteDocumentsByIds = (data) => {
  return service({
    url: '/documents/deleteDocumentsByIds',
    method: 'delete',
    data,
  })
}

// @Tags Documents
// @Summary update document
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Documents true "update document"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/updateDocuments [put]
export const updateDocuments = (data) => {
  return service({
    url: '/documents/updateDocuments',
    method: 'put',
    data,
  })
}

// @Tags Documents
// @Summary update basic document information
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body documents.UpdateBasic true "update basic document information"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/updateBasicDocuments [put]
export const updateBasicDocuments = (data) => {
  return service({
    url: '/documents/updateBasicDocuments',
    method: 'put',
    data,
  })
}

// @Tags Documents
// @Summary update related document information
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body documents.UpdateRelation true "update related document information"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/updateRelatedDocuments [put]
export const updateRelatedDocuments = (data) => {
  return service({
    url: '/documents/updateRelatedDocuments',
    method: 'put',
    data,
  })
}

// @Tags Documents
// @Summary update document attached files
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body documents.UpdateFile true "update document attached files"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/updateDocumentFiles [put]
export const updateDocumentFiles = (data) => {
  return service({
    url: '/documents/updateDocumentFiles',
    method: 'put',
    data,
  })
}

// @Tags Documents
// @Summary get document by ID
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Documents true "get document by ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/findDocuments [get]
export const findDocuments = (params) => {
  return service({
    url: '/documents/findDocuments',
    method: 'get',
    params,
  })
}

// @Tags Documents
// @Summary get list of documents
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "get list of documents"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/getDocumentsList [get]
export const getDocumentsList = (params) => {
  return service({
    url: '/documents/getDocumentsList',
    method: 'get',
    params,
  })
}

// @Tags Documents
// @Summary get list of revisions
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.GetByID true "get list of revisions"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/getDocumentRevisions [get]
export const getDocumentRevisions = (params) => {
  return service({
    url: '/documents/getDocumentRevisions',
    method: 'get',
    params,
  })
}

// @Tags Documents
// @Summary get list of attached files
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.GetByID true "get list of attached files"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/getDocumentsList [get]
export const getDocumentFiles = (params) => {
  return service({
    url: '/documents/getDocumentFiles',
    method: 'get',
    params,
  })
}

// @Tags Documents
// @Summary make duplicated version of the given document
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.GetByID true "make duplicated version of the given document"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documents/makeDuplication [post]
export const makeDuplication = (params) => {
  return service({
    url: '/documents/makeDuplication',
    method: 'post',
    params,
  })
}
