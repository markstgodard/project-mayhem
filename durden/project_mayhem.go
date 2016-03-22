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
	config           *Config
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
		config: LoadConfig(),
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
