import service from '@/utils/request'

// @Tags DocumentCategories
// @Summary create new category
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentCategories true "create new category"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentCategories/createDocumentCategories [post]
export const createDocumentCategories = (data) => {
  return service({
    url: '/documentCategories/createDocumentCategories',
    method: 'post',
    data
  })
}

// @Tags DocumentCategories
// @Summary delete category by ID
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentCategories true "delete category by ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentCategories/deleteDocumentCategories [delete]
export const deleteDocumentCategories = (data) => {
  return service({
    url: '/documentCategories/deleteDocumentCategories',
    method: 'delete',
    data
  })
}

// @Tags DocumentCategories
// @Summary bulk delete category by IDs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "bulk delete category by IDs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentCategories/deleteDocumentCategories [delete]
export const deleteDocumentCategoriesByIds = (data) => {
  return service({
    url: '/documentCategories/deleteDocumentCategoriesByIds',
    method: 'delete',
    data
  })
}

// @Tags DocumentCategories
// @Summary update category
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocumentCategories true "update category"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentCategories/updateDocumentCategories [put]
export const updateDocumentCategories = (data) => {
  return service({
    url: '/documentCategories/updateDocumentCategories',
    method: 'put',
    data
  })
}

// @Tags DocumentCategories
// @Summary get category by ID
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DocumentCategories true "get category by ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentCategories/findDocumentCategories [get]
export const findDocumentCategories = (params) => {
  return service({
    url: '/documentCategories/findDocumentCategories',
    method: 'get',
    params
  })
}

// @Tags DocumentCategories
// @Summary get list of categories
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "get list of categories"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /documentCategories/getDocumentCategoriesList [get]
export const getDocumentCategoriesList = (params) => {
  return service({
    url: '/documentCategories/getDocumentCategoriesList',
    method: 'get',
    params
  })
}
