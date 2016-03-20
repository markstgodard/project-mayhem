//go:generate counterfeiter -o fakes/fake_iaas.go . Iaas
package infrastructure

type Iaas interface {
	Destroy(id string) error
}
