package models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"
)

//a struct to rep user account
type Race struct {
	gorm.Model
	startOfEarlyReg         *time.Time
	PeriodPlusIncreases     []PeriodPlusIncrease //Gives you a range of increases
	Courses                 []Course
	Coupons                 []Coupon //Number of individual race options
	GroupsOfItems           []GroupOfItems
	price                   float32
	Email                   string `json:"email"`
	Password                string `json:"password"`
	Token                   string `json:"token";sql:"-"`
	ConfigurationRegisterID uint   //One to
	ConfigurationAdminID    uint
}

type PeriodPlusIncrease struct {
	gorm.Model
	startTime       *time.Time    //Start of this course
	endTime         *time.Time    //Start of this course
	percentIncrease sql.NullInt64 `json:"increase",gorm:"default:10"` //Percent of increase until 10%
	RaceID          uint
}

type Course struct { //Copy these from the year before
	gorm.Model
	Length            float32       // km or miles
	UnitOfMeasurement string        `gorm:"default:'miles'"` //miles, kilometers, feet, centimeters
	Title             string        // title
	ImageUrl          string        //A image of the course
	Description       string        //Description of the course
	Instructions      string        //instruction to racers of this course
	AgeStartRange     sql.NullInt64 `gorm:"default:0"`     //Default 0
	AgeEndRange       sql.NullInt64 `gorm:"default:150"`   //Default 150
	GendersAllowed    string        `gorm:"default:'All'"` // List of genders: All, Male, Female
	StartTime         *time.Time    //Start of this course
	RaceID            uint
}

type Coupon struct {
	gorm.Model
	Name             string          //Name of organization or what ever
	DisplayGroupName bool            //Instead of being just a coupon display this groups name and ask for validation code
	PercentOff       sql.NullFloat64 `gorm:"default:10"` // 25% off of current prices
	Code             string          //Entry code needed to get said discount or just to validate your part of this group
	ExpirationDate   *time.Time      //coupon expires
	RaceID           uint
}

//Merchandised Free or otherwise
type GroupOfItems struct {
	gorm.Model
	Name   string //Name of stuff
	Items  []Item
	RaceID uint
}

type Item struct {
	gorm.Model
	Description string        //Women large size
	ItemOrder   sql.NullInt64 `gorm:"default:0"` //Order of displaying this item
	//Sizes string  `gorm:"default:'YS,YM,YL,XS,S,M,L,XL,XXL,3XL'"` //Array might be hard to do so comma dilimated string YS,YM,YL,XS,S,M,L,XL,XXL,3XL
	Cost           sql.NullFloat64 `gorm:"default:0"` //Zero considering free item
	Size           string          `gorm:"default:'M'"`
	GroupOfItemsID uint
}

type ConfigurationRegister struct {
	gorm.Model
	Races []Race
	//Not sure what will be here yet
}
type ConfigurationAdmin struct {
	gorm.Model
	Races []Race
	//Not sure what will be here yet
}
