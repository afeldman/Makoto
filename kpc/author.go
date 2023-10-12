package kpc

// Copyright Anton Feldmann
//
// KPC is a Karel package managemant file.
// This structure is used to inform the tools about the structure of the project
//

import (
	"github.com/asaskevich/govalidator"
	log "github.com/sirupsen/logrus"
)

// Author is a structure with name and email
// address of the user
// an author is defined by name and the email
type Author struct {
	Name  string `toml:"name"`
	Email string `toml:"email"`
}

// InitAuthor creates and author
//
// the email is set to an empty string.
//
// @param {string} name the author name
//
// @return {*Author} the author structure
func InitAuthor(name string) *Author {
	return &Author{
		Name:  name,
		Email: ""}
}

// SetAuthorEmail change the email address
// each author has a empty email address
//
// this function is pagt of the Author structure
//
// @param {string} email. The email address for the author
func (auth *Author) SetAuthorEmail(email string) {
	if govalidator.IsEmail(email) {
		auth.Email = email
	} else {
		log.Errorf("email %s is not a valid mail address", email)
	}
}

// GetAuthorEmail returns the authors email
// get the author's email
//
// this function is pagt of the Author structure
//
// @return {*string} the Email address
func (auth *Author) GetAuthorEmail() *string {
	return &auth.Email
}

// SetAuthorName changes the name of the author
// each author has a name. Maybe if the name was written wrong.
// This method change the name
//
// this function is pagt of the Author structure
//
// @param {string} name. the new name for the author
func (auth *Author) SetAuthorName(name string) {
	auth.Name = name
}

// GetAuthorName of the author
// get the author's name
//
// this function is pagt of the Author structure
//
// @return {*string} the Author's name
func (auth *Author) GetAuthorName() *string {
	return &(auth.Name)
}
