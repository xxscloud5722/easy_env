package bean

type Pair struct {
	Id          int    `json:"id,omitempty"`
	Key         string `json:"key,omitempty"`
	Value       string `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
}
