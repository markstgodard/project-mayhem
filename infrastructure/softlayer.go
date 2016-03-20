package infrastructure

import (
	log "github.com/Sirupsen/logrus"
	"github.com/cloudfoundry/bosh-softlayer-cpi/softlayer/vm"
)

type Softlayer struct {
}

func (sl *Softlayer) destroy(id string) error {

	_ = vm.SoftLayerVM{}

	log.Printf("About to delete vm [%s]..")

	return nil
}
