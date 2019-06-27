package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/james-woo/wakeus/api/models"
	"github.com/james-woo/wakeus/api/rpc"
	"github.com/james-woo/wakeus/api/utils"
	"net/http"
	"time"
)

var Basic = func(w http.ResponseWriter, r *http.Request) {
	var basic = models.Basic{}
	err := json.NewDecoder(r.Body).Decode(&basic)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	fmt.Printf("%s: Performing hardware basic with data: %+v\n", time.Now(), basic)
	result, err := rpc.PerformBasic(
		r.Context(),
		models.Color{R: basic.Color.R, G: basic.Color.G, B: basic.Color.B},
		basic.Intensity,
	)
	if result {
		resp := utils.Message(true, "success")
		utils.Respond(w, resp)
	} else {
		message := fmt.Sprintf("Failed to perform basic with request: %+v. Error: %s", basic, err)
		utils.Respond(w, utils.Message(false, message))
	}
}

var Fade = func(w http.ResponseWriter, r *http.Request) {
	fade := &models.Fade{}
	err := json.NewDecoder(r.Body).Decode(fade)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	fmt.Printf("%s: Performing hardware fade with data: %+v\n", time.Now(), fade)
	result, err := rpc.PerformFade(
		r.Context(),
		models.Color{R: fade.StartColor.R, G: fade.StartColor.G, B: fade.StartColor.B},
		models.Color{R: fade.EndColor.R, G: fade.EndColor.G, B: fade.EndColor.B},
		fade.StartIntensity,
		fade.EndIntensity,
		fade.Duration,
	)
	if result {
		resp := utils.Message(true, "success")
		utils.Respond(w, resp)
	} else {
		message := fmt.Sprintf("Failed to perform fade with request: %+v. Error: %s", fade, err)
		utils.Respond(w, utils.Message(false, message))
	}
}

var Rainbow = func(w http.ResponseWriter, r *http.Request) {
	rainbow := &models.Rainbow{}
	err := json.NewDecoder(r.Body).Decode(rainbow)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}
	fmt.Printf("%s: Performing hardware rainbow\n", time.Now())
	result, err := rpc.PerformRainbow(
		r.Context(),
		rainbow.Cycles,
	)
	if result {
		resp := utils.Message(true, "success")
		utils.Respond(w, resp)
	} else {
		message := fmt.Sprintf("Failed to perform rainbow. Error: %s", err)
		utils.Respond(w, utils.Message(false, message))
	}
}

var Clear = func(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s: Performing hardware clear\n", time.Now())
	result, err := rpc.PerformClear(r.Context())
	if result {
		resp := utils.Message(true, "success")
		utils.Respond(w, resp)
	} else {
		message := fmt.Sprintf("Failed to perform clear. Error: %s", err)
		utils.Respond(w, utils.Message(false, message))
	}
}
