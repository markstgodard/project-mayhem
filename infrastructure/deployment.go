//go:generate counterfeiter -o fakes/fake_deployment.go . Deployment
package infrastructure

type Deployment interface {
	ListVMs() (VMs, error)
}

type BoshDeployment struct {
	Name string
}
