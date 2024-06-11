package models

type Book struct {
    ID          string   `json:"id"`
    Name        string   `json:"name"`
    Description string   `json:"description"`
    Author      string   `json:"author"`
    Publisher   Publisher `json:"publisher"`
    IsFree      bool     `json:"isFree"`
    ReleaseDate string   `json:"releaseDate"`
    MinimumAge  int      `json:"minimumAge"`
}