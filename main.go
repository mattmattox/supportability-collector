package main

import (
	"github.com/mattmattox/supportability-collector/modules/cli"
	"github.com/mattmattox/supportability-collector/modules/health"
	"github.com/mattmattox/supportability-collector/modules/run"
)

func main() {
	settings := cli.Settings()
	health.PrintVersion()
	health.StartHealthServer(settings)
	run.Run(settings)
}
