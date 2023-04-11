package pod

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type PodClient struct {
	client *kubernetes.Clientset
	ctx    context.Context
}

type Pod interface {
	ListPods(ns string) []string
	DeletePod(ns, podname string)
	Logs(ns, podname string)
}

func NewPodClient(c *kubernetes.Clientset) *PodClient {
	ctx := context.TODO()
	return &PodClient{
		client: c,
		ctx:    ctx,
	}
}

func (c *PodClient) ListPods(ns string) []string {
	var podname []string
	pods, err := c.client.CoreV1().Pods(ns).List(c.ctx, metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, v := range pods.Items {
		podname = append(podname, v.Name)
	}
	return podname
}

func (c *PodClient) DeletePod(ns, podname string) {
	err := c.client.CoreV1().Pods(ns).Delete(context.TODO(), podname, metav1.DeleteOptions{})
	if err != nil {
		log.Fatal(err)
	}
}

func (c *PodClient) Logs(ns, podname string) {
	req := c.client.CoreV1().Pods(ns).GetLogs(podname, &v1.PodLogOptions{Follow: true})

	stream, err := req.Stream(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	r := bufio.NewReader(stream)
	for {
		bytes, err := r.ReadBytes('\n')
		message := string(bytes)
		fmt.Print(message)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
		}
	}
}
