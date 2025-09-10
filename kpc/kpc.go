// Copyright Anton Feldmann
//
// KPC is a Karel package management structure.
// This structure is used to inform the tools about the structure of the project
package kpc

import "fmt"

// KpcVersion changed the KPC Version. now we use TOML
const KpcVersion = "0.3.0"

// kpc datastructurwe
type KPC struct {
	KPCVersion string `toml:"kpc_version"`

	// --- Package Information ---
	Name         string        `toml:"name"`
	Description  string        `toml:"description"`
	Version      string        `toml:"version"`
	Homepage     string        `toml:"url,omitempty"` // Projekt-Homepage
	Requirements []Requirement `toml:"requirements,omitempty"`
	Conflicts    []Conflict    `toml:"conflicts,omitempty"`
	Authors      []Author      `toml:"authors,omitempty"`
	Source       Repository    `toml:"source,omitempty"` // Git-Repo des Projekts
	Issue        string        `toml:"issues,omitempty"`
	Keywords     []string      `toml:"keywords,omitempty"`

	// --- Path Settings ---
	Prefix     string `toml:"prefix"`
	SrcDir     string `toml:"srcdir,omitempty"`
	TypeDir    string `toml:"typedir,omitempty"`
	IncludeDir string `toml:"includedir,omitempty"`
	ConstDir   string `toml:"constdir,omitempty"`
	FormDir    string `toml:"formdir,omitempty"`
	DictDir    string `toml:"dictdir,omitempty"`

	// --- File Includes ---
	Main     string   `toml:"main"`
	Dicts    []string `toml:"dicts,omitempty"`
	Forms    []string `toml:"forms,omitempty"`
	Types    []string `toml:"types,omitempty"`
	Includes []string `toml:"includes,omitempty"`
	Consts   []string `toml:"consts,omitempty"`
}

// initilize a KPC object
func InitKPC(name string) *KPC {
	return &KPC{
		KPCVersion:   KpcVersion,
		Name:         name,
		Description:  "",
		Version:      "",
		Homepage:     "",
		Source:       *(InitRepository()),
		Issue:        "",
		Prefix:       "",
		SrcDir:       "",
		TypeDir:      "",
		IncludeDir:   "",
		ConstDir:     "",
		Main:         "",
		Requirements: []Requirement{},
		Conflicts:    []Conflict{},
		Authors:      []Author{},
		Dicts:        []string{},
		Includes:     []string{},
		Forms:        []string{},
		Types:        []string{},
		Consts:       []string{},
		Keywords:     []string{},
	}
}

// return KPC version version
//
// @return (string) version
func (kpc_obj *KPC) GetKPCVersion() string {
	return kpc_obj.KPCVersion
}

// RejectConflict you did well and fixed an issue. then delete the conflict
//
// @param (string) name of conflivt package
func (kpc_obj *KPC) RejectConflict(name string) {
	var tmp []Conflict
	for _, v := range kpc_obj.Conflicts {
		if v.Name == name {
			continue
		} else {
			tmp = append(tmp, v)
		}
	}
	kpc_obj.Conflicts = tmp
}

// ConflictSize how many conflicts are known
func (kpc_obj *KPC) ConflictSize() int {
	return len(kpc_obj.Conflicts)
}

// AddConflict you found a new conflict so add the conflict to the list
func (kpc_obj *KPC) AddConflict(con Conflict) {
	kpc_obj.Conflicts = append(kpc_obj.Conflicts, con)
}

// GetConflicts get conflicts
func (kpc_obj *KPC) GetConflicts() []Conflict {
	return kpc_obj.Conflicts
}

// GetConflict get a specific conflict
func (kpc_obj *KPC) GetConflict(name string) *Conflict {
	for _, val := range kpc_obj.Conflicts {
		if val.Name == name {
			return &val
		}
	}
	return nil
}

// GetAuthor is the author structure of the KPC struct
//
// @params (strung) the authors name
//
// @return (*Author) the author structure
// @return (error) if the author is unknown throws error
func (kpc_obj *KPC) GetAuthor(authorsName string) (*Author, error) {
	for _, val := range kpc_obj.Authors {
		if val.Name == authorsName {
			return &val, nil
		}
	}
	return nil, fmt.Errorf("can not find author %s", authorsName)
}

// AddAuthor add a new Author to the list
//
// @param (Author) an author struct
func (kpc_obj *KPC) AddAuthor(aut Author) {
	_, err := kpc_obj.GetAuthor(*aut.GetAuthorName())
	if err != nil {
		// author do not exists in list
		kpc_obj.Authors = append(kpc_obj.Authors, aut)
	}
}

// NumberOfAuthors is the number of authors in the author list
//
// @return (int) the number of authors in kpc
func (kpc_obj *KPC) NumberOfAuthors() int {
	return len(kpc_obj.Authors)
}

// RejectAuthor deletes an author in the list of authors
//
// @param (string) authors name
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

// get all rewuirements
func (kpc_obj *KPC) GetRequirement(name string) *Requirement {
	for _, val := range kpc_obj.Requirements {
		if val.Name == name {
			return &val
		}
	}
	return nil
}

// add a requirement
func (kpc_obj *KPC) AddRequirement(req Requirement) {
	kpc_obj.Requirements = append(kpc_obj.Requirements, req)
}

// how many requirements are stored
func (kpc_obj *KPC) RequirementSize() int {
	return len(kpc_obj.Requirements)
}

// delete a requirement
func (kpc_obj *KPC) RejectRequirement(name string) {
	var tmp []Requirement
	for _, v := range kpc_obj.Requirements {
		if v.Name == name {
			continue
		} else {
			tmp = append(tmp, v)
		}
	}
	kpc_obj.Requirements = tmp
}
