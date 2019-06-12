package controllers

import (
	"encoding/json"
	"net/http"
	"webstuff/go-contacts/models"
	u "webstuff/go-contacts/utils"
)

var CreateParticipant = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("group").(uint) //Grab the id of the user that send the request
	participant := &models.Participant{}

	err := json.NewDecoder(r.Body).Decode(participant)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	participant.GroupId = user
	resp := participant.Create()
	u.Respond(w, resp)
}

var GetParticipant = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetContacts(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var UpdateParticipant = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetContacts(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var DeleteParticipant = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetContacts(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
