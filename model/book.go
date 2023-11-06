package model

type Book struct {
	Name string `json:"name" gorm:"not null;column:name;size:255"`
	Author string `json:"author" gorm:"not null;column:author;size:255"`
	Publisher string `json:"pubisher" gorm:"not null;column:pubisher;size:255"`
}