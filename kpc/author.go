// Copyright Anton Feldmann
//
// KPC is a Karel package managemant file.
// This structure is used to inform the tools about the structure of the project
//
package kpc

// a package has to be written by atleased one author
// an author is defined by name and the email
type Author struct {
	Name  string `toml:"name"`
	Email string `toml:"email"`
}

// an author needs atleased a name
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

// set a new email address
// each author has a empty email address
//
// this function is pagt of the Author structure
//
// @param {string} email. The email address for the author
func (auth *Author) SetEmail(email string) {
	auth.Email = email
}

// set a new name.
// each author has a name. Maybe if the name was written wrong.
// This method change the name
//
// this function is pagt of the Author structure
//
// @param {string} name. the new name for the author
func (auth *Author) SetName(name string) {
	auth.Name = name
}

// get email of author
// get the author's email
//
// this function is pagt of the Author structure
//
// @return {*string} the Email address
func (auth *Author) GetEmail() *string {
	return &(auth.Email)
}

// get name of author
// get the author's name
//
// this function is pagt of the Author structure
//
// @return {*string} the Author's name
func (auth *Author) GetName() *string {
	return &(auth.Name)
}

// the the author
func (kpc_obj *KPC) GetAuthor(name string) *Author {
	for _, val := range kpc_obj.Authors {
		if val.Name == name {
			return &val
		}
	}
	return nil
}

// add an author to the list
func (kpc_obj *KPC) AddAuthor(aut Author) {
	kpc_obj.Authors = append(kpc_obj.Authors, aut)
}

// how many authors are working on this package
func (kpc_obj *KPC) AuthorsSize() int {
	return len(kpc_obj.Authors)
}

// delete a author from the package
func (kpc_obj *KPC) RejectAuthor(name string) {
	var tmp []Author
	for _, v := range kpc_obj.Authors {
		if v.Name == name {
			continue
		} else {
			tmp = append(tmp, v)
		}
	}
	kpc_obj.Authors = tmp
}
