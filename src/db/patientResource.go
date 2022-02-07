package db

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model

	Id         string       `json:"id" gorm:"unique_index"`
	Identifier []Identifier `json:"identifier,omitempty" gorm:"foreignKey:Value"`
	Name       []Name       `json:"name,omitempty" gorm:"foreignKey:Given"`
	Gender     string       `json:"gender,omitempty"`
	BirthDate  string       `json:"birthDate,omitempty"`
	Telecoms   []Telecom    `json:"telecom,omitempty" gorm:"foreignKey:Value"`
	Active     bool         `json:"active,omitempty"`
	Deceased   Deceased     `json:"deceased,omitempty" gorm:"foreignKey:DeceasedBoolean"`
	Addressess []Address    `json:"address,omitempty" gorm:"foreignKey:Use"`
	Contacts   []Contact    `json:"contact,omitempty" gorm:"foreignKey:Relationship"`
}

type Name struct {
	gorm.Model

	Given  string `json:"given,omitempty"`
	Use    string `json:"use,omitempty"`
	Family string `json:"family,omitempty"`
}

type Telecom struct {
	gorm.Model

	Value int    `json:"value,omitempty"`
	Use   string `json:"use,omitempty"`
}

type Identifier struct {
	gorm.Model

	Value string `json:"value,omitempty"`
}

type Deceased struct {
	gorm.Model

	DeceasedBoolean  bool      `json:"deceasedBoolean,omitempty"`
	DeceasedDateTime time.Time `json:"deceasedDateTime,omitempty"`
}

type Address struct {
	gorm.Model

	Use        string `json:"use,omitempty"`
	Type       string `json:"type,omitempty"`
	Text       string `json:"text,omitempty"`
	Line       string `json:"line,omitempty"`
	City       string `json:"city,omitempty"`
	District   string `json:"district,omitempty"`
	State      string `json:"state,omitempty"`
	Postalcode string `json:"postalcode,omitempty"`
	Country    string `json:"country,,omitempty"`
	Period     string `json:"period,omitempty"`
}

type Contact struct {
	gorm.Model

	Relationship string    `json:"relationship,omitempty"`
	Name         []Name    `json:"name,omitempty" gorm:"foreignKey:Given"`
	Telecom      []Telecom `json:"telecom,omitempty" gorm:"foreignKey:Value"`
	Address      []Address `json:"address,omitempty" gorm:"foreignKey:Use"`
	Gender       string    `json:"gender,omitempty"`
	Organization string    `json:"organization,omitempty"`
	Period       string    `json:"period,omitempty"`
}
