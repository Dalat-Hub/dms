package request

type DocumentRuleInfo struct {
	View          string `json:"view"`
	Download      string `json:"download"`
	Edit          string `json:"edit"`
	Owner         string `json:"owner"`
	ViewUsers     []uint `json:"viewUsers"`
	ViewRoles     []uint `json:"viewRoles"`
	DownloadUsers []uint `json:"downloadUsers"`
	DownloadRoles []uint `json:"downloadRoles"`
	EditUsers     []uint `json:"editUsers"`
	EditRoles     []uint `json:"editRoles"`
	OwnerUsers    []uint `json:"ownerUsers"`
	OwnerRoles    []uint `json:"ownerRoles"`
}
