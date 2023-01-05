package collect

import (
	"io"
	"os"

	"github.com/mattmattox/supportability-collector/modules/kubernetes"
	"gopkg.in/yaml.v2"
)

func RancherK8sYamlDir(dir string) string {
	rancherK8sYaml := dir + "/rancher-k8s-yaml"
	err := os.Mkdir(rancherK8sYaml, 0755)
	if err != nil {
		log.Warningln("Rancher install YAML folder creation failed - Error %s", err)
	}
	return rancherK8sYaml
}

func RancherAllNamespaceYaml(dir string) {
	client, err := kubernetes.GetClient()
	if err != nil {
		log.Warningln("Kubernetes client creation failed - Error %s", err)
	}
	namespaces, err := kubernetes.GetNamespaces(client)
	if err != nil {
		log.Warningln("Rancher namespace collection failed - Error %s", err)
	}
	namespaceDir := dir + "/rancher-all-namespace-yaml"
	err = os.Mkdir(namespaceDir, 0755)
	if err != nil {
		log.Info(err)
		log.Warningln("Rancher namespace folder creation failed - Error %s", err)
	}
	for _, namespace := range namespaces {
		log.Infof("Grabbing YAML for namespace: %s", namespace)
		namespaceData, err := kubernetes.GetNamespaceYaml(client, namespace)
		if err != nil {
			log.Warningln("Rancher namespace YAML collection failed - Error %s", err)
		}
		namespaceYaml, err := yaml.Marshal(namespaceData)
		if err != nil {
			log.Warningln("Rancher namespace YAML marshalling failed - Error %s", err)
		}
		file, err := os.Create(namespaceDir + "/" + namespace + ".yaml")
		if err != nil {
			log.Warningln("Rancher namespace YAML file creation failed - Error %s", err)
		}
		_, err = io.WriteString(file, string(namespaceYaml))
		if err != nil {
			log.Warningln("Rancher namespace YAML file write failed - Error %s", err)
		}
	}
}

func RancherK8sYamlPods(dir string) {
	client, err := kubernetes.GetClient()
	if err != nil {
		log.Warningln("Kubernetes client creation failed - Error %s", err)
	}
	pods, err := kubernetes.GetPods(client, "cattle-system")
	if err != nil {
		log.Warningln("List of pods in cattle-system failed - Error %s", err)
	}
	podDir := dir + "/pods"
	err = os.Mkdir(podDir, 0755)
	if err != nil {
		log.Warningln("Folder creation for pods in cattle-system failed - Error %s", err)
	}
	for _, pod := range pods {
		log.Infof("Grabbing YAML for pod: %s", pod)
		podData, err := kubernetes.GetPodYaml(client, "cattle-system", pod)
		if err != nil {
			log.Warningln("Pod YAML collection failed - Error %s", err)
		}
		podYaml, err := yaml.Marshal(podData)
		if err != nil {
			log.Warningln("Pod YAML marshalling failed - Error %s", err)
		}
		file, err := os.Create(podDir + "/" + pod + ".yaml")
		if err != nil {
			log.Warningln("Pod YAML file creation failed - Error %s", err)
		}
		defer file.Close()
		_, err = file.Write(podYaml)
		if err != nil {
			log.Warningln("Pod YAML file write failed - Error %s", err)
		}
	}
}

func RancherK8sYamlDeployments(dir string) {
	client, err := kubernetes.GetClient()
	if err != nil {
		log.Warningln("Kubernetes client creation failed - Error %s", err)
	}
	deployments, err := kubernetes.GetDeployments(client, "cattle-system")
	if err != nil {
		log.Warningln("List of deployments in cattle-system failed - Error %s", err)
	}
	deploymentDir := dir + "/deployments"
	err = os.Mkdir(deploymentDir, 0755)
	if err != nil {
		log.Warningln("Folder creation for deployments in cattle-system failed - Error %s", err)
	}
	for _, deployment := range deployments {
		log.Infof("Grabbing YAML for deployment: %s", deployment)
		deploymentData, err := kubernetes.GetDeploymentYaml(client, "cattle-system", deployment)
		if err != nil {
			log.Warningln("Deployment YAML collection failed - Error %s", err)
		}
		deploymentYaml, err := yaml.Marshal(deploymentData)
		if err != nil {
			log.Warningln("Deployment YAML marshalling failed - Error %s", err)
		}
		file, err := os.Create(deploymentDir + "/" + deployment + ".yaml")
		if err != nil {
			log.Warningln("Deployment YAML file creation failed - Error %s", err)
		}
		defer file.Close()
		_, err = file.Write(deploymentYaml)
		if err != nil {
			log.Warningln("Deployment YAML file write failed - Error %s", err)
		}
	}
}

