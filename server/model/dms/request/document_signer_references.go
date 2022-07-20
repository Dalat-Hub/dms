package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
)

type DocumentSignerReferencesSearch struct {
	dms.DocumentSignerReferences
	request.PageInfo
}
