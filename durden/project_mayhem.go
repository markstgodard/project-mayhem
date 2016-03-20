package durden

import (
	"encoding/json"
	"net/http"

	"github.com/markstgodard/project-mayhem/infrastructure"
)

const (
	StatusStopped = "stopped"
	StatusRunning = "running"
)

type ProjectMayhem struct {
	Status           string
	TargetDeployment string
	VirtualMachines  infrastructure.VMs
}

type systemStatus struct {
	Status           string
	TargetDeployment string
	VirtualMachines  infrastructure.VMs
}

func NewProjectMayhem() *ProjectMayhem {
	return &ProjectMayhem{
		Status: StatusStopped,
	}
}

func (pm *ProjectMayhem) StatusHandler(w http.ResponseWriter, r *http.Request) {
	status := &systemStatus{
		Status:           pm.Status,
		TargetDeployment: pm.TargetDeployment,
		VirtualMachines:  pm.VirtualMachines,
	}
	json.NewEncoder(w).Encode(status)
}
