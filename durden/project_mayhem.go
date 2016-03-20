package durden

import (
	"encoding/json"
	"net/http"
)

type systemStatus struct {
	Status           string
	TargetDeployment string
	VirtualMachines  VMs
}

type VM struct {
	Id    string
	State string
}
type VMs []VM

type ProjectMayhem struct {
	Status           string
	TargetDeployment string
	VirtualMachines  VMs
}

func NewProjectMayhem() *ProjectMayhem {
	return &ProjectMayhem{
		Status: "stopped",
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
