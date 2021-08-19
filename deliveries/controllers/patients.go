package controllers

import (
	"net/http"

	"github.com/K-Phanudet/covid-patients-stats-api/models"
	"github.com/K-Phanudet/covid-patients-stats-api/usecases"
	"github.com/gin-gonic/gin"
)

type PatientController struct {
	patientUseCase usecases.PatientUseCase
}

func NewPatientController(usecase usecases.PatientUseCase) *PatientController {
	return &PatientController{
		patientUseCase: usecase,
	}
}

// Controller for API route return covid summary stats of province and age group
func (p *PatientController) GetSummaryPatientStats(ctx *gin.Context) {
	var patients models.Patients
	var err error
	response := map[string]map[string]int{}
	if patients, err = p.patientUseCase.GetAllPatients(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}
	patientGroupByProvince := groupByProvince(patients.Data)
	patientGroupByAge := groupByAge(patients.Data)
	response["Province"] = countValueInMap(patientGroupByProvince)
	response["AgeGroup"] = countValueInMap(patientGroupByAge)
	ctx.JSON(http.StatusOK, response)
}

// groupByProvince : Group patients by province return map key is province , value is slice of patient infomation
func groupByProvince(patients []models.PatientInfo) map[string][]models.PatientInfo {
	patientGroup := map[string][]models.PatientInfo{}
	for _, patient := range patients {
		patientGroup[patient.Province] = append(patientGroup[patient.Province], patient)
	}
	return patientGroup
}

// groupByProvince : Group patients by age range return map key is age range , value is slice of patient infomation
func groupByAge(patients []models.PatientInfo) map[string][]models.PatientInfo {
	patientGroup := map[string][]models.PatientInfo{
		"61+":   {},
		"31-60": {},
		"0-30":  {},
		"N/A":   {},
	}
	for _, patient := range patients {
		if patient.Age >= 61 {
			patientGroup["61+"] = append(patientGroup["61+"], patient)
			continue
		}
		if patient.Age >= 31 {
			patientGroup["31-60"] = append(patientGroup["31-60"], patient)
			continue
		}
		if patient.Age >= 0 {
			patientGroup["0-30"] = append(patientGroup["0-30"], patient)
			continue
		}
		patientGroup["N/A"] = append(patientGroup["N/A"], patient)
	}
	return patientGroup
}

// countValueInMap : count number of value in each key then return new map that value is number
func countValueInMap(object map[string][]models.PatientInfo) map[string]int {
	result := map[string]int{}
	for key, value := range object {
		result[key] = len(value)
	}
	return result
}
