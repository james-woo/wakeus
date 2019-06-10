package controllers

import (
	"encoding/json"
	"github.com/james-woo/wakeus/server/models"
	"github.com/james-woo/wakeus/server/utils"
	"net/http"
)

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) // Decode the request body into struct and failed if any error occur
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(account.Username, account.Password)
	utils.Respond(w, resp)
}