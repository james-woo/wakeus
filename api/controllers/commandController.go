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
	rpc.PerformBasic(
		r.Context(),
		models.Color{R: basic.Color.R, G: basic.Color.G, B: basic.Color.B},
		basic.Intensity,
	)
	resp := utils.Message(true, "success")
	utils.Respond(w, resp)
}

var Fade = func(w http.ResponseWriter, r *http.Request) {
	fade := &models.Fade{}
	err := json.NewDecoder(r.Body).Decode(fade)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	fmt.Printf("%s: Performing hardware fade with data: %+v\n", time.Now(), fade)
	rpc.PerformFade(
		r.Context(),
		models.Color{R: fade.StartColor.R, G: fade.StartColor.G, B: fade.StartColor.B},
		models.Color{R: fade.EndColor.R, G: fade.EndColor.G, B: fade.EndColor.B},
		fade.StartIntensity,
		fade.EndIntensity,
		fade.Duration,
	)
	resp := utils.Message(true, "success")
	utils.Respond(w, resp)
}

var Rainbow = func(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s: Performing hardware rainbow\n", time.Now())
	rpc.PerformRainbow(r.Context())
	resp := utils.Message(true, "success")
	utils.Respond(w, resp)
}

var Clear = func(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s: Performing hardware clear\n", time.Now())
	rpc.PerformClear(r.Context())
	resp := utils.Message(true, "success")
	utils.Respond(w, resp)
}
