package main

import (
	"encoding/json"
	"net/http"
)

var patients []Patient

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

	patient.Id = len(patients) + 1

	patients = append(patients, patient)

	result, err := json.Marshal(patient)

	if err == nil {
		res.Write(result)
		res.WriteHeader(http.StatusOK)
	} else {

		res.Write([]byte(err.Error()))
	}

}
