package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/bc-class/model"
	"github.com/bc-class/service"
	"github.com/bc-class/utils"
)

func DeletePod(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pod := params.ByName("name")

	err := service.DeletePod(pod)
	if err != nil {
		utils.RespMsg(w, r, err)
		return
	}

	utils.RespMsg(w, r, &model.CommonResp{
		Code:    utils.OK,
		Message: "Delete successful",
	})
}

func ScaleDeployment(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	deployment, amount := params.ByName("deployment"), params.ByName("amount")
	a, _ := strconv.Atoi(amount)

	err := service.ScaleDeployment(deployment, int32(a))
	if err != nil {
		utils.RespMsg(w, r, err)
		return
	}

	utils.RespMsg(w, r, &model.CommonResp{
		Code:    utils.OK,
		Message: "Scale successful",
	})
}

func ListDeployment(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	deployments, err := service.ListDeployment()
	if err != nil {
		utils.RespMsg(w, r, err)
		return
	}

	utils.RespMsg(w, r, &model.ListDeploymentResp{
		Code: utils.OK,
		Data: deployments,
	})
}

func ListPods(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	deployment := params.ByName("deployment")

	pods, err := service.ListPods(deployment)
	if err != nil {
		utils.RespMsg(w, r, err)
		return
	}

	utils.RespMsg(w, r, &model.ListPodResp{
		Code: utils.OK,
		Data: pods,
	})
}
