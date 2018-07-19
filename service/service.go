package service

import (
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/retry"

	"github.com/bc-class/model"
)

var clientset *kubernetes.Clientset

const k8sNamespace = "prj-tin"

// init k8s clientset
func init() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Printf("create in-cluster config error")
		panic(err.Error())
	}
	// creates the clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Printf("create clientset error")
		panic(err.Error())
	}
}

// ListDeployment return all deployment for special namespace
func ListDeployment() (deployments []*model.Deployment, err error) {
	deploymentClient := clientset.AppsV1().Deployments(k8sNamespace)
	list, err := deploymentClient.List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, d := range list.Items {
		tmpDeployment := &model.Deployment{
			Name:     d.Name,
			Avalable: d.Status.AvailableReplicas,
		}
		deployments = append(deployments, tmpDeployment)
	}

	return
}

// ScaleDeployment scale deployment, make it plus amount
func ScaleDeployment(name string, amount int32) error {
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		deploymentClient := clientset.AppsV1().Deployments(k8sNamespace)

		deployment, getErr := deploymentClient.Get(name, metav1.GetOptions{})
		if getErr != nil {
			return getErr
		}

		deployment.Spec.Replicas = int32Ptr(*deployment.Spec.Replicas + amount)

		_, updateErr := deploymentClient.Update(deployment)
		return updateErr
	})
}

// ListPods list all pod in special deployment
func ListPods(deployment string) ([]*model.Pod, error) {
	podList, listErr := clientset.CoreV1().Pods(k8sNamespace).List(metav1.ListOptions{
		LabelSelector: fmt.Sprintf("app=%s", deployment),
	})
	if listErr != nil {
		return nil, listErr
	}

	podArr := []*model.Pod{}

	for _, pod := range podList.Items {
		tmpPod := &model.Pod{
			Name: pod.Name,
			IP:   pod.Status.HostIP,
		}
		podArr = append(podArr, tmpPod)
	}

	return podArr, nil
}

// DeletePod delete a pod by pod name
func DeletePod(pod string) error {
	return clientset.CoreV1().Pods(k8sNamespace).Delete(pod, &metav1.DeleteOptions{})
}

func int32Ptr(i int32) *int32 {
	return &i
}
