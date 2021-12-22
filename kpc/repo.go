// Copyright Anton Feldmann
//
// KPC is a Karel package management structure.
// This structure is used to inform the tools about the structure of the project
package kpc

//Repository is a structure to represend the developmentfiles
//the basic repository protocol is git
//the tag is the hash in the git tree.
type Repository struct {
	URL string `toml:"url"`
	Tag string `toml:"tag"`
}

// init a repo with default values
func InitRepository() *Repository {
	return &Repository{URL: "", Tag: ""}
}

// set a tag
func (repo *Repository) SetTag(tag string) {
	repo.Tag = tag
}

// get a tag
func (repo *Repository) GetTag() *string {
	return &(repo.Tag)
}

// set the url
func (repo *Repository) SetURL(address string) {
	repo.URL = address
}

// get project url
func (repo *Repository) GetURL() *string {
	return &(repo.URL)
}

// get the repository
func (kpc_obj *KPC) GetRepo() *Repository {
	return &(kpc_obj.Repository)
}

/*********************** Setter ***********************/

func (kpc_obj *KPC) AddRepo(repo Repository) {
	kpc_obj.Repository = repo
}