func RancherK8sYamlDaemonSets(dir string) {
	client, err := kubernetes.GetClient()
	if err != nil {
		log.Warningln("Kubernetes client creation failed - Error %s", err)
	}
	daemonsets, err := kubernetes.GetDaemonSets(client, "cattle-system")
	if err != nil {
		log.Warningln("List of daemonsets in cattle-system failed - Error %s", err)
	}
	daemonsetDir := dir + "/daemonsets"
	err = os.Mkdir(daemonsetDir, 0755)
	if err != nil {
		log.Warningln("Folder creation for deployments in cattle-system failed - Error %s", err)
	}
	for _, daemonset := range daemonsets {
		log.Infof("Grabbing deployment: %s", daemonset)
		daemonsetData, err := kubernetes.GetDeploymentYaml(client, "cattle-system", daemonset)
		if err != nil {
			log.Warningln("DaemonSets YAML collection failed - Error %s", err)
		}
		daemonsetYaml, err := yaml.Marshal(daemonsetData)
		if err != nil {
			log.Warningln("DaemonSets YAML marshalling failed - Error %s", err)
		}
		file, err := os.Create(daemonsetDir + "/" + daemonset + ".yaml")
		if err != nil {
			log.Warningln("DaemonSets YAML file creation failed - Error %s", err)
		}
		defer file.Close()
		_, err = file.Write(daemonsetYaml)
		if err != nil {
			log.Warningln("DaemonSets YAML file write failed - Error %s", err)
		}
	}
}

func RancherK8sYamlStatefulSets(dir string) {
	client, err := kubernetes.GetClient()
	if err != nil {
		log.Warningln("Kubernetes client creation failed - Error %s", err)
	}
	statefulsets, err := kubernetes.GetStatefulSets(client, "cattle-system")
	if err != nil {
		log.Warningln("List of statefulsets in cattle-system failed - Error %s", err)
	}
	statefulsetDir := dir + "/statefulsets"
	err = os.Mkdir(statefulsetDir, 0755)
	if err != nil {
		log.Warningln("Folder creation for statefulsets in cattle-system failed - Error %s", err)
	}
	for _, statefulset := range statefulsets {
		log.Infof("Grabbing YAML for statefulset: %s", statefulset)
		statefulsetData, err := kubernetes.GetStatefulSetYaml(client, "cattle-system", statefulset)
		if err != nil {
			log.Warningln("Statefulset YAML collection failed - Error %s", err)
		}
		statefulsetYaml, err := yaml.Marshal(statefulsetData)
		if err != nil {
			log.Warningln("Statefulset YAML marshalling failed - Error %s", err)
		}
		file, err := os.Create(statefulsetDir + "/" + statefulset + ".yaml")
		if err != nil {
			log.Warningln("Statefulset YAML file creation failed - Error %s", err)
		}
		defer file.Close()
		_, err = file.Write(statefulsetYaml)
		if err != nil {
			log.Warningln("Statefulset YAML file write failed - Error %s", err)
		}
	}
}

