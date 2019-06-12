package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
	u "webstuff/go-contacts/utils"
)

/*
JWT claims struct
*/
type GroupToken struct {
	GroupId uint
	jwt.StandardClaims
}

type GroupPayment struct {
	gorm.Model
	Payments                  []GroupParticipantPayment
	RaceID                    uint   //Information concerning the race that the group is attending
	Email                     string `json:"email"`         //Group Email to send receipt to
	Paid                      bool   `gorm:"default:false"` //Has every member of the group paid
	Password                  string `json:"password"`      //Password and Email can be used to log on
	CreditCardNumber          string //Incase you would like store credit cards
	CreditCardExpirationMonth uint8  //1,2..12
	CreditCardExpirationYear  uint16 //2012...
	CreditCardSecurityCode    string //Hopefully I will not store this information
	CreditCardNameOnCard      string //Name on Card
	GroupToken                string `json:"token";sql:"-"`
}

//Validate incoming user details...
func (group *GroupPayment) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(group.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(group.Password) < 6 {
		return u.Message(false, "Password is must be 6 characters long"), false
	} else if len(group.Password) < 1 {
		return u.Message(false, "Password is required"), false
	}

	//Email must be unique
	temp := &GroupPayment{}

	//check for errors and duplicate emails
	err := GetDB().Table("grouppayments").Where("email = ?", group.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (group *GroupPayment) CreateGroup() map[string]interface{} {

	if resp, ok := group.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(group.Password), bcrypt.DefaultCost)
	group.Password = string(hashedPassword)

	GetDB().Create(group)

	if group.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}

	//Create new JWT token for the newly registered account
	tk := &GroupToken{GroupId: group.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	group.GroupToken = tokenString

	group.Password = "" //delete password

	response := u.Message(true, "Account has been created")
	response["account"] = group
	return response
}

func LoginGroup(email, password string) map[string]interface{} {

	group := &GroupPayment{}
	err := GetDB().Table("grouppayments").Where("email = ?", email).First(group).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(group.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again")
	}
	//Worked! Logged In
	group.Password = ""

	//Create JWT token
	tk := &GroupToken{GroupId: group.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	group.GroupToken = tokenString //Store the token in the response

	resp := u.Message(true, "Logged In")
	resp["account"] = group
	return resp
}

func GetGroup(u uint) *GroupPayment {

	group := &GroupPayment{}
	GetDB().Table("accounts").Where("id = ?", u).First(group)
	if group.Email == "" { //User not found!
		return nil
	}

	group.Password = ""
	return group
}
