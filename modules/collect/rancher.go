package collect

import (
	"os"
	"time"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type RancherInfo struct {
	Timestamp time.Time `yaml:"timestamp"`
	Version   string    `yaml:"version"`
	UUID      string    `yaml:"uuid"`
	ServerUrl string    `yaml:"serverUrl"`
	EulaDate  string    `yaml:"eulaDate"`
}

func CollectRancherData(tempDirRoot string) {
	log.Infoln("Collecting Rancher Data")

	rancherDataDir := RancherDataDir(tempDirRoot)

	rancherInfo := RancherInfo{
		Timestamp: time.Now(),
	}

	log.Infoln("Connecting to upsteam cluster")
	config, err := rest.InClusterConfig()
	if err != nil {
		kubeconfig := os.Getenv("KUBECONFIG")
		if kubeconfig == "" {
			kubeconfig = os.Getenv("HOME") + "/.kube/config"
		}

		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			log.Fatalln("Failed to connect to upstream cluster")
		}
	}
	log.Infoln("Connected to upstream cluster")

	// Collecting details about Rancher itself
	rancherInfo.UUID = RancherDataUUID(config, rancherDataDir)
	rancherInfo.Version = RancherDataVersion(config, rancherDataDir)
	rancherInfo.ServerUrl = RancherDataServerUrl(config, rancherDataDir)
	rancherInfo.EulaDate = RancherDataEulaDate(config, rancherDataDir)

	// Writing this data to files
	RancherDataWriteYaml(rancherDataDir, &rancherInfo)
	RancherDataWriteJson(rancherDataDir, &rancherInfo)

	//Collecting YAML output for all namespaces in upstream cluster
	RancherAllNamespaceYaml(tempDirRoot)

	//Collecting YAMLs from cattle-system namespace
	rancherK8sYamlDir := RancherK8sYamlDir(tempDirRoot)
	RancherK8sYamlDeployments(rancherK8sYamlDir)
	RancherK8sYamlDaemonSets(rancherK8sYamlDir)
	RancherK8sYamlStatefulSets(rancherK8sYamlDir)
	RancherK8sYamlCronjobs(rancherK8sYamlDir)
	RancherK8sYamlJobs(rancherK8sYamlDir)
	RancherK8sYamlPods(rancherK8sYamlDir)
	RancherK8sYamlReplicaSets(rancherK8sYamlDir)
	RancherK8sYamlServices(rancherK8sYamlDir)
	RancherK8sYamlEndpoints(rancherK8sYamlDir)
	RancherK8sYamlIngresses(rancherK8sYamlDir)

	// Collecting Rancher resources
	rancherResourceDir := RancherResourcesDir(tempDirRoot)
	RancherResourcesClusters(config, rancherResourceDir)
	RancherResourcesClusterNodes(config, rancherResourceDir)
	RancherResourcesClusterNodePools(config, rancherResourceDir)
	//RancherResourcesClusterNodeTemplates(config, rancherResourceDir)
	//RancherResourcesClusterTemplates(config, rancherResourceDir)
	//RancherResourcesClusterTemplateRevisions(config, rancherResourceDir)
	//RancherResourcesFeatures(config, rancherResourceDir)

	log.Infoln("Rancher information collection complete")

}
