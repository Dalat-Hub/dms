package documents

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type UpdateRelation struct {
	request.GetById

	BasedDocuments   []uint `json:"basedDocuments" form:"basedDocuments"`
	RelatedDocuments []uint `json:"relatedDocuments" form:"relatedDocuments"`
	RelatedAgencies  []uint `json:"relatedAgencies" form:"relatedAgencies"`
	RelatedUsers     []uint `json:"relatedUsers" form:"relatedUsers"`
	UpdatedBy        uint   `json:"updatedBy" form:"updatedBy"`
	BeResponsibleBy  uint   `json:"beResponsibleBy" form:"beResponsibleBy"`
}
