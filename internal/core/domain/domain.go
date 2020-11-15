package domain

type Product struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Type        string   `json:"state"`
	Url         string   `json:"url"`
	Creator     string   `json:"creator"`
	Description string   `json:"desciption"`
	Tags        []string `json:"tags"`
}
