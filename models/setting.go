package models

type Setting struct {
	ID    string `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type NewSetting struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type UpdateSetting struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
