package request

type PreloadDocument struct {
	PreloadCategory        int `json:"preloadCategory" form:"preloadCategory"`
	PreloadAgency          int `json:"preloadAgency" form:"preloadAgency"`
	PreloadFields          int `json:"preloadFields" form:"preloadFields"`
	PreloadSigners         int `json:"preloadSigners" form:"preloadSigners"`
	PreloadCreatedBy       int `json:"preloadCreatedBy" form:"preloadCreatedBy"`
	PreloadUpdatedBy       int `json:"preloadUpdatedBy" form:"preloadUpdatedBy"`
	PreloadBeResponsibleBy int `json:"preloadBeResponsibleBy" form:"preloadBeResponsibleBy"`
	PreloadBasedDocs       int `json:"preloadBasedDocs" form:"preloadBasedDocs"`
	PreloadRelatedDocs     int `json:"preloadRelatedDocs" form:"preloadRelatedDocs"`
	PreloadRelatedUsers    int `json:"preloadRelatedUsers" form:"preloadRelatedUsers"`
	PreloadRelatedAgencies int `json:"preloadRelatedAgencies" form:"preloadRelatedAgencies"`
	PreloadAuthority       int `json:"preloadAuthority" form:"preloadAuthority"`
}
