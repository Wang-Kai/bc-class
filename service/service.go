package service

import (
	"encoding/json"
	"fmt"

	"github.com/apex/log"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
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
		log.Error("create in-cluster config error")
		panic(err.Error())
	}
	// creates the clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Error("create clientset error")
		panic(err.Error())
	}
}

// CreateDeployment create deployment in k8s
func CreateDeployment(deploymentConf *model.DeploymentConf) (err error) {
	// generate Container
	var containerArr = make([]apiv1.Container, len(deploymentConf.Pod.Containers))

	for index, containerConf := range deploymentConf.Pod.Containers {
		// generate ContainerPorts for this container
		var containerPortArr = make([]apiv1.ContainerPort, len(containerConf.ContainerPorts))
		for index, containerPortConf := range containerConf.ContainerPorts {
			containerPortArr[index] = apiv1.ContainerPort{
				Name:          containerPortConf.Name,
				ContainerPort: containerPortConf.ContainerPort,
			}
		}

		containerArr[index] = apiv1.Container{
			Name:    containerConf.Name,
			Image:   containerConf.Image,
			Ports:   containerPortArr,
			Command: containerConf.Command,
			Args:    containerConf.Args,
		}
	}

	body, _ := json.Marshal(containerArr)
	log.Infof("%s", body)

	// generate a deployment struct by deployment config
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deploymentConf.Name,
			Namespace: k8sNamespace,
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: deploymentConf.Labels,
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: deploymentConf.Pod.Labels,
				},
				Spec: apiv1.PodSpec{
					Containers: containerArr,
				},
			},
		},
	}

	deploymentClient := clientset.AppsV1().Deployments(k8sNamespace)
	_, err = deploymentClient.Create(deployment)
	if err != nil {
		log.WithError(err).Error("CreateDeploymentError")
	}

	return
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
			IP:   pod.Status.PodIP,
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
