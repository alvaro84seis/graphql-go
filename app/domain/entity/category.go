package entity

type Category struct {
	ID           int       `json:"id"`
	CreatedAt    string    `json:"created_at"`
	UpdatedAt    string    `json:"updated_at"`
	Code         string    `json:"code"`
	DefaultLabel string    `json:"default_label"`
	Path         string    `json:"path"`
	Level        int       `json:"level"`
	ParentCode   string    `json:"parent_code"`
	LabelPath    string    `json:"label_path"`
	Parent       *Category `json:"parent,omitempty"`
}
