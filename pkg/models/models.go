package models

// Book model Struct
type Book struct {
	ID     int    `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author Author `json:"author"`
	// Author *Author `json:"author"`
}

// Author model Struct
type Author struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
