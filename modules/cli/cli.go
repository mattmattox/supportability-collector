package cli

import (
	"os"

	"github.com/mattmattox/supportability-collector/modules/logging"
)

type Cli struct {
	HealthCheckPort  string
	RancherAccessKey string
	RancherSecretKey string
}

var log = logging.SetupLogging()

func Settings() Cli {
	healthCheckPort := os.Getenv("HEALTH_CHECK_PORT")
	if healthCheckPort == "" {
		healthCheckPort = "9000"
	}
	rancherAccessKey := os.Getenv("RANCHER_ACCESS_KEY")
	if rancherAccessKey == "" {
		log.Fatal("RANCHER_ACCESS_KEY is not set")
	}
	rancherSecretKey := os.Getenv("RANCHER_SECRET_KEY")
	if rancherSecretKey == "" {
		log.Fatal("RANCHER_SECRET_KEY is not set")
	}

	settings := Cli{
		HealthCheckPort:  healthCheckPort,
		RancherAccessKey: rancherAccessKey,
		RancherSecretKey: rancherSecretKey,
	}

	return settings
}
