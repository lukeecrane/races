package models

import (
	"github.com/jinzhu/gorm"
	"time"
	u "webstuff/go-contacts/utils"
)

type GroupParticipantPayment struct {
	gorm.Model
	ParticipantId       uint       //Location to get permanent data
	UnDiscountedTotal   float32    `grom:"default:0.0"` //Amount with out a coupon and early registration
	DateDiscountedTotal float32    `grom:"default:0.0"` //Total Amount because of early registration
	TotalDue            float32    `grom:"default:0.0"` //Actual amount due after applying the two possible discounts (early and coupon)
	CouponId            uint       //Need to find out if a coupon was used, possible
	CourseId            uint       //What course did they sign up for
	paid                bool       `gorm:"default:false"` //Have they paid
	PaidDate            *time.Time //Date at which they paid should be same as registration Date
	CheckedInDate       *time.Time //Date of checkin
	CheckedIn           bool       `gorm:"default:false"` //Have the checkin and got their race packet
	RegistrationDate    *time.Time //Date of completed registration
	SignedWaiver        bool       `gorm:"default:false"` //Did that sign liability waiver
	methodOfPayment     string     `gorm:"default:'CC'"`  //Credit Card, Check, Cash  should be group level but isn't just in case
	locationOfPayment   string     `gorm:"default:'web'"` //In person, online, phone
	note                string     //List of instruction about this person
	GroupPaymentId      uint       //A payment Group
}

func (participantpayment *GroupParticipantPayment) Create() map[string]interface{} {

	if resp, ok := participantpayment.Validate(); !ok {
		return resp
	}

	GetDB().Create(participantpayment)

	resp := u.Message(true, "success")
	resp["participant payment"] = participantpayment
	return resp
}

func (participantpayment *GroupParticipantPayment) Validate() (map[string]interface{}, bool) {

	//if participantpayment.FirstName == "" {
	//	return u.Message(false, "First Name is empty"), false
	//}

	//if participantpayment.PhoneNumber == "" {
	//	return u.Message(false, "Phone Number is empty"), false
	//}
	//if !phoneValid.MatchString(participantpayment.PhoneNumber) {
	//	return u.Message(false, "Phone Number is invalid"), false
	//}

	//if participantpayment.GroupId <= 0 {
	//	return u.Message(false, "Group is not valid"), false
	//}

	//All the required parameters are present
	return u.Message(true, "success"), true
}