func RancherK8sYamlCronjobs(dir string) {
	client, err := kubernetes.GetClient()
	if err != nil {
		log.Warningln("Kubernetes client creation failed - Error %s", err)
	}
	cronjobs, err := kubernetes.GetCronJobs(client, "cattle-system")
	if err != nil {
		log.Warningln("List of cronjobs in cattle-system failed - Error %s", err)
	}
	cronjobDir := dir + "/cronjobs"
	err = os.Mkdir(cronjobDir, 0755)
	if err != nil {
		log.Warningln("Folder creation for cronjobs in cattle-system failed - Error %s", err)
	}
	for _, cronjob := range cronjobs {
		log.Infof("Grabbing YAML for cronjob: %s", cronjob)
		cronjobData, err := kubernetes.GetCronJobYaml(client, "cattle-system", cronjob)
		if err != nil {
			log.Warningln("Cronjob YAML collection failed - Error %s", err)
		}
		cronjobYaml, err := yaml.Marshal(cronjobData)
		if err != nil {
			log.Warningln("Cronjob YAML marshalling failed - Error %s", err)
		}
		file, err := os.Create(cronjobDir + "/" + cronjob + ".yaml")
		if err != nil {
			log.Warningln("Cronjob YAML file creation failed - Error %s", err)
		}
		defer file.Close()
		_, err = file.Write(cronjobYaml)
		if err != nil {
			log.Warningln("Cronjob YAML file write failed - Error %s", err)
		}
	}
}

func RancherK8sYamlJobs(dir string) {
	client, err := kubernetes.GetClient()
	if err != nil {
		log.Warningln("Kubernetes client creation failed - Error %s", err)
	}
	jobs, err := kubernetes.GetJobs(client, "cattle-system")
	if err != nil {
		log.Warningln("List of jobs in cattle-system failed - Error %s", err)
	}
	jobDir := dir + "/jobs"
	err = os.Mkdir(jobDir, 0755)
	if err != nil {
		log.Warningln("Folder creation for jobs in cattle-system failed - Error %s", err)
	}
	for _, job := range jobs {
		log.Infof("Grabbing YAML for job: %s", job)
		jobData, err := kubernetes.GetJobYaml(client, "cattle-system", job)
		if err != nil {
			log.Warningln("Job YAML collection failed - Error %s", err)
		}
		jobYaml, err := yaml.Marshal(jobData)
		if err != nil {
			log.Warningln("Job YAML marshalling failed - Error %s", err)
		}
		file, err := os.Create(jobDir + "/" + job + ".yaml")
		if err != nil {
			log.Warningln("Job YAML file creation failed - Error %s", err)
		}
		defer file.Close()
		_, err = file.Write(jobYaml)
		if err != nil {
			log.Warningln("Job YAML file write failed - Error %s", err)
		}
	}
}

func RancherK8sYamlReplicaSets(dir string) {
	client, err := kubernetes.GetClient()
	if err != nil {
		log.Warningln("Kubernetes client creation failed - Error %s", err)
	}
	replicasets, err := kubernetes.GetReplicaSets(client, "cattle-system")
	if err != nil {
		log.Warningln("List of replicasets in cattle-system failed - Error %s", err)
	}
	replicasetDir := dir + "/replicasets"
	err = os.Mkdir(replicasetDir, 0755)
	if err != nil {
		log.Warningln("Folder creation for replicasets in cattle-system failed - Error %s", err)
	}
	for _, replicaset := range replicasets {
		log.Infof("Grabbing YAML for replicaset: %s", replicaset)
		replicasetData, err := kubernetes.GetReplicaSetYaml(client, "cattle-system", replicaset)
		if err != nil {
			log.Warningln("Replicaset YAML collection failed - Error %s", err)
		}
		replicasetYaml, err := yaml.Marshal(replicasetData)
		if err != nil {
			log.Warningln("Replicaset YAML marshalling failed - Error %s", err)
		}
		file, err := os.Create(replicasetDir + "/" + replicaset + ".yaml")
		if err != nil {
			log.Warningln("Replicaset YAML file creation failed - Error %s", err)
		}
		defer file.Close()
		_, err = file.Write(replicasetYaml)
		if err != nil {
			log.Warningln("Replicaset YAML file write failed - Error %s", err)
		}
	}
}

