package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/james-woo/wakeus/api/models"
	"github.com/james-woo/wakeus/api/utils"
	"net/http"
	"strconv"
)

var CreateTask = func(w http.ResponseWriter, r *http.Request) {
	task := &models.Task{}
	err := json.NewDecoder(r.Body).Decode(task)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}
	resp := task.Create()
	utils.Respond(w, resp)
}

var DeleteTask = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskId, err := strconv.Atoi(params["task_id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error in your request, missing task_id"))
		return
	}
	task := models.GetTask(uint(taskId))
	resp := task.Delete()
	utils.Respond(w, resp)
}

var GetTasks = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetTasks()
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}

var GetTask = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskId, err := strconv.Atoi(params["task_id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error in your request, missing task_id"))
		return
	}
	data := models.GetTask(uint(taskId))
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}

var UpdateTask = func(w http.ResponseWriter, r *http.Request) {

}