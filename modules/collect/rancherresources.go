package collect

import (
	"os"

	"github.com/mattmattox/supportability-collector/modules/kubernetes"
	"gopkg.in/yaml.v2"
	"k8s.io/client-go/rest"
)

func RancherResourcesDir(dir string) string {
	rancherDataDir := dir + "/rancher-resources"
	err := os.Mkdir(rancherDataDir, 0755)
	if err != nil {
		log.Warningln("Rancher install YAML folder creation failed - Error %s", err)
	}
	return rancherDataDir
}

func RancherResourcesClusters(config *rest.Config, dir string) {
	clusters, err := kubernetes.GetRancherClusters(config)
	if err != nil {
		log.Warningln("Rancher get clusters failed - Error %s", err)
	}
	clusterDir := dir + "/clusters"
	err = os.Mkdir(clusterDir, 0755)
	if err != nil {
		log.Warningln("Rancher clusters YAML folder creation failed - Error %s", err)
	}
	for _, cluster := range clusters {
		log.Infof("Grabbing Rancher cluster: %s", cluster)
		clusterData, err := kubernetes.GetRancherClusterYaml(config, cluster)
		if err != nil {
			log.Warningln("Rancher cluster YAML collection failed - Error %s", err)
		}
		clusterYaml, err := yaml.Marshal(clusterData)
		if err != nil {
			log.Warningln("Rancher cluster YAML marshalling failed - Error %s", err)
		}
		file, err := os.Create(clusterDir + "/" + cluster + ".yaml")
		if err != nil {
			log.Warningln("Rancher cluster YAML file creation failed - Error %s", err)
		}
		defer file.Close()
		_, err = file.Write(clusterYaml)
		if err != nil {
			log.Warningln("Rancher cluster YAML file write failed - Error %s", err)
		}
	}
}

func RancherResourcesClusterNodes(config *rest.Config, dir string) {
	clusters, err := kubernetes.GetRancherClusters(config)
	if err != nil {
		log.Warningln("Rancher cluster nodes failed - Error %s", err)
	}
	clusterDir := dir + "/cluster-nodes"
	err = os.Mkdir(clusterDir, 0755)
	if err != nil {
		log.Warningln("Rancher cluster node YAML folder creation failed - Error %s", err)
	}
	for _, cluster := range clusters {
		log.Infof("Grabbing Rancher cluster nodes: %s", cluster)
		clusterNodes, err := kubernetes.GetRancherClusterNodes(config, cluster)
		if err != nil {
			log.Warningln("Rancher cluster node collection failed - Error %s", err)
		}
		clusterNodeDir := clusterDir + "/" + cluster
		err = os.Mkdir(clusterNodeDir, 0755)
		if err != nil {
			log.Warningln("Rancher cluster node YAML folder creation failed - Error %s", err)
		}
		for _, node := range clusterNodes {
			log.Infof("Grabbing Rancher cluster node: %s", node)
			nodeData, err := kubernetes.GetRancherClusterNodeYaml(config, cluster, node)
			if err != nil {
				log.Warningln("Rancher cluster node YAML collection failed - Error %s", err)
			}
			nodeYaml, err := yaml.Marshal(nodeData)
			if err != nil {
				log.Warningln("Rancher cluster node YAML marshalling failed - Error %s", err)
			}
			file, err := os.Create(clusterNodeDir + "/" + node + ".yaml")
			if err != nil {
				log.Warningln("Rancher cluster node YAML file creation failed - Error %s", err)
			}
			defer file.Close()
			_, err = file.Write(nodeYaml)
			if err != nil {
				log.Warningln("Rancher cluster node YAML file write failed - Error %s", err)
			}
		}
	}
}

