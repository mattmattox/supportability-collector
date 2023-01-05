package kubernetes

import (
	"context"
	"encoding/json"
	"log"
	"os"

	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	networkingV1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/yaml"
)

// Data represents the JSON data returned by the API.
type Data struct {
	Value string `json:"value"`
}

// parseJSON parses the JSON data and returns the "value" field.
func parseJSON(data []byte) (string, error) {
	var d Data
	err := json.Unmarshal(data, &d)
	if err != nil {
		return "", err
	}
	return d.Value, nil
}

func GetClient() (*kubernetes.Clientset, error) {
	if os.Getenv("KUBECONFIG") != "" {
		// If the KUBECONFIG environment variable is set, use it to build the client configuration
		kubeConfigPath := os.Getenv("KUBECONFIG")
		kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
		if err != nil {
			return nil, err
		}
		client, err := kubernetes.NewForConfig(kubeConfig)
		if err != nil {
			return nil, err
		}
		return client, nil
	}

	// If the KUBECONFIG environment variable is not set, try to use the in-cluster configuration
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetRancherVersion(config *rest.Config) (string, error) {
	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: "management.cattle.io", Version: "v3"}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()
	crdClient, err := rest.RESTClientFor(&crdConfig)
	if err != nil {
		panic(err.Error())
	}
	result, err := crdClient.
		Get().
		AbsPath("/apis/management.cattle.io/v3/settings/server-version").
		DoRaw(context.TODO())
	if err != nil {
		panic(err.Error())
	}
	rancherVersion, err := parseJSON(result)
	if err != nil {
		log.Fatalln(err)
	}
	return rancherVersion, nil
}

func GetRancherUUID(config *rest.Config) (string, error) {
	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: "management.cattle.io", Version: "v3"}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()
	crdClient, err := rest.RESTClientFor(&crdConfig)
	if err != nil {
		panic(err.Error())
	}
	result, err := crdClient.
		Get().
		AbsPath("/apis/management.cattle.io/v3/settings/install-uuid").
		DoRaw(context.TODO())
	if err != nil {
		panic(err.Error())
	}
	rancherUUID, err := parseJSON(result)
	if err != nil {
		log.Fatalln(err)
	}
	return rancherUUID, nil
}

func GetRancherServerURL(config *rest.Config) (string, error) {
	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: "management.cattle.io", Version: "v3"}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()
	crdClient, err := rest.RESTClientFor(&crdConfig)
	if err != nil {
		panic(err.Error())
	}
	result, err := crdClient.
		Get().
		AbsPath("/apis/management.cattle.io/v3/settings/server-url").
		DoRaw(context.TODO())
	if err != nil {
		panic(err.Error())
	}
	rancherServerURL, err := parseJSON(result)
	if err != nil {
		log.Fatalln(err)
	}
	return rancherServerURL, nil
}

func GetRancherEulaDate(config *rest.Config) (string, error) {
	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: "management.cattle.io", Version: "v3"}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()
	crdClient, err := rest.RESTClientFor(&crdConfig)
	if err != nil {
		panic(err.Error())
	}
	result, err := crdClient.
		Get().
		AbsPath("/apis/management.cattle.io/v3/settings/eula-agreed").
		DoRaw(context.TODO())
	if err != nil {
		panic(err.Error())
	}
	rancherStartDate, err := parseJSON(result)
	if err != nil {
		log.Fatalln(err)
	}
	return rancherStartDate, nil
}

