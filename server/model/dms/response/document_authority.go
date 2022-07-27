package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/system"

type DocumentAuthority struct {
	PublicToView         bool                  `json:"publicToView"`
	PublicToDownload     bool                  `json:"publicToDownload"`
	OnlyAdminCanView     bool                  `json:"onlyAdminCanView"`
	OnlyAdminCanDownload bool                  `json:"onlyAdminCanDownload"`
	ViewLimitUsers       []system.SysUser      `json:"viewLimitUsers"`
	ViewLimitRoles       []system.SysAuthority `json:"viewLimitRoles"`
	DownloadLimitUsers   []system.SysUser      `json:"downloadLimitUsers"`
	DownloadLimitRoles   []system.SysAuthority `json:"downloadLimitRoles"`
	OnlyMeCanUpdate      bool                  `json:"onlyMeCanUpdate"`
	UpdateLimitUsers     []system.SysUser      `json:"updateLimitUsers"`
	UpdateLimitRoles     []system.SysAuthority `json:"updateLimitRoles"`
	OnlyMeBeOwner        bool                  `json:"onlyMeBeOwner"`
	OwnerLimitUsers      []system.SysUser      `json:"ownerLimitUsers"`
	OwnerLimitRoles      []system.SysAuthority `json:"ownerLimitRoles"`
}
