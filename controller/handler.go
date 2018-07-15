package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/bc-class/model"
	"github.com/bc-class/service"
	"github.com/bc-class/utils"
)

func ListDeployment(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	deployments, err := service.ListDeployment()
	if err != nil {
		utils.RespMsg(w, r, err)
	}

	utils.RespMsg(w, r, &model.ListDeploymentResp{
		Code: utils.OK,
		Data: deployments,
	})
}

func HandleAccess(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	deployment, user := params.ByName("deployment"), params.ByName("user")
	pod, err := service.HandleAccess(deployment, user)
	if err != nil {
		utils.RespMsg(w, r, err)
	}

	utils.RespMsg(w, r, &model.HandleAccessResp{
		Code: utils.OK,
		Data: pod,
	})
}
