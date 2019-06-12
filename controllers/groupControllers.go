package controllers

import (
	"encoding/json"
	"net/http"
	"webstuff/go-contacts/models"
	u "webstuff/go-contacts/utils"
)

var CreateGroup = func(w http.ResponseWriter, r *http.Request) {

	group := &models.GroupPayment{}
	err := json.NewDecoder(r.Body).Decode(group) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := group.CreateGroup() //Create account
	u.Respond(w, resp)
}

var AuthenticateGroup = func(w http.ResponseWriter, r *http.Request) {

	group := &models.GroupPayment{}
	err := json.NewDecoder(r.Body).Decode(group) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.LoginGroup(group.Email, group.Password)
	u.Respond(w, resp)
}
