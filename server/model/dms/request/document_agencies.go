package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
)

type DocumentAgenciesSearch struct {
	dms.DocumentAgencies
	request.PageInfo
}
