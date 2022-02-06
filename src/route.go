package main

import (
	"encoding/json"
	"net/http"
)

type Patient struct {
	Id         int        `json:"id"`
	Identifier Identifier `json:"identifier"`
	Name       Name       `json:"name"`
}

type Name struct {
	Given  string `json:"given"`
	Use    string `json:"use"`
	Family string `json:"family"`
}

type Identifier struct {
	Value string `json:"value"`
}

var patients []Patient
var id = 0

func getPatient(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-type", "application/json")

	result, err := json.Marshal(patients)

	if err == nil {
		res.Write(result)
		res.WriteHeader(http.StatusOK)
	} else {

		res.Write([]byte(err.Error()))
	}
}

func addPatient(res http.ResponseWriter, req *http.Request) {

	var patient Patient

	json.NewDecoder(req.Body).Decode(&patient)
	patient.Id = id
	id++
	patients = append(patients, patient)

	result, err := json.Marshal(patient)

	if err == nil {
		res.Write(result)
		res.WriteHeader(http.StatusOK)
	} else {

		res.Write([]byte(err.Error()))
	}

}