func RancherResourcesClusterNodePools(config *rest.Config, dir string) {
	clusters, err := kubernetes.GetRancherClusters(config)
	if err != nil {
		log.Warningln("Rancher cluster collection failed")
	}
	clusterDir := dir + "/cluster-node-pools"
	err = os.Mkdir(clusterDir, 0755)
	if err != nil {
		log.Warningln("Rancher cluster node pool YAML folder creation failed")
	}
	for _, cluster := range clusters {
		log.Infof("Grabbing Rancher cluster node pools: %s", cluster)
		clusterNodePools, err := kubernetes.GetRancherClusterNodePools(config, cluster)
		if err != nil {
			log.Warningln("Rancher cluster node pool collection failed")
		}
		clusterNodePoolDir := clusterDir + "/" + cluster
		err = os.Mkdir(clusterNodePoolDir, 0755)
		if err != nil {
			log.Warningln("Rancher cluster node pool YAML folder creation failed")
		}
		for _, nodePool := range clusterNodePools {
			log.Infof("Grabbing Rancher cluster node pool: %s", nodePool)
			nodePoolData, err := kubernetes.GetRancherClusterNodePoolYaml(config, cluster, nodePool)
			if err != nil {
				log.Warningln("Rancher cluster node pool YAML collection failed")
			}
			nodePoolYaml, err := yaml.Marshal(nodePoolData)
			if err != nil {
				log.Warningln("Rancher cluster node pool YAML marshalling failed")
			}
			file, err := os.Create(clusterNodePoolDir + "/" + nodePool + ".yaml")
			if err != nil {
				log.Warningln("Rancher cluster node pool YAML file creation failed")
			}
			defer file.Close()
			_, err = file.Write(nodePoolYaml)
			if err != nil {
				log.Warningln("Rancher cluster node pool YAML file write failed")
			}
		}
	}
}

func RancherResourcesClusterNodeTemplates(config *rest.Config, dir string) {
	clusters, err := kubernetes.GetRancherClusters(config)
	if err != nil {
		log.Warningln("Rancher cluster collection failed")
	}
	clusterDir := dir + "/cluster-node-templates"
	err = os.Mkdir(clusterDir, 0755)
	if err != nil {
		log.Warningln("Rancher cluster node template YAML folder creation failed")
	}
	for _, cluster := range clusters {
		log.Infof("Grabbing Rancher cluster node templates: %s", cluster)
		clusterNodeTemplates, err := kubernetes.GetRancherClusterNodeTemplates(config, cluster)
		if err != nil {
			log.Warningln("Rancher cluster node template collection failed")
		}
		clusterNodeTemplateDir := clusterDir + "/" + cluster
		err = os.Mkdir(clusterNodeTemplateDir, 0755)
		if err != nil {
			log.Warningln("Rancher cluster node template YAML folder creation failed")
		}
		for _, nodeTemplate := range clusterNodeTemplates {
			log.Infof("Grabbing Rancher cluster node template: %s", nodeTemplate)
			nodeTemplateData, err := kubernetes.GetRancherClusterNodeTemplateYaml(config, cluster, nodeTemplate)
			if err != nil {
				log.Warningln("Rancher cluster node template YAML collection failed")
			}
			nodeTemplateYaml, err := yaml.Marshal(nodeTemplateData)
			if err != nil {
				log.Warningln("Rancher cluster node template YAML marshalling failed")
			}
			file, err := os.Create(clusterNodeTemplateDir + "/" + nodeTemplate + ".yaml")
			if err != nil {
				log.Warningln("Rancher cluster node template YAML file creation failed")
			}
			defer file.Close()
			_, err = file.Write(nodeTemplateYaml)
			if err != nil {
				log.Warningln("Rancher cluster node template YAML file write failed")
			}
		}
	}
}

