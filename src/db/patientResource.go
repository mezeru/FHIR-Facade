package db

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model

	ID         string       `json:"id,omitempty" gorm:"primaryKey"`
	Identifier []Identifier `json:"identifier,omitempty" gorm:"foreignKey:PatientId" `
	Name       []Name       `json:"name,omitempty" gorm:"foreignKey:PatientId" `
	Gender     string       `json:"gender,omitempty"`
	BirthDate  string       `json:"birthDate,omitempty"`
	Telecom    []Telecom    `json:"telecom,omitempty" gorm:"foreignKey:PatientId"`
	Active     bool         `json:"active,omitempty"`
	// Deceased   Deceased     `json:"deceased,omitempty" gorm:"foreignKey:PatientId"`
	// Address    []Address    `json:"address,omitempty" gorm:"foreignKey:PatientId"`
	// Contacts   []Contact    `json:"contact,omitempty" gorm:"foreignKey:PatientId"`
}

type Name struct {
	gorm.Model

	Use       string `json:"use,omitempty"`
	Given     string `json:"given,omitempty"`
	Family    string `json:"family,omitempty"`
	PatientId string `json:"patientId,omitempty"`
}

type Identifier struct {
	gorm.Model

	Value     string `json:"value,omitempty"`
	PatientId string `json:"patientId,omitempty"`
}

type Telecom struct {
	gorm.Model

	Use       string `json:"use,omitempty"`
	Value     int    `json:"value,omitempty"`
	PatientId string `json:"patientId,omitempty"`
}

type Deceased struct {
	gorm.Model

	DeceasedDateTime time.Time `json:"deceasedDateTime,omitempty"`
	DeceasedBoolean  bool      `json:"deceasedBoolean,omitempty"`
	PatientId        string    `json:"patientId,omitempty"`
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
	PatientId  string `json:"patientId,omitempty"`
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
	PatientId    string    `json:"patientId,omitempty"`
}
