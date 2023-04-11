package deploy

import (
	"context"
	"log"

	"k8s.io/client-go/kubernetes"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeployClient struct {
	client *kubernetes.Clientset
	ctx    context.Context
}

type Deploy interface {
	ListDeploys(ns string) []string
	DeleteDeploy(ns, deployname string)
}

func NewDeployClient(c *kubernetes.Clientset) *DeployClient {
	ctx := context.TODO()
	return &DeployClient{
		client: c,
		ctx:    ctx,
	}
}

func (c *DeployClient) ListDeploys(ns string) []string {
	var podname []string
	pods, err := c.client.AppsV1().Deployments(ns).List(c.ctx, metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, v := range pods.Items {
		podname = append(podname, v.Name)
	}
	return podname
}

func (c *DeployClient) DeleteDeploy(ns, deployname string) {
	deletePolicy := metav1.DeletePropagationForeground
	err := c.client.AppsV1().Deployments(ns).Delete(context.TODO(), deployname, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
	if err != nil {
		log.Fatal(err)
	}
}
