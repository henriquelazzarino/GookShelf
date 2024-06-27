package models

type Book struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
	IsFree      bool   `json:"isFree"`
	ReleaseDate string `json:"releaseDate"`
	MinimumAge  int    `json:"minimumAge"`
}