func GetRancherClusters(config *rest.Config) ([]string, error) {
	// Set up the CRD client configuration
	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: "management.cattle.io", Version: "v3"}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	// Create the CRD client
	crdClient, err := rest.RESTClientFor(&crdConfig)
	if err != nil {
		return nil, err
	}

	// Retrieve the list of clusters
	result, err := crdClient.
		Get().
		AbsPath("/apis/management.cattle.io/v3/clusters").
		DoRaw(context.TODO())
	if err != nil {
		return nil, err
	}

	// Define a struct to hold the complete response from the Rancher API
	type RancherResponse struct {
		APIVersion string `json:"apiVersion"`
		Items      []struct {
			APIVersion string `json:"apiVersion"`
			Kind       string `json:"kind"`
			Metadata   struct {
				Name string `json:"name"`
			}
		}
	}

	// Unmarshal the response into the RancherResponse struct
	var response RancherResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}

	var rancherClusters []string
	for _, item := range response.Items {
		rancherClusters = append(rancherClusters, item.Metadata.Name)
	}
	return rancherClusters, nil
}

func GetRancherClusterYaml(config *rest.Config, clusterID string) (string, error) {
	// Set up the CRD client configuration
	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: "management.cattle.io", Version: "v3"}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	// Create the CRD client
	crdClient, err := rest.RESTClientFor(&crdConfig)
	if err != nil {
		return "", err
	}

	// Retrieve the cluster data
	result, err := crdClient.
		Get().
		AbsPath("/apis/management.cattle.io/v3/clusters/" + clusterID).
		DoRaw(context.TODO())
	if err != nil {
		return "", err
	}

	// Convert the cluster data to YAML
	yamlData, err := yaml.JSONToYAML(result)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(yamlData), nil
}

func GetRancherClusterNodes(config *rest.Config, clusterID string) ([]string, error) {
	// Set up the CRD client configuration
	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: "management.cattle.io", Version: "v3"}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	// Create the CRD client
	crdClient, err := rest.RESTClientFor(&crdConfig)
	if err != nil {
		return nil, err
	}

	// Retrieve the list of nodes
	result, err := crdClient.
		Get().
		AbsPath("/apis/management.cattle.io/v3/nodes/").
		Param("namespace", clusterID).
		DoRaw(context.TODO())

	if err != nil {
		return nil, err
	}

	// Define a struct to hold the complete response from the Rancher API
	type RancherResponse struct {
		APIVersion string `json:"apiVersion"`
		Items      []struct {
			APIVersion string `json:"apiVersion"`
			Kind       string `json:"kind"`
			Metadata   struct {
				Name string `json:"name"`
			}
		}
	}

	// Unmarshal the response into the RancherResponse struct
	var response RancherResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}

	var rancherClusterNodes []string
	for _, item := range response.Items {
		rancherClusterNodes = append(rancherClusterNodes, item.Metadata.Name)
	}
	return rancherClusterNodes, nil
}

func GetRancherClusterNodeYaml(config *rest.Config, clusterID string, nodeID string) (string, error) {
	// Set up the CRD client configuration
	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: "management.cattle.io", Version: "v3"}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	// Create the CRD client
	crdClient, err := rest.RESTClientFor(&crdConfig)
	if err != nil {
		return "", err
	}

	// Retrieve the node data
	result, err := crdClient.
		Get().
		AbsPath("/apis/management.cattle.io/v3/nodes/").
		Param("namespace", clusterID).
		Param("name", nodeID).
		DoRaw(context.TODO())
	if err != nil {
		return "", err
	}

	// Convert the node data to YAML
	yamlData, err := yaml.JSONToYAML(result)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(yamlData), nil
}

func GetRancherClusterNodePools(config *rest.Config, clusterID string) ([]string, error) {
	// Set up the CRD client configuration
	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: "management.cattle.io", Version: "v3"}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	// Create the CRD client
	crdClient, err := rest.RESTClientFor(&crdConfig)
	if err != nil {
		return nil, err
	}

	// Retrieve the list of node pools
	result, err := crdClient.
		Get().
		AbsPath("/apis/management.cattle.io/v3/nodepools/").
		Param("namespace", clusterID).
		DoRaw(context.TODO())

	if err != nil {
		return nil, err
	}

	// Parse the list of node pools from the JSON data
	var rancherNodePools []string
	err = json.Unmarshal(result, &rancherNodePools)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return rancherNodePools, nil
}

