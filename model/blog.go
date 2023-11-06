package model

type Blog struct {
	Name string `json:"name" gorm:"not null;column:name;size:255"`
	Author string `json:"author" gorm:"not null;column:author;size:255"`
}