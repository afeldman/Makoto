package kpc

// Copyright Anton Feldmann
//
// KPC is a Karel package managemant file.
// This structure is used to inform the tools about the structure of the project

import log "github.com/sirupsen/logrus"

// Conflict karel packages can have cinfkicts with order projects.
// to avoid a unknown behavior in your projects make sure to set all projects
// your project can get unwonted intersection with into this list
//
// a conflict is targeted with a project name and one or different version of this project
type Conflict struct {
	Name     string   `toml:"name"`
	Versions []string `toml:"versions"`
}

// InitConflict building a conflict information
//
// a conflict has to start with a name and one version.
// if there is non name then there is not conflict and
// if there is not a start version, then there is no conflict
//
// @param {string} name the name of the conflicting project
// @param {string} version the version of the conflict
//
// @return {*Conflict} the conflict information
func InitConflict(name, version string) *Conflict {
	return &Conflict{
		Name:     name,
		Versions: []string{version},
	}
}

// AddVersion add a version number
//
// another version can be added to the version set.
//
// @param {string} version the version to add to the list of conflicts
func (con *Conflict) AddVersion(version string) {

	if !con.ContainsVersion(version) {
		con.Versions = append(con.Versions, version)
	} else {
		log.Debugf("version %s alredy added", version)
	}
}

// NumberOfConflicts the number of versions for that project that builds into conflicts
//
// @return {int} the number of known versions for a conflict
func (con *Conflict) NumberOfConflicts() int {
	return len(con.Versions)
}

// ChangeVersion change a version, because the error is not happen anymore or the version string is wrong
//
// @param {string} old the old number
// @param {string} new the replaycement
//
// @return {bool} the version is changed
func (con *Conflict) ChangeVersion(old, newversion string) bool {
	for idx, val := range con.Versions {
		if val == old {
			con.Versions[idx] = newversion
			return true
		}
	}
	return false
}

// ContainsVersion contains the version
// is the requested version into the set of versions
//
// @param {string} the version to check
//
// @return {bool} true if the version is in the list else false
func (con *Conflict) ContainsVersion(version string) bool {
	log.Debugf("search for version %s", version)
	for _, val := range con.Versions {
		log.Debugf("got %s from list", val)
		if val == version {
			return true
		}
	}

	return false
}

// GetName get the name of the package that triggers the error
//
// @return {*string} the name of the project
func (con *Conflict) GetConflictName() *string {
	return &con.Name
}

// SetName set the name. Maybe the project name changes or the name is wrong
//
// @param {string} name the new name for the project
func (con *Conflict) SetName(name string) {
	con.Name = name
}

// RejectVersion reject a version from the list of conflicts.
// If a version is deleted then the list is resized.
// if the list has only one element, then the element cannot deleted.
// you need atlease one version to have a conflict
//
// @param {string} version the version to delete
func (con *Conflict) RejectVersion(version string) {
	if len(con.Versions) >= 1 {
		var tmp []string
		for _, v := range con.Versions {
			if v == version {
				continue
			} else {
				tmp = append(tmp, v)
			}
		}
		con.Versions = tmp
	}
}
