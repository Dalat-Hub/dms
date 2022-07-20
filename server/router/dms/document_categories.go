package dms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DocumentCategoriesRouter struct {
}

// InitDocumentCategoriesRouter Initialize DocumentCategories routing information
func (s *DocumentCategoriesRouter) InitDocumentCategoriesRouter(Router *gin.RouterGroup) {
	documentCategoriesRouter := Router.Group("documentCategories").Use(middleware.OperationRecord())
	documentCategoriesRouterWithoutRecord := Router.Group("documentCategories")
	var documentCategoriesApi = v1.ApiGroupApp.DmsApiGroup.DocumentCategoriesApi
	{
		documentCategoriesRouter.POST("createDocumentCategories", documentCategoriesApi.CreateDocumentCategories)
		documentCategoriesRouter.DELETE("deleteDocumentCategories", documentCategoriesApi.DeleteDocumentCategories)
		documentCategoriesRouter.DELETE("deleteDocumentCategoriesByIds", documentCategoriesApi.DeleteDocumentCategoriesByIds)
		documentCategoriesRouter.PUT("updateDocumentCategories", documentCategoriesApi.UpdateDocumentCategories)
	}
	{
		documentCategoriesRouterWithoutRecord.GET("findDocumentCategories", documentCategoriesApi.FindDocumentCategories)
		documentCategoriesRouterWithoutRecord.GET("getDocumentCategoriesList", documentCategoriesApi.GetDocumentCategoriesList)
	}
}
