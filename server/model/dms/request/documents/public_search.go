package documents

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	request2 "github.com/flipped-aurora/gin-vue-admin/server/model/dms/request"
	"gopkg.in/guregu/null.v4"
)

type PublicSearch struct {
	request.PageInfo

	Category       uint      `json:"category" form:"category"`
	Field          uint      `json:"field" form:"field"`
	Agency         uint      `json:"agency" form:"agency"`
	FromDate       null.Time `json:"fromDate" form:"fromDate"`
	ToDate         null.Time `json:"toDate" form:"toDate"`
	StillValid     null.Bool `json:"stillValid" form:"stillValid"`
	OrderBy        string    `json:"orderBy" form:"orderBy"`
	OrderDirection string    `json:"orderDirection" form:"orderDirection"`

	request2.PreloadDocument
}
