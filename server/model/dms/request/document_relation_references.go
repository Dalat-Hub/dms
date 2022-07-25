package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
)

type DocumentRelationReferencesSearch struct {
	dms.DocumentRelationReferences
	request.PageInfo
}
