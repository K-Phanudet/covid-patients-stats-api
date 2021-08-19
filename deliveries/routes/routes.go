package routes

import (
	"net/http"
	"time"

	"github.com/K-Phanudet/covid-patients-stats-api/deliveries/controllers"
	"github.com/K-Phanudet/covid-patients-stats-api/repositories"
	"github.com/K-Phanudet/covid-patients-stats-api/usecases"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	repo := repositories.NewCovidRepository(&http.Client{Timeout: 10 * time.Second})
	useCase := usecases.NewPatientUseCase(repo)
	controller := controllers.NewPatientController(useCase)

	r := gin.Default()
	covid := r.Group("/covid")
	{
		covid.GET("summary", controller.GetSummaryPatientStats)
	}
	return r
}
