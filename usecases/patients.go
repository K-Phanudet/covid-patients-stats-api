package usecases

import (
	"github.com/K-Phanudet/covid-patients-stats-api/models"
	"github.com/K-Phanudet/covid-patients-stats-api/repositories"
)

type patientUseCase struct {
	repository repositories.PatientRepository
}

func NewPatientUseCase(repo repositories.PatientRepository) PatientUseCase {
	return &patientUseCase{
		repository: repo,
	}
}

func (p *patientUseCase) GetAllPatients() (models.Patients, error) {
	patients := models.Patients{}
	err := p.repository.GetAllPatients(&patients)
	return patients, err
}
