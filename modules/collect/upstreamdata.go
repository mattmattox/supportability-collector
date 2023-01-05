package collect

import (
	"fmt"
	"os"

	"github.com/mattmattox/supportability-collector/modules/kubernetes"
)

func UpstreamDataDir(dir string) string {
	log.Infoln("Collecting upstream cluster information")
	upstreamDir := dir + "/upstream"
	err := os.MkdirAll(upstreamDir, 0755)
	if err != nil {
		log.Fatalln("Upstream cluster directory creation failed")
	}
	log.Infoln("Upstream cluster directory created successfully")
	return upstreamDir
}

func UpstreamClusterNodes(dir string) {
	log.Infoln("Collecting upstream cluster nodes")
	client, err := kubernetes.GetClient()
	if err != nil {
		log.Fatalln("Failed to connect to upstream cluster")
	}
	dir = dir + "/nodes"
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatalln("Upstream cluster nodes directory creation failed")
	}
	log.Infoln("Upstream cluster nodes directory created successfully")
	nodes, err := kubernetes.GetNodes(client)
	if err != nil {
		log.Fatalln("Failed to get upstream cluster nodes")
	}
	for _, node := range nodes {
		log.Infoln(node)
		nodeData, err := kubernetes.GetNodeYaml(client, node)
		if err != nil {
			log.Fatalln("Failed to get upstream cluster node data")
		}
		f, err := os.Create(dir + "/" + node + ".yaml")
		if err != nil {
			log.Fatalln("Upstream cluster node file creation failed")
		}
		fmt.Fprint(f, nodeData)
		f.Close()
		log.Infoln("Upstream cluster node file created successfully")
	}
}
