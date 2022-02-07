package main

import (
	"encoding/json"
	"net/http"

	uuid "github.com/google/uuid"
	"github.com/mezeru/FHIR-Facade/src/db"
)

func getPatient(res http.ResponseWriter, req *http.Request) {

	var patient []db.Patient

	res.Header().Set("Content-type", "application/json")

	db.PgDataBase.Select("id", "gender", "birthDate").Find(&patient)

	result, err := json.Marshal(patient)

	if err == nil {
		res.Write(result)
		res.WriteHeader(http.StatusOK)
	} else {

		res.Write([]byte(err.Error()))
	}
}

func addPatient(res http.ResponseWriter, req *http.Request) {

	var patient db.Patient

	json.NewDecoder(req.Body).Decode(&patient)

	output, err := json.Marshal(patient)

	u := uuid.New()

	patient.ID = u.String()

	if err == nil {
		res.Write(output)
		res.WriteHeader(http.StatusOK)
	} else {

		res.Write([]byte(err.Error()))
	}

	db.PgDataBase.Create(&patient)

}
