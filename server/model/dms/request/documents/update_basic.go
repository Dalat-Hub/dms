package documents

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gopkg.in/guregu/null.v4"
)

type UpdateBasic struct {
	request.GetById

	Title           string    `json:"title" form:"title"`
	Expert          string    `json:"expert" form:"expert"`
	Content         string    `json:"content" form:"content"`
	DateIssued      null.Time `json:"dateIssued" form:"dateIssued"`
	EffectDate      null.Time `json:"effectDate" form:"effectDate"`
	ExpirationDate  null.Time `json:"expirationDate" form:"expirationDate"`
	SignNumber      uint      `json:"signNumber" form:"signNumber"`
	SignYear        uint      `json:"signYear" form:"signYear"`
	SignCategory    string    `json:"signCategory" form:"signCategory"`
	SignAgency      string    `json:"signAgency" form:"signAgency"`
	SignText        string    `json:"signText" form:"signText"`
	CategoryId      uint      `json:"categoryId" form:"categoryId"`
	Agencies        []uint    `json:"agencies" form:"agencies"`
	Fields          []uint    `json:"fields" form:"fields"`
	Signers         []uint    `json:"signers" form:"signers"`
	Priority        int       `json:"priority" form:"priority"`
	Status          int       `json:"status" form:"status"`
	UpdatedBy       uint      `json:"updatedBy" form:"updatedBy"`
	BeResponsibleBy uint      `json:"beResponsibleBy" form:"beResponsibleBy"`
}
