package main

type Patient struct {
	Id         int        `json:"id"`
	Identifier Identifier `json:"identifier"`
	Name       Name       `json:"name"`
}

type Name []struct {
	Given  string `json:"given"`
	Use    string `json:"use"`
	Family string `json:"family"`
}

type Identifier []struct {
	Value string `json:"value"`
}
