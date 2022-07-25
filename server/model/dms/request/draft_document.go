package request

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type DraftDocument struct {
	global.GVA_MODEL
	Title        string `json:"title"`
	RelatedUsers []uint `json:"relatedUsers"`
}
