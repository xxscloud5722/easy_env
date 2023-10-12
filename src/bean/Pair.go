package bean

type Pair struct {
	Id          int    `json:"id,omitempty"`
	Key         string `json:"key,omitempty"`
	Value       string `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
}

type Script struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Path        string `json:"path,omitempty"`
	Value       string `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
}
