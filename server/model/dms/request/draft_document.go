package request

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type DraftDocument struct {
	global.GVA_MODEL
	Title        string `json:"title" form:"title"`
	RelatedUsers []uint `json:"relatedUsers" form:"relatedUsers"`
	Category     uint   `json:"category" form:"category"`
	Agency       uint   `json:"agency" form:"agency"`
	SignNumber   uint   `json:"signNumber" form:"signNumber"`
	SignYear     uint   `json:"signYear" form:"signYear"`
	SignText     string `json:"signText" form:"signText"`
	CategoryText string `json:"categoryText" form:"categoryText"`
	AgencyText   string `json:"agencyText" form:"agencyText"`
}