func RancherResourcesClusterTemplates(config *rest.Config, dir string) {
	clusterTemplates, err := kubernetes.GetRancherClusterTemplates(config)
	if err != nil {
		log.Warningln("Rancher cluster template collection failed")
	}
	clusterTemplateDir := dir + "/cluster-templates"
	err = os.Mkdir(clusterTemplateDir, 0755)
	if err != nil {
		log.Warningln("Rancher cluster template YAML folder creation failed")
	}
	for _, clusterTemplate := range clusterTemplates {
		log.Infof("Grabbing Rancher cluster template: %s", clusterTemplate)
		clusterTemplateData, err := kubernetes.GetRancherClusterTemplateYaml(config, clusterTemplate)
		if err != nil {
			log.Warningln("Rancher cluster template YAML collection failed")
		}
		clusterTemplateYaml, err := yaml.Marshal(clusterTemplateData)
		if err != nil {
			log.Warningln("Rancher cluster template YAML marshalling failed")
		}
		file, err := os.Create(clusterTemplateDir + "/" + clusterTemplate + ".yaml")
		if err != nil {
			log.Warningln("Rancher cluster template YAML file creation failed")
		}
		defer file.Close()
		_, err = file.Write(clusterTemplateYaml)
		if err != nil {
			log.Warningln("Rancher cluster template YAML file write failed")
		}
	}
}

func RancherResourcesClusterTemplateRevisions(config *rest.Config, dir string) {
	clusterTemplates, err := kubernetes.GetRancherClusterTemplates(config)
	if err != nil {
		log.Warningln("Rancher cluster template collection failed")
	}
	clusterTemplateRevisionDir := dir + "/cluster-template-revisions"
	err = os.Mkdir(clusterTemplateRevisionDir, 0755)
	if err != nil {
		log.Warningln("Rancher cluster template revision YAML folder creation failed")
	}
	for _, clusterTemplate := range clusterTemplates {
		log.Infof("Grabbing Rancher cluster template revisions: %s", clusterTemplate)
		clusterTemplateRevisions, err := kubernetes.GetRancherClusterTemplateRevisions(config, clusterTemplate)
		if err != nil {
			log.Warningln("Rancher cluster template revision collection failed")
		}
		clusterTemplateRevisionTemplateDir := clusterTemplateRevisionDir + "/" + clusterTemplate
		err = os.Mkdir(clusterTemplateRevisionTemplateDir, 0755)
		if err != nil {
			log.Warningln("Rancher cluster template revision YAML folder creation failed")
		}
		for _, clusterTemplateRevision := range clusterTemplateRevisions {
			log.Infof("Grabbing Rancher cluster template revision: %s", clusterTemplateRevision)
			clusterTemplateRevisionData, err := kubernetes.GetRancherClusterTemplateRevisionYaml(config, clusterTemplate, clusterTemplateRevision)
			if err != nil {
				log.Warningln("Rancher cluster template revision YAML collection failed")
			}
			clusterTemplateRevisionYaml, err := yaml.Marshal(clusterTemplateRevisionData)
			if err != nil {
				log.Warningln("Rancher cluster template revision YAML marshalling failed")
			}
			file, err := os.Create(clusterTemplateRevisionTemplateDir + "/" + clusterTemplateRevision + ".yaml")
			if err != nil {
				log.Warningln("Rancher cluster template revision YAML file creation failed")
			}
			defer file.Close()
			_, err = file.Write(clusterTemplateRevisionYaml)
			if err != nil {
				log.Warningln("Rancher cluster template revision YAML file write failed")
			}
		}
	}
}

func RancherResourcesFeatures(config *rest.Config, dir string) {
	features, err := kubernetes.GetRancherFeatures(config)
	if err != nil {
		log.Warningln("Rancher feature collection failed")
	}
	featureDir := dir + "/features"
	err = os.Mkdir(featureDir, 0755)
	if err != nil {
		log.Warningln("Rancher feature YAML folder creation failed")
	}
	for _, featureid := range features {
		log.Infof("Grabbing Rancher feature: %s", featureid)
		featureData, err := kubernetes.GetRancherFeatureYaml(config, featureid)
		if err != nil {
			log.Warningln("Rancher feature YAML collection failed")
		}
		featureYaml, err := yaml.Marshal(featureData)
		if err != nil {
			log.Warningln("Rancher feature YAML marshalling failed")
		}
		file, err := os.Create(featureDir + "/" + featureid + ".yaml")
		if err != nil {
			log.Warningln("Rancher feature YAML file creation failed")
		}
		defer file.Close()
		_, err = file.Write(featureYaml)
		if err != nil {
			log.Warningln("Rancher feature YAML file write failed")
		}
	}
}
