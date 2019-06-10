package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/james-woo/wakeus/server/models"
	"github.com/james-woo/wakeus/server/utils"
	"net/http"
	"strconv"
)

var CreateTask = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	task := &models.Task{}

	err := json.NewDecoder(r.Body).Decode(task)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	task.UserId = user
	resp := task.Create()
	utils.Respond(w, resp)
}

var DeleteTask = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := r.Context().Value("user").(uint)
	taskId, err := strconv.Atoi(params["task_id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error in your request, missing task_id"))
		return
	}
	task := models.GetTask(uint(userId), uint(taskId))
	resp := task.Delete()
	utils.Respond(w, resp)
}

var GetTasks = func(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user").(uint)

	data := models.GetTasks(uint(userId))
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}

var GetTask = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := r.Context().Value("user").(uint)
	taskId, err := strconv.Atoi(params["task_id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error in your request, missing task_id"))
		return
	}

	data := models.GetTask(uint(userId), uint(taskId))
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}

var UpdateTask = func(w http.ResponseWriter, r *http.Request) {

}