package infrastructure

type Provider interface {
	Destroy(id string) error
}
