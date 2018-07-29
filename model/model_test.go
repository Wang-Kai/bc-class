package model

import (
	"encoding/json"
	"testing"

	apiv1 "k8s.io/api/core/v1"
)

func TestNewDeployment(t *testing.T) {
	var deploymentConf = &DeploymentConf{
		Name: "bc-class",
		Labels: map[string]string{
			"app": "bc-class",
		},
		Strategy: "Recreate",

		Pod: &PodConf{
			Labels: map[string]string{
				"app": "bc-class",
			},
			Containers: []*ContainerConf{
				&ContainerConf{
					Name:    "novnc",
					Image:   "uhub.service.ucloud.cn/safehouse/novnc",
					Command: []string{"/bin/sh"},
					Args:    []string{"-c", "/usr/src/app/noVNC/utils/launch.sh --vnc localhost:5091"},
					ContainerPorts: []*ContainerPortConf{
						&ContainerPortConf{
							ContainerPort: 6080,
						},
					},
				},

				&ContainerConf{
					Name:  "ubuntu-xfce-vnc",
					Image: "uhub.service.ucloud.cn/safehouse/ubuntu-xfce-vnc",
					ContainerPorts: []*ContainerPortConf{
						&ContainerPortConf{
							ContainerPort: 5901,
						},
					},
				},
			},
		},
	}

	res, _ := json.Marshal(deploymentConf)
	t.Logf("%s", res)

	var containerArr = make([]apiv1.Container, len(deploymentConf.Pod.Containers))

	t.Logf("Length of containers is %d \n", len(deploymentConf.Pod.Containers))

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

		containerJSON, _ := json.Marshal(containerArr)
		t.Logf("===> %s\n", containerJSON)
	}

}
