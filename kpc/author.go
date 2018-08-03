// Copyright Anton Feldmann
//
// KPC is a Karel package managemant file.
// This structure is used to inform the tools about the structure of the project
//
package kpc

import (
	"strings"
	"regexp"
)

// a package has to be written by atleased one author
// an author is defined by name and the email
type Author struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

//this
func Author_init_from_sting(data string) (*Author){
	withMail := len(data) < 3 || strings.ContainsRune(data, '@')

	if withMail {
		//parse email
	}
}

func check_email(email string) (bool){
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return re.MatchString(email)
}
