package models

type Student struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	GPA    float32 `json:"gpa"`
	IsMale bool    `json:"isMale"`
}