func RancherK8sYamlServices(dir string) {
	client, err := kubernetes.GetClient()
	if err != nil {
		log.Warningln("Kubernetes client creation failed - Error %s", err)
	}
	services, err := kubernetes.GetServices(client, "cattle-system")
	if err != nil {
		log.Warningln("List of services in cattle-system failed - Error %s", err)
	}
	serviceDir := dir + "/services"
	err = os.Mkdir(serviceDir, 0755)
	if err != nil {
		log.Warningln("Folder creation for services in cattle-system failed - Error %s", err)
	}
	for _, service := range services {
		log.Infof("Grabbing YAML for service: %s", service)
		serviceData, err := kubernetes.GetServiceYaml(client, "cattle-system", service)
		if err != nil {
			log.Warningln("Service YAML collection failed - Error %s", err)
		}
		serviceYaml, err := yaml.Marshal(serviceData)
		if err != nil {
			log.Warningln("Service YAML marshalling failed - Error %s", err)
		}
		file, err := os.Create(serviceDir + "/" + service + ".yaml")
		if err != nil {
			log.Warningln("Service YAML file creation failed - Error %s", err)
		}
		defer file.Close()
		_, err = file.Write(serviceYaml)
		if err != nil {
			log.Warningln("Service YAML file write failed - Error %s", err)
		}
	}
}

func RancherK8sYamlEndpoints(dir string) {
	client, err := kubernetes.GetClient()
	if err != nil {
		log.Warningln("Kubernetes client creation failed - Error %s", err)
	}
	endpoints, err := kubernetes.GetEndpoints(client, "cattle-system")
	if err != nil {
		log.Warningln("List of endpoints in cattle-system failed - Error %s", err)
	}
	endpointDir := dir + "/endpoints"
	err = os.Mkdir(endpointDir, 0755)
	if err != nil {
		log.Warningln("Folder creation for endpoints in cattle-system failed - Error %s", err)
	}
	for _, endpoint := range endpoints {
		log.Infof("Grabbing YAML for endpoint: %s", endpoint)
		endpointData, err := kubernetes.GetEndpointYaml(client, "cattle-system", endpoint)
		if err != nil {
			log.Warningln("Endpoint YAML collection failed - Error %s", err)
		}
		endpointYaml, err := yaml.Marshal(endpointData)
		if err != nil {
			log.Warningln("Endpoint YAML marshalling failed - Error %s", err)
		}
		file, err := os.Create(endpointDir + "/" + endpoint + ".yaml")
		if err != nil {
			log.Warningln("Endpoint YAML file creation failed - Error %s", err)
		}
		defer file.Close()
		_, err = file.Write(endpointYaml)
		if err != nil {
			log.Warningln("Endpoint YAML file write failed - Error %s", err)
		}
	}
}

func RancherK8sYamlIngresses(dir string) {
	client, err := kubernetes.GetClient()
	if err != nil {
		log.Warningln("Kubernetes client creation failed - Error %s", err)
	}
	ingresses, err := kubernetes.GetIngresses(client, "cattle-system")
	if err != nil {
		log.Warningln("List of ingresses in cattle-system failed - Error %s", err)
	}
	ingressDir := dir + "/ingresses"
	err = os.Mkdir(ingressDir, 0755)
	if err != nil {
		log.Warningln("Folder creation for ingresses in cattle-system failed - Error %s", err)
	}
	for _, ingress := range ingresses {
		log.Infof("Grabbing YAML for ingress: %s", ingress)
		ingressData, err := kubernetes.GetIngressYaml(client, "cattle-system", ingress)
		if err != nil {
			log.Warningln("Ingress YAML collection failed - Error %s", err)
		}
		ingressYaml, err := yaml.Marshal(ingressData)
		if err != nil {
			log.Warningln("Ingress YAML marshalling failed - Error %s", err)
		}
		file, err := os.Create(ingressDir + "/" + ingress + ".yaml")
		if err != nil {
			log.Warningln("Ingress YAML file creation failed - Error %s", err)
		}
		defer file.Close()
		_, err = file.Write(ingressYaml)
		if err != nil {
			log.Warningln("Ingress YAML file write failed - Error %s", err)
		}
	}
}
