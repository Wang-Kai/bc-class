package controller

import (
	"encoding/json"
	"testing"

	"github.com/bc-class/model"
)

func TestReq2Struct(t *testing.T) {
	var jsonBody = `
		{"name":"bc-class","labels":{"app":"bc-class"},"strategy":"Recreate","pod":{"labels":{"app":"bc-class"},"containers":[{"name":"novnc","image":"uhub.service.ucloud.cn/safehouse/novnc","command":["/bin/sh"],"args":["-c","/usr/src/app/noVNC/utils/launch.sh --vnc localhost:5091"],"containerPorts":[{"container_port":6080}]},{"name":"ubuntu-xfce-vnc","image":"uhub.service.ucloud.cn/safehouse/ubuntu-xfce-vnc","containerPorts":[{"container_port":5901}]}]}}
	`

	deploymentConf := &model.DeploymentConf{}

	err := json.Unmarshal([]byte(jsonBody), deploymentConf)
	if err != nil {
		t.Fatal(err)
	}

	byteBody, _ := json.Marshal(deploymentConf)

	t.Logf("%s", byteBody)
}
