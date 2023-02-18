package client

import (
	"pain/pin8s/client/config"
	"pain/pin8s/client/pod"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type K8sClient struct {
	Pod    pod.Pod
	Config config.Config
}

func NewK8sClient() (*K8sClient, error) {
	client, err := loadClient()
	if err != nil {
		return nil, err
	}

	pod := pod.NewPodClient(client)
	config := config.NewConfigClient(clientcmd.NewDefaultPathOptions())

	return &K8sClient{
		Pod:    pod,
		Config: config,
	}, nil
}

func loadClient() (*kubernetes.Clientset, error) {
	config, err := getConfig()
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

func getConfig() (*rest.Config, error) {
	flags := getFlags()
	//return flags.ToRawKubeConfigLoader().ClientConfig()
	return flags.ToRESTConfig()
}

func getFlags() *genericclioptions.ConfigFlags {
	//return genericclioptions.NewConfigFlags(true).ToRawKubeConfigLoader()
	return genericclioptions.NewConfigFlags(true)
}
