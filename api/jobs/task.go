package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/james-woo/wakeus/api/models"
	"github.com/james-woo/wakeus/api/rpc"
	"gopkg.in/robfig/cron.v3"
	"log"
	"time"
)

var TaskJobs = make(map[uint]cron.EntryID)

var LaunchJobs = func(ctx context.Context) {
	tasks := models.GetTasks()

	for _, task := range tasks {
		t := task
		fmt.Printf("Adding task: %s %s\n", t.Type, t.Data)
		AddJob(ctx, t)
	}
}

var AddJob = func(ctx context.Context, task *models.Task) {
	c := cron.New()
	fmt.Printf("Adding task: %s %s\n", task.Type, task.Data)
	switch task.Type {
	case "basic":
		{
			id, err := c.AddFunc(task.Schedule, func() {
				fmt.Printf("%s: Performing hardware basic with data: %s, schedule: %s\n", time.Now(), task.Data, task.Schedule)
				var basic= models.Basic{}
				if err := json.Unmarshal([]byte(task.Data), &basic); err != nil {
					log.Fatal(err)
				}
				rpc.PerformBasic(
					ctx,
					rpc.Color{R: basic.Color.R, G: basic.Color.G, B: basic.Color.B},
					basic.Intensity,
				)
			})
			if err != nil {
				log.Fatal(err)
			}
			TaskJobs[task.ID] = id
		}
	case "fade":
		{
			id, err := c.AddFunc(task.Schedule, func() {
				fmt.Printf("%s: Performing hardware fade with data: %s, schedule: %s\n", time.Now(), task.Data, task.Schedule)
				var fade= models.Fade{}
				if err := json.Unmarshal([]byte(task.Data), &fade); err != nil {
					log.Fatal(err)
				}
				rpc.PerformFade(
					ctx,
					rpc.Color{R: fade.StartColor.R, G: fade.StartColor.G, B: fade.StartColor.B},
					rpc.Color{R: fade.EndColor.R, G: fade.EndColor.G, B: fade.EndColor.B},
					fade.StartIntensity,
					fade.EndIntensity,
					fade.Duration,
				)
			})
			if err != nil {
				log.Fatal(err)
			}
			TaskJobs[task.ID] = id
		}
	case "clear":
		{
			id, err := c.AddFunc(task.Schedule, func() {
				fmt.Printf("%s: Performing hardware clear\n", time.Now())
				rpc.PerformClear(ctx)
			})
			if err != nil {
				log.Fatal(err)
			}
			TaskJobs[task.ID] = id
		}
	}
	c.Start()
}

var DeleteJob = func(taskId uint) {
	c := cron.New()
	c.Remove(TaskJobs[taskId])
}