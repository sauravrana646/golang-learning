package model

type Thread struct {
	Id          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description,omitempty"`
	URL         string   `json:"url"`
	Tags        []string `json:"tags"`
}
