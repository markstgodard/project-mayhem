package spacemonkeys

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
	config   *Config
	Status   string
	director infrastructure.Director
}

type systemStatus struct {
	Status      string
	Deployments infrastructure.BoshDeployments
}

func NewProjectMayhem() *ProjectMayhem {
	cfg := LoadConfig()
	return &ProjectMayhem{
		config:   cfg,
		director: infrastructure.NewDirector(cfg.Deployment.Host, cfg.Deployment.Username, cfg.Deployment.Password),
		Status:   StatusStopped,
	}
}

func (pm *ProjectMayhem) StatusHandler(w http.ResponseWriter, r *http.Request) {
	deployments, err := pm.director.GetDeployments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	status := &systemStatus{
		Status:      pm.Status,
		Deployments: deployments,
	}
	json.NewEncoder(w).Encode(status)
}

func (pm *ProjectMayhem) ListDeploymentsHandler(w http.ResponseWriter, r *http.Request) {
	deployments, err := pm.director.GetDeployments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(deployments)
}

func (pm *ProjectMayhem) ListVmsHandler(w http.ResponseWriter, r *http.Request) {
	_, err := pm.director.GetDeployments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	vms := infrastructure.VMs{}
	json.NewEncoder(w).Encode(vms)
}
