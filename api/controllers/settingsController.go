package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/james-woo/wakeus/api/models"
	"github.com/james-woo/wakeus/api/utils"
)

var UpdateSettings = func(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s: Update setting\n", time.Now())
	setting := &models.Setting{}
	err := json.NewDecoder(r.Body).Decode(setting)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	models.GetSettings().Pause = setting.Pause
}

var Settings = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetSettings()
	resp := utils.Message(true, "success")
	resp["settings"] = data
	utils.Respond(w, resp)
}
