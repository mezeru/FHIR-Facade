package db

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model

	Id         string       `json:"id" gorm:"unique_index"`
	Identifier []Identifier `json:"identifier,omitempty"`
	Name       []Name       `json:"name,omitempty" validate:"required"`
	Gender     string       `json:"gender,omitempty"`
	BirthDate  string       `json:"birthDate,omitempty"`
	Telecom    []Telecom    `json:"telecom,omitempty"`
	Active     bool         `json:"active,omitempty"`
	Deceased   Deceased     `json:"deceased,omitempty"`
	Address    []Address    `json:"address,omitempty"`
	Contact    []Contact    `json:"contact,omitempty"`
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
	Name         []Name    `json:"name,omitempty"`
	Telecom      []Telecom `json:"telecom,omitempty"`
	Address      []Address `json:"address,omitempty"`
	Gender       string    `json:"gender,omitempty"`
	Organization string    `json:"organization,omitempty"`
	Period       string    `json:"period,omitempty"`
}
