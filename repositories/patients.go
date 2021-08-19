package repositories

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/K-Phanudet/covid-patients-stats-api/models"
)

type patientRepository struct {
	conn *http.Client
}

func NewCovidRepository(conn *http.Client) PatientRepository {
	return &patientRepository{conn}
}

func (p *patientRepository) GetAllPatients(patients *models.Patients) error {
	API := os.Getenv("COVID_API")
	res, err := p.conn.Get(API)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(patients)
}
