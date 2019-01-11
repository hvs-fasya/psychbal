package models

//Tab tab structure
type Tab struct {
	ID              int64    `json:"id"`
	Title           string   `json:"title"`
	UserTypeVisible []string `json:"user_type_visible"`
	Index           int64    `json:"index"`
	Enabled         bool     `json:"enabled"`
	*TabType        `json:"tab_type"`
}

//TabType tab_type structure
type TabType struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
}
