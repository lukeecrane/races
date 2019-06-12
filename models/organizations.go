package models

//a struct to organization
type Organization struct {
	Name                        string
	LogoUrl                     string //Logo of Race
	CustomLayoutRegisterPageUrl string //Pages that include custom CSS and basic Layout for Page
	CustomLayoutAdminPageUrl    string //Custom headers and footers for Admin stuff
	CustomLayoutEmailUrl        string //Custom headers and footers for emails
	Races                       []Race
	Email                       string `json:"email"`
	PhoneNumber                 string
	StreetAddress               string
	City                        string
	ZipCode                     string
	State                       string
	Country                     string
}
