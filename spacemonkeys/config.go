package spacemonkeys

import (
	"os"

	log "github.com/Sirupsen/logrus"
)

type BoshDeployment struct {
	Host     string
	Username string
	Password string
}

type Config struct {
	Deployment BoshDeployment
}

func LoadConfig() *Config {
	log.Info("Loading config...")
	boshDeployment := BoshDeployment{
		Host:     os.Getenv("BOSH_DIRECTOR_HOST"),
		Username: os.Getenv("BOSH_DIRECTOR_USERNAME"),
		Password: os.Getenv("BOSH_DIRECTOR_PASSWORD"),
	}
	return &Config{
		Deployment: boshDeployment,
	}
}
