package models

type Song struct {
	ID     string  `json:"id,omitempty"`
	Name   string  `json:"name,omitempty"`
	Labels []Label `json:"labels"`
}
