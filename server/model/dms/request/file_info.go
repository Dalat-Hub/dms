package request

type FileInfo struct {
	Name  string  `json:"name"`
	Key   string  `json:"key"`
	Url   string  `json:"url"`
	Tag   string  `json:"tag"`
	Size  float64 `json:"size"`
	Order int     `json:"order"`
	Path  string  `json:"path"`
}
