package documents

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type UpdateFile struct {
	request.GetById

	Name  string  `json:"name" form:"name"`
	Key   string  `json:"key" form:"key"`
	Url   string  `json:"url" form:"url"`
	Tag   string  `json:"tag" form:"tag"`
	Size  float64 `json:"size" form:"size"`
	Order int     `json:"order" form:"order"`
	Path  string  `json:"path" form:"path"`

	UpdatedBy       uint `json:"updatedBy" form:"updatedBy"`
	BeResponsibleBy uint `json:"beResponsibleBy" form:"beResponsibleBy"`
}
