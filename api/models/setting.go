package models

import (
	"sync"
)

type Setting struct {
	Pause bool
}

var instance *Setting
var once sync.Once

func GetSettings() *Setting {
	once.Do(func() {
		instance = &Setting{
			Pause: false,
		}
	})
	return instance
}
