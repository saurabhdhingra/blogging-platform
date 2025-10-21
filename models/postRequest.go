package models


type PostRequest struct {
	Title		string		`json:"title"`
	Content		string		`json:"content"`
	Category	string		`json:"category"`
	Tags		[]string	`json:"tags"`
}