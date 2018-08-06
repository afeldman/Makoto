// Copyright Anton Feldmann
//
// KPC is a Karel package managemant file.
// This structure is used to inform the tools about the structure of the project
//
package kpc

import (
	"github.com/emirpasic/gods/sets/hashset"
)

// different projects can contain different conflicts on the given requirements
//
// a conflict is targeted with a project name and one or different version of this project
type Conflict struct {
	Name     string  `json:"name"`
	Versions *hashset.Set `json:"version,omitempty"`
}

// building a conflict information
//
// a conflict has to start with a name and one version.
// if there is non name then there is not conflict and
// if there is not a start version, then there is no conflict
//
// @param {string} name the name of the conflicting project
// @param {string} version the version of the conflict
//
// @return {*Conflict} the conflict information
func Conflict_Init(name, version string) (*Conflict){
	m := hashset.New()
	m.Add(version)

	return &Conflict{
		Name: name,
		Versions: m,
	}
}

// add an version number
//
// another version can be added to the version set.
//
// @param {string} version the version to add to the list of conflicts
func (this *Conflict)AddElement(version string){
	this.Versions.Add(version)
}

// the number of versions for that project that builds into conflicts
//
// @return {int} the number of known versions for a conflict
func (this *Conflict)GetLength() (int) {
	return this.Versions.Size()
}

// change a version, because the error is not happen anymore or the version string is wrong
//
// @param {string} old the old number
// @param {string} new the replaycement
//
// @return {bool} the version is changed
func (this *Conflict)ChangeVersion(old, new string) (bool){
	b := this.Versions.Contains(old)

	if b {
		this.Versions.Remove(old)
		this.Versions.Add(new)
		return true
	} else {
		return false
	}
}

// contains the version
// is the requested version into the set of versions
//
// @param {string} the version to check
//
// @return {bool} true if the version is in the list else false
func (this *Conflict)ContainsVersion(version string) (bool){
	return this.Versions.Contains(version)
}

// get the name of the package that triggers the error
//
// @return {*string} the name of the project
func (this *Conflict)GetName()(*string){
	return &(this.Name)
}

// set the name. Maybe the project name changes or the name is wrong
//
// @param {string} name the new name for the project
func (this *Conflict)SetName(name string){
	this.Name = name
}

// reject a version from the list of conflicts.
// If a version is deleted then the list is resized.
// if the list has only one element, then the element cannot deleted.
// you need atlease one version to have a conflict
//
// @param {string} version the version to delete
func (this *Conflict)RejectVersion(version string){
	this.Versions.Remove(version)
}
