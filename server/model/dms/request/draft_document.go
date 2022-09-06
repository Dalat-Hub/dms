package request

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type DraftDocument struct {
	global.GVA_MODEL
	Title        string `json:"title" form:"title"`
	RelatedUsers []uint `json:"relatedUsers" form:"relatedUsers"`
	Agencies     []uint `json:"agencies" form:"agencies"`
	Category     uint   `json:"category" form:"category"`
	SignNumber   uint   `json:"signNumber" form:"signNumber"`
	SignYear     uint   `json:"signYear" form:"signYear"`
	SignText     string `json:"signText" form:"signText"`
	CategoryText string `json:"categoryText" form:"categoryText"`
	AgencyText   string `json:"agencyText" form:"agencyText"`
}
