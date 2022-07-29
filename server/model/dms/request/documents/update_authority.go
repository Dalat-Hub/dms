package documents

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type UpdateAuthority struct {
	request.GetById

	View          string `json:"view" form:"view"`
	Download      string `json:"download" form:"download"`
	Edit          string `json:"edit" form:"edit"`
	Owner         string `json:"owner" form:"owner"`
	ViewUsers     []uint `json:"viewUsers" form:"viewUsers"`
	ViewRoles     []uint `json:"viewRoles" form:"viewRoles"`
	DownloadUsers []uint `json:"downloadUsers" form:"downloadUsers"`
	DownloadRoles []uint `json:"downloadRoles" form:"downloadRoles"`
	EditUsers     []uint `json:"editUsers" form:"editUsers"`
	EditRoles     []uint `json:"editRoles" form:"editRoles"`
	OwnerUsers    []uint `json:"ownerUsers" form:"ownerUsers"`
	OwnerRoles    []uint `json:"ownerRoles" form:"ownerRoles"`
}
