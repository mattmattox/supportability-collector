package collect

import (
	"encoding/json"
	"os"

	"github.com/mattmattox/supportability-collector/modules/kubernetes"
	"gopkg.in/yaml.v3"
	"k8s.io/client-go/rest"
)

func RancherDataDir(dir string) string {
	rancherDataDir := dir + "/rancher-data"
	err := os.MkdirAll(rancherDataDir, 0755)
	if err != nil {
		log.Fatalln("Rancher install YAML folder creation failed")
	}
	return rancherDataDir
}

func RancherDataWriteYaml(dir string, RancherData *RancherInfo) {
	rancherDataYaml, err := os.Create(dir + "/rancher-data.yaml")
	if err != nil {
		log.Fatalln("Rancher install YAML file creation failed")
	}
	defer rancherDataYaml.Close()
	yamlEncoder := yaml.NewEncoder(rancherDataYaml)
	yamlEncoder.SetIndent(2)
	err = yamlEncoder.Encode(RancherData)
	if err != nil {
		log.Fatalln("Rancher install YAML file write failed")
	}
}

func RancherDataWriteJson(dir string, RancherData *RancherInfo) {
	rancherDataJson, err := os.Create(dir + "/rancher-data.json")
	if err != nil {
		log.Fatalln("Rancher install JSON file creation failed")
	}
	defer rancherDataJson.Close()
	jsonEncoder := json.NewEncoder(rancherDataJson)
	jsonEncoder.SetIndent("", "  ")
	err = jsonEncoder.Encode(RancherData)
	if err != nil {
		log.Fatalln("Rancher install JSON file write failed")
	}
}

func RancherDataVersion(config *rest.Config, dir string) string {
	rancherVersion, err := kubernetes.GetRancherVersion(config)
	if err != nil {
		log.Fatalln("Rancher version collection failed")
	}
	log.Infof("Rancher version: %s", rancherVersion)
	return rancherVersion
}

func RancherDataUUID(config *rest.Config, dir string) string {
	rancherUUID, err := kubernetes.GetRancherUUID(config)
	if err != nil {
		log.Fatalln("Rancher UUID collection failed")
	}
	log.Infof("Rancher UUID: %s", rancherUUID)
	return rancherUUID
}

func RancherDataServerUrl(config *rest.Config, dir string) string {
	rancherURL, err := kubernetes.GetRancherServerURL(config)
	if err != nil {
		log.Fatalln("Rancher Server URL collection failed")
	}
	log.Infof("Rancher Server URL: %s", rancherURL)
	return rancherURL
}

func RancherDataEulaDate(config *rest.Config, dir string) string {
	rancherEulaDate, err := kubernetes.GetRancherEulaDate(config)
	if err != nil {
		log.Fatalln("Rancher EULA date collection failed")
	}
	log.Infof("Rancher EULA date: %s", rancherEulaDate)
	return rancherEulaDate
}
