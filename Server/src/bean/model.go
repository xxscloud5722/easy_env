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

type ScriptDirectory struct {
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	ParentId int    `json:"parentId,omitempty"`
}

type FileInfo struct {
	Name        string `json:"name,omitempty"`
	Last        string `json:"last,omitempty"`
	Size        string `json:"size,omitempty"`
	Description string `json:"description,omitempty"`
	Path        string `json:"path,omitempty"`
}

type DirInfo struct {
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
}
