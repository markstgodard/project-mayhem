package infrastructure

type Deployment interface {
	ListVMs() (VMs, error)
}
