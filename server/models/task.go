package models

import (
	"github.com/james-woo/wakeus/server/utils"
	"github.com/jinzhu/gorm"
	"log"
)

type Task struct {
	gorm.Model
	Commands string `json:"commands"`
	UserId uint `json:"user_id"`
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

func GetTask(user uint, id uint) *Task {
	task := &Task{}
	err := GetDB().Table("tasks").Where("user_id = ? and id = ?", user, id).First(task).Error
	if err != nil {
		log.Println(err)
		return nil
	}
	return task
}

func GetTasks(user uint) []*Task {
	tasks := make([]*Task, 0)
	err := GetDB().Table("tasks").Where("user_id = ?", user).Find(&tasks).Error
	if err != nil {
		log.Println(err)
		return nil
	}
	return tasks
}

func GetAllTasks() []*Task {
	tasks := make([]*Task, 0)
	err := GetDB().Table("tasks").Find(&tasks).Error
	if err != nil {
		log.Println(err)
		return nil
	}
	return tasks
}