package db

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model

	Id          string       `json:"id" gorm:"unique_index"`
	Identifiers []{ Value string `json:"value,omitempty"`} `json:"identifier,omitempty"`
	Names       []Name       `json:"name,omitempty" gorm:"foreignKey:Use;references:Use"`
	Gender      string       `json:"gender,omitempty"`
	BirthDate   string       `json:"birthDate,omitempty"`
	Telecoms    []Telecom    `json:"telecom,omitempty" gorm:"foreignKey:Value;references:Use"`
	Active      bool         `json:"active,omitempty"`
	Deceased    Deceased     `json:"deceased,omitempty" gorm:"foreignKey:DeceasedBoolean ;references:deceasedDateTime"`
	Addressess  []Address    `json:"address,omitempty" gorm:"foreignKey:Type;references:Use"`
	Contacts    []Contact    `json:"contact,omitempty" gorm:"foreignKey:Organisation;references:Relationship"`
}

type Name struct {
	gorm.Model

	Use    string `json:"use,omitempty"`
	Given  string `json:"given,omitempty"`
	Family string `json:"family,omitempty"`
}

type Telecom struct {
	gorm.Model

	Use   string `json:"use,omitempty"`
	Value int    `json:"value,omitempty"`
}

type Deceased struct {
	gorm.Model

	DeceasedDateTime time.Time `json:"deceasedDateTime,omitempty"`
	DeceasedBoolean  bool      `json:"deceasedBoolean,omitempty"`
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
	Names        []Name    `json:"name,omitempty" gorm:"foreignKey:Given"`
	Telecoms     []Telecom `json:"telecom,omitempty" gorm:"foreignKey:Value"`
	Addressess   []Address `json:"address,omitempty" gorm:"foreignKey:Type"`
	Gender       string    `json:"gender,omitempty"`
	Organization string    `json:"organization,omitempty"`
	Period       string    `json:"period,omitempty"`
}
