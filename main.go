package main

import (
	"github.com/K-Phanudet/covid-patients-stats-api/deliveries/routes"
)

func main() {
	engin := routes.SetupRouter()
	engin.Run(":8080")
}
