package run

import (
	"github.com/mattmattox/supportability-collector/modules/cli"
	"github.com/mattmattox/supportability-collector/modules/collect"
	"github.com/mattmattox/supportability-collector/modules/logging"
)

var log = logging.SetupLogging()

func Run(cli.Cli) {
	log.Infoln("Starting Rancher Supportability Collector")
	collect.CollectData()
}
