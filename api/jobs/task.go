package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/james-woo/wakeus/api/models"
	"github.com/james-woo/wakeus/api/rpc"
	"github.com/robfig/cron"
	"log"
	"time"
)

var LaunchTasks = func(ctx context.Context) {
	c := cron.New()
	tasks := models.GetTasks()

	for _, task := range tasks {
		fmt.Printf("Adding task: %s %s\n", task.Type, task.Data)
		switch task.Type {
		case "basic":
			{
				err := c.AddFunc(task.Schedule, func() {
					fmt.Printf("Performing hardware basic with data: %s\n", task.Data)
					var basic = models.Basic{}
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
			}
		case "fade":
			{
				err := c.AddFunc(task.Schedule, func() {
					fmt.Printf("Performing hardware fade with data: %s\n", task.Data)
					var fade = models.Fade{}
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
			}
		}
	}

	for _, entry := range c.Entries() {
		fmt.Printf("Job: %s, Schedule: %s\n", entry.Job, entry.Schedule.Next(time.Now()))
	}

	c.Start()
}
