package usecases

import (
	"github.com/K-Phanudet/covid-patients-stats-api/models"
)

type PatientUseCase interface {
	GetAllPatients() (models.Patients, error)
}