func GetRancherClusterNodePoolYaml(config *rest.Config, clusterID string, nodePoolID string) (string, error) {
	// Set up the CRD client configuration
	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: "management.cattle.io", Version: "v3"}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	// Create the CRD client
	crdClient, err := rest.RESTClientFor(&crdConfig)
	if err != nil {
		return "", err
	}

	// Retrieve the node pool data
	result, err := crdClient.
		Get().
		AbsPath("/apis/management.cattle.io/v3/nodepools/").
		Param("namespace", clusterID).
		Param("name", nodePoolID).
		DoRaw(context.TODO())
	if err != nil {
		return "", err
	}

	// Convert the node pool data to YAML
	yamlData, err := yaml.JSONToYAML(result)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(yamlData), nil
}

func GetRancherClusterNodeTemplates(config *rest.Config, clusterID string) ([]string, error) {
	// Set up the CRD client configuration
	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: "management.cattle.io", Version: "v3"}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	// Create the CRD client
	crdClient, err := rest.RESTClientFor(&crdConfig)
	if err != nil {
		return nil, err
	}

	// Retrieve the list of node templates
	result, err := crdClient.
		Get().
		AbsPath("/apis/management.cattle.io/v3/clusters/" + clusterID + "/nodetemplates").
		DoRaw(context.TODO())

	if err != nil {
		return nil, err
	}

	// Parse the list of node templates from the JSON data
	var rancherNodeTemplates []string
	err = json.Unmarshal(result, &rancherNodeTemplates)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return rancherNodeTemplates, nil
}

func GetRancherClusterNodeTemplateYaml(config *rest.Config, clusterID string, nodeTemplateID string) (string, error) {
	// Set up the CRD client configuration
	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: "management.cattle.io", Version: "v3"}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	// Create the CRD client
	crdClient, err := rest.RESTClientFor(&crdConfig)
	if err != nil {
		return "", err
	}

	// Retrieve the node template data
	result, err := crdClient.
		Get().
		AbsPath("/apis/management.cattle.io/v3/clusters/" + clusterID + "/nodetemplates/" + nodeTemplateID).
		DoRaw(context.TODO())
	if err != nil {
		return "", err
	}

	// Convert the node template data to YAML
	yamlData, err := yaml.JSONToYAML(result)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(yamlData), nil
}

func GetRancherClusterTemplates(config *rest.Config) ([]string, error) {
	// Set up the CRD client configuration
	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: "management.cattle.io", Version: "v3"}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	// Create the CRD client
	crdClient, err := rest.RESTClientFor(&crdConfig)
	if err != nil {
		return nil, err
	}

	// Retrieve the list of cluster templates
	result, err := crdClient.
		Get().
		AbsPath("/apis/management.cattle.io/v3/clustertemplates").
		DoRaw(context.TODO())

	if err != nil {
		return nil, err
	}

	// Parse the list of cluster templates from the JSON data
	var rancherClusterTemplates []string
	err = json.Unmarshal(result, &rancherClusterTemplates)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return rancherClusterTemplates, nil
}

func GetRancherClusterTemplateYaml(config *rest.Config, clusterTemplateID string) (string, error) {
	// Set up the CRD client configuration
	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: "management.cattle.io", Version: "v3"}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	// Create the CRD client
	crdClient, err := rest.RESTClientFor(&crdConfig)
	if err != nil {
		return "", err
	}

	// Retrieve the cluster template data
	result, err := crdClient.
		Get().
		AbsPath("/apis/management.cattle.io/v3/clustertemplates/" + clusterTemplateID).
		DoRaw(context.TODO())
	if err != nil {
		return "", err
	}

	// Convert the cluster template data to YAML
	yamlData, err := yaml.JSONToYAML(result)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(yamlData), nil
}

