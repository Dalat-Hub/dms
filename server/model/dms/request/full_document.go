package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
)

type FullDocument struct {
	dms.Documents

	ReqAgencies           []uint            `json:"agencies"`
	ReqFields             []uint            `json:"fields"`
	ReqSigners            []uint            `json:"signers"`
	ReqDocumentBaseOns    []uint            `json:"documentBaseOns"`
	ReqDocumentReferences []uint            `json:"documentReferences"`
	ReqUsersRelated       []uint            `json:"documentUsersRelated"`
	ReqAgenciesRelated    []uint            `json:"documentAgenciesRelated"`
	ReqFileInfo           *FileInfo         `json:"fileInfo"`
	ReqRuleInfo           *DocumentRuleInfo `json:"ruleInfo"`
}
