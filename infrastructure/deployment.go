//go:generate counterfeiter -o fakes/fake_deployment.go . Deployment
package infrastructure

type BoshDeployments []BoshDeployment

type Deployment interface {
	Name() string
	ListVMs() (VMs, error)
}

type BoshDeployment struct {
	Name string
}
