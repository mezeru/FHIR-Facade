package main

import (
	"time"
)

type Patient struct {
	Id         int          `json:"id"`
	Identifier []Identifier `json:"identifier,omitempty"`
	Name       []Name       `json:"name,omitempty"`
	Gender     string       `json:"gender,omitempty"`
	BirthDate  string       `json:"birthDate,omitempty"`
	Telecom    []Telecom    `json:"telecom,omitempty"`
	Active     bool         `json:"active,omitempty"`
	Deceased   Deceased     `json:"deceased,omitempty"`
	Address    []Address    `json:"address,omitempty"`
	Contact    []Contact    `json:"contact,omitempty"`
}

type Name struct {
	Given  string `json:"given,omitempty"`
	Use    string `json:"use,omitempty"`
	Family string `json:"family,omitempty"`
}

type Telecom struct {
	Value int    `json:"value,omitempty"`
	Use   string `json:"use,omitempty"`
}

type Identifier struct {
	Value string `json:"value,omitempty"`
}

type Deceased struct {
	DeceasedBoolean  bool      `json:"deceasedBoolean,omitempty"`
	DeceasedDateTime time.Time `json:"deceasedDateTime,omitempty"`
}

type Address struct {
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
	Relationship string    `json:"relationship,omitempty"`
	Name         []Name    `json:"name,omitempty"`
	Telecom      []Telecom `json:"telecom,omitempty"`
	Address      []Address `json:"address,omitempty"`
	Gender       string    `json:"gender,omitempty"`
	Organization string    `json:"organization,omitempty"`
	Period       string    `json:"period,omitempty"`
}
