package models

import (
	"github.com/jinzhu/gorm"
	"time"
	u "webstuff/go-contacts/utils"
)

//a struct to rep user account
type Participant struct {
	gorm.Model
	FirstName     string
	LastName      string
	Email         string `json:"email"`
	PhoneNumber   string
	StreetAddress string
	City          string
	ZipCode       string
	State         string
	Country       string
	Age           uint8  `gorm:"default:0"`
	Gender        string `gorm:"default:'Female'"`
	GroupToken    string `json:"token";sql:"-"`
	GroupId       uint   `json:"group_id"` //The user that this contact belongs to
}

func (participant *Participant) Create() map[string]interface{} {

	if resp, ok := participant.Validate(); !ok {
		return resp
	}

	GetDB().Create(participant)

	resp := u.Message(true, "success")
	resp["participant"] = participant
	return resp
}

func (participant *Participant) Validate() (map[string]interface{}, bool) {

	if participant.FirstName == "" {
		return u.Message(false, "First Name is empty"), false
	}

	if participant.PhoneNumber == "" {
		return u.Message(false, "Phone Number is empty"), false
	}
	if !phoneValid.MatchString(participant.PhoneNumber) {
		return u.Message(false, "Phone Number is invalid"), false
	}

	if participant.GroupId <= 0 {
		return u.Message(false, "Group is not valid"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}
