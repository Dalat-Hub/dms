package request

type DocumentUsersDistribution struct {
	ID         uint   `json:"id" form:"id"`
	DocumentId uint   `json:"documentId" form:"documentId"`
	UserIds    []uint `json:"userIds" form:"userIds"`
}
