package collect

import (
	"time"
)

type UpstreamClusterInfo struct {
	Timestamp time.Time `yaml:"timestamp"`
	Version   string    `yaml:"version"`
	UUID      string    `yaml:"uuid"`
	ServerUrl string    `yaml:"serverUrl"`
	EulaDate  string    `yaml:"eulaDate"`
}

func CollectUpstreamCluster(tempDirRoot string) {
	log.Infoln("Collecting upstream cluster information")

	upstreamClusterDir := UpstreamDataDir(tempDirRoot)

	// Collecting details about the cluster
	UpstreamClusterNodes(upstreamClusterDir)

}
