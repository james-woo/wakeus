package models

import (
	"github.com/james-woo/wakeus/api/utils"
	"github.com/jinzhu/gorm"
	"log"
)

type Task struct {
	gorm.Model
	Schedule string `json:"schedule"`
	Type string `json:"type"`
	Data string `json:"data"`
}

func (task *Task) Create() map[string]interface{} {
	GetDB().Create(task)
	resp := utils.Message(true, "success")
	resp["task"] = task
	return resp
}

func (task *Task) Delete() map[string]interface{} {
	GetDB().Delete(task)
	resp := utils.Message(true, "success")
	resp["task"] = task
	return resp
}

func GetTask(id uint) *Task {
	task := &Task{}
	err := GetDB().Table("tasks").Where("id = ?", id).First(task).Error
	if err != nil {
		log.Println(err)
		return nil
	}
	return task
}

func GetTasks() []*Task {
	tasks := make([]*Task, 0)
	err := GetDB().Table("tasks").Find(&tasks).Error
	if err != nil {
		log.Println(err)
		return nil
	}
	return tasks
}
