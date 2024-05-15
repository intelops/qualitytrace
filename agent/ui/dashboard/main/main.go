package main

import (
	"context"
	"fmt"

	"github.com/intelops/qualitytrace/agent/ui/dashboard"
	"github.com/intelops/qualitytrace/agent/ui/dashboard/models"
	"github.com/intelops/qualitytrace/agent/ui/dashboard/sensors"
)

func main() {
	err := dashboard.StartDashboard(context.Background(), models.EnvironmentInformation{
		OrganizationID: "Ana",
		EnvironmentID:  "Empregada",
		AgentVersion:   "0.15.5",
	}, sensors.NewSensor())
  
	if err != nil {
		fmt.Println(err.Error())
	}
}
