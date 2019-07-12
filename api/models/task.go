package models

import (
	"github.com/jinzhu/gorm"
	"log"
)

type Task struct {
	gorm.Model
	Schedule string `json:"schedule"`
	Type string `json:"type"`
	Data string `json:"data"`
}

func (task *Task) Create() *Task {
	GetDB().Create(task)
	return task
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

func (task *Task) Update(taskId uint, updated Task) *Task {
	GetDB().Model(task).Updates(updated)
	return task
}

func (task *Task) Delete() *Task {
	GetDB().Delete(task)
	return task
}