func GetRancherClusterTemplateRevisions(config *rest.Config, clusterTemplateID string) ([]string, error) {
	// Set up the CRD client configuration
	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: "management.cattle.io", Version: "v3"}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	// Create the CRD client
	crdClient, err := rest.RESTClientFor(&crdConfig)
	if err != nil {
		return nil, err
	}

	// Retrieve the list of cluster template revisions
	result, err := crdClient.
		Get().
		AbsPath("/apis/management.cattle.io/v3/clustertemplates/" + clusterTemplateID + "/revisions").
		DoRaw(context.TODO())

	if err != nil {
		return nil, err
	}

	// Parse the list of cluster template revisions from the JSON data
	var rancherClusterTemplateRevisions []string
	err = json.Unmarshal(result, &rancherClusterTemplateRevisions)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return rancherClusterTemplateRevisions, nil
}

func GetRancherClusterTemplateRevisionYaml(config *rest.Config, clusterTemplateID string, clusterTemplateRevisionID string) (string, error) {
	// Set up the CRD client configuration
	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: "management.cattle.io", Version: "v3"}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	// Create the CRD client
	crdClient, err := rest.RESTClientFor(&crdConfig)
	if err != nil {
		return "", err
	}

	// Retrieve the cluster template revision data
	result, err := crdClient.
		Get().
		AbsPath("/apis/management.cattle.io/v3/clustertemplates/" + clusterTemplateID + "/revisions/" + clusterTemplateRevisionID).
		DoRaw(context.TODO())
	if err != nil {
		return "", err
	}

	// Convert the cluster template revision data to YAML
	yamlData, err := yaml.JSONToYAML(result)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(yamlData), nil
}

func GetRancherFeatures(config *rest.Config) ([]string, error) {
	// Set up the CRD client configuration
	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: "management.cattle.io", Version: "v3"}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	// Create the CRD client
	crdClient, err := rest.RESTClientFor(&crdConfig)
	if err != nil {
		return nil, err
	}

	// Retrieve the list of Rancher features
	result, err := crdClient.
		Get().
		AbsPath("/apis/management.cattle.io/v3/features").
		DoRaw(context.TODO())
	if err != nil {
		return nil, err
	}

	// Parse the list of Rancher features from the JSON data
	var rancherFeatures []string
	err = json.Unmarshal(result, &rancherFeatures)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return rancherFeatures, nil
}

func GetRancherFeatureYaml(config *rest.Config, featureID string) (string, error) {
	// Set up the CRD client configuration
	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: "management.cattle.io", Version: "v3"}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	// Create the CRD client
	crdClient, err := rest.RESTClientFor(&crdConfig)
	if err != nil {
		return "", err
	}

	// Retrieve the feature data
	result, err := crdClient.
		Get().
		AbsPath("/apis/management.cattle.io/v3/features/" + featureID).
		DoRaw(context.TODO())
	if err != nil {
		return "", err
	}

	// Convert the feature data to YAML
	yamlData, err := yaml.JSONToYAML(result)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(yamlData), nil
}

func GetRancherGlobalDNSProviders(config *rest.Config) (string, error) {
	// Set up the CRD client configuration
	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: "management.cattle.io", Version: "v3"}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	// Create the CRD client
	crdClient, err := rest.RESTClientFor(&crdConfig)
	if err != nil {
		return "", err
	}

	// Retrieve the global DNS providers data
	result, err := crdClient.
		Get().
		AbsPath("/apis/management.cattle.io/v3/globaldnsproviders").
		DoRaw(context.TODO())
	if err != nil {
		return "", err
	}

	// Convert the global DNS providers data to YAML
	yamlData, err := yaml.JSONToYAML(result)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(yamlData), nil
}

func GetNamespaces(client *kubernetes.Clientset) ([]string, error) {
	namespaces, err := client.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var namespaceList []string
	for _, namespace := range namespaces.Items {
		namespaceList = append(namespaceList, namespace.Name)
	}
	return namespaceList, nil
}

func GetNamespaceYaml(client *kubernetes.Clientset, namespace string) (*v1.Namespace, error) {
	ns, err := client.CoreV1().Namespaces().Get(context.Background(), namespace, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return ns, nil
}

func GetDeployments(client *kubernetes.Clientset, namespace string) ([]string, error) {
	deployments, err := client.AppsV1().Deployments(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var deploymentList []string
	for _, deployment := range deployments.Items {
		deploymentList = append(deploymentList, deployment.Name)
	}
	return deploymentList, nil
}

func GetDeploymentYaml(client *kubernetes.Clientset, namespace string, deployment string) (*appsv1.Deployment, error) {
	deploy, err := client.AppsV1().Deployments(namespace).Get(context.Background(), deployment, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return deploy, nil
}

func GetDaemonSets(client *kubernetes.Clientset, namespace string) ([]string, error) {
	daemonsets, err := client.AppsV1().DaemonSets(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var daemonsetList []string
	for _, daemonset := range daemonsets.Items {
		daemonsetList = append(daemonsetList, daemonset.Name)
	}
	return daemonsetList, nil
}

func GetDaemonSetYaml(client *kubernetes.Clientset, namespace string, daemonset string) (*appsv1.DaemonSet, error) {
	ds, err := client.AppsV1().DaemonSets(namespace).Get(context.Background(), daemonset, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return ds, nil
}

func GetStatefulSets(client *kubernetes.Clientset, namespace string) ([]string, error) {
	statefulsets, err := client.AppsV1().StatefulSets(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var statefulsetList []string
	for _, statefulset := range statefulsets.Items {
		statefulsetList = append(statefulsetList, statefulset.Name)
	}
	return statefulsetList, nil
}

func GetStatefulSetYaml(client *kubernetes.Clientset, namespace string, statefulset string) (*appsv1.StatefulSet, error) {
	ss, err := client.AppsV1().StatefulSets(namespace).Get(context.Background(), statefulset, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return ss, nil
}

func GetJobs(client *kubernetes.Clientset, namespace string) ([]string, error) {
	jobs, err := client.BatchV1().Jobs(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var jobList []string
	for _, job := range jobs.Items {
		jobList = append(jobList, job.Name)
	}
	return jobList, nil
}

func GetJobYaml(client *kubernetes.Clientset, namespace string, job string) (*batchv1.Job, error) {
	j, err := client.BatchV1().Jobs(namespace).Get(context.Background(), job, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func GetCronJobs(client *kubernetes.Clientset, namespace string) ([]string, error) {
	cronjobs, err := client.BatchV1().CronJobs(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var cronjobList []string
	for _, cronjob := range cronjobs.Items {
		cronjobList = append(cronjobList, cronjob.Name)
	}
	return cronjobList, nil
}

func GetCronJobYaml(client *kubernetes.Clientset, namespace string, cronjob string) (*batchv1.CronJob, error) {
	cj, err := client.BatchV1().CronJobs(namespace).Get(context.Background(), cronjob, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return cj, nil
}

func GetPods(client *kubernetes.Clientset, namespace string) ([]string, error) {
	pods, err := client.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var podList []string
	for _, pod := range pods.Items {
		podList = append(podList, pod.Name)
	}
	return podList, nil
}

func GetPodYaml(client *kubernetes.Clientset, namespace string, pod string) (*v1.Pod, error) {
	p, err := client.CoreV1().Pods(namespace).Get(context.Background(), pod, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return p, nil
}

func GetServices(client *kubernetes.Clientset, namespace string) ([]string, error) {
	services, err := client.CoreV1().Services(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var serviceList []string
	for _, service := range services.Items {
		serviceList = append(serviceList, service.Name)
	}
	return serviceList, nil
}

func GetServiceYaml(client *kubernetes.Clientset, namespace string, service string) (*v1.Service, error) {
	s, err := client.CoreV1().Services(namespace).Get(context.Background(), service, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return s, nil
}

func GetEndpoints(client *kubernetes.Clientset, namespace string) ([]string, error) {
	endpoints, err := client.CoreV1().Endpoints(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var endpointList []string
	for _, endpoint := range endpoints.Items {
		endpointList = append(endpointList, endpoint.Name)
	}
	return endpointList, nil
}

func GetEndpointYaml(client *kubernetes.Clientset, namespace string, endpoint string) (*v1.Endpoints, error) {
	e, err := client.CoreV1().Endpoints(namespace).Get(context.Background(), endpoint, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return e, nil
}

func GetIngresses(client *kubernetes.Clientset, namespace string) ([]string, error) {
	ingresses, err := client.NetworkingV1().Ingresses(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var ingressList []string
	for _, ingress := range ingresses.Items {
		ingressList = append(ingressList, ingress.Name)
	}
	return ingressList, nil
}

func GetIngressYaml(client *kubernetes.Clientset, namespace string, ingress string) (*networkingV1.Ingress, error) {
	i, err := client.NetworkingV1().Ingresses(namespace).Get(context.Background(), ingress, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return i, nil
}

func GetPersistentVolumes(client *kubernetes.Clientset) ([]string, error) {
	pvs, err := client.CoreV1().PersistentVolumes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var pvList []string
	for _, pv := range pvs.Items {
		pvList = append(pvList, pv.Name)
	}
	return pvList, nil
}

func GetPersistentVolumeYaml(client *kubernetes.Clientset, pv string) (*v1.PersistentVolume, error) {
	p, err := client.CoreV1().PersistentVolumes().Get(context.Background(), pv, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return p, nil
}

func GetPersistentVolumeClaims(client *kubernetes.Clientset, namespace string) ([]string, error) {
	pvcs, err := client.CoreV1().PersistentVolumeClaims(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var pvcList []string
	for _, pvc := range pvcs.Items {
		pvcList = append(pvcList, pvc.Name)
	}
	return pvcList, nil
}

func GetPersistentVolumeClaimYaml(client *kubernetes.Clientset, namespace string, pvc string) (*v1.PersistentVolumeClaim, error) {
	p, err := client.CoreV1().PersistentVolumeClaims(namespace).Get(context.Background(), pvc, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return p, nil
}

func GetConfigMaps(client *kubernetes.Clientset, namespace string) ([]string, error) {
	configmaps, err := client.CoreV1().ConfigMaps(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var configmapList []string
	for _, configmap := range configmaps.Items {
		configmapList = append(configmapList, configmap.Name)
	}
	return configmapList, nil
}

func GetConfigMapYaml(client *kubernetes.Clientset, namespace string, configmap string) (*v1.ConfigMap, error) {
	cm, err := client.CoreV1().ConfigMaps(namespace).Get(context.Background(), configmap, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return cm, nil
}

func GetReplicaSets(client *kubernetes.Clientset, namespace string) ([]string, error) {
	replicasets, err := client.AppsV1().ReplicaSets(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var replicasetsList []string
	for _, replicasets := range replicasets.Items {
		replicasetsList = append(replicasetsList, replicasets.Name)
	}
	return replicasetsList, nil
}

func GetReplicaSetYaml(client *kubernetes.Clientset, namespace string, replicasets string) (*appsv1.ReplicaSet, error) {
	rs, err := client.AppsV1().ReplicaSets(namespace).Get(context.Background(), replicasets, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func GetNodes(client *kubernetes.Clientset) ([]string, error) {
	nodes, err := client.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var nodeList []string
	for _, node := range nodes.Items {
		nodeList = append(nodeList, node.Name)
	}
	return nodeList, nil
}

func GetNodeYaml(client *kubernetes.Clientset, node string) (*v1.Node, error) {
	n, err := client.CoreV1().Nodes().Get(context.Background(), node, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return n, nil
}
