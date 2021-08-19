package repositories

import (
	"github.com/K-Phanudet/covid-patients-stats-api/models"
)

type PatientRepository interface {
	GetAllPatients(patients *models.Patients) error
}
