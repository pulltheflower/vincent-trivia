package models

import "gorm.io/gorm"

type Fact struct {
	gorm.Model
	Question string `json:"question" gorm:"text;not null;default:null"`
	Answer   string `json:"anwser" gorm:"text;not null;default:null"`
}
