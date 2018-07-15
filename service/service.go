package service

import (
	"errors"
	"log"

	"github.com/gomodule/redigo/redis"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/retry"

	. "github.com/bc-class/db"
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

	// init redis db
	InitRedis("[::]:6379")
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

// HandleAccess scale deployent, and save user/podIP into redis
func HandleAccess(deploymentName, user string) (*model.Pod, error) {
	// scale deployment
	retry.RetryOnConflict(retry.DefaultRetry, func() error {
		deploymentClient := clientset.AppsV1().Deployments(k8sNamespace)
		deployment, getErr := deploymentClient.Get(deploymentName, metav1.GetOptions{})
		if getErr != nil {
			return getErr
		}

		deployment.Spec.Replicas = int32Ptr(*deployment.Spec.Replicas + 1)
		_, updateErr := deploymentClient.Update(deployment)
		return updateErr
	})

	// find an idle IP
	podsList, err := clientset.CoreV1().Pods(k8sNamespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	conn := Pool.Get()
	defer conn.Close()

	for _, pod := range podsList.Items {
		ip, name := pod.Status.HostIP, pod.Name

		exist, err := redis.Bool(conn.Do("HEXISTS", deploymentName, ip))
		if err != nil {
			return nil, err
		}

		if !exist {
			_, err := conn.Do("HSET", deploymentName, ip, user)
			if err != nil {
				return nil, err
			}

			return &model.Pod{
				Name: name,
				IP:   ip,
			}, nil
		}
	}

	return nil, errors.New("NoIdleIP")
}

func int32Ptr(i int32) *int32 {
	return &i
}
