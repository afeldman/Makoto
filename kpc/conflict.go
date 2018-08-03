// Copyright Anton Feldmann
//
// KPC is a Karel package managemant file.
// This structure is used to inform the tools about the structure of the project
//
package kpc

// different projects can contain different conflicts on the given requirements
//
// a conflict is targeted with a project name and one or different version of this project
type Conflict struct {
	Name     string   `yaml:"name"`
	Versions []string `yaml:"version"`
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
	return &Conflict{
		Name: name,
		Versions: []string{version},
	}
}

// add an version number
//
// another version can be added to the version set.
//
// @param {string} version the version to add to the list of conflicts
func (this *Conflict)AddElement(version string){
	if !this.ContainsVersion(version) {
		this.Versions = append(this.Versions, version)
	}
}

// the number of versions for that project that builds into conflicts
//
// @return {int} the number of known versions for a conflict
func (this *Conflict)GetLength() (int) {
	return len(this.Versions)
}

// a small private library. Chacking if a version is already in the set
//
// @param {string} version. the version to request
//
// @return {bool} true if the element is in the list otherwise false
// @return {int}  -1 if the element is not in the list otherwise the index in the set
func (this *Conflict)find_elem(version string) (bool, int){
	for i := range this.Versions {
		if this.Versions[i] == version {
			return true, i
		}
	}

	return false, -1
}

// change a version, because the error is not happen anymore or the version string is wrong
//
// @param {string} old the old number
// @param {string} new the replaycement
//
// @return {bool} the version is changed
func (this *Conflict)ChangeVersion(old, new string) (bool){
	has_elem, idx := this.find_elem(old)
	if  has_elem{
		this.Versions[idx] = new
		return true;
	}
	return false
}

// contains the version
// is the requested version into the set of versions
//
// @param {string} the version to check
//
// @return {bool} true if the version is in the list else false
func (this *Conflict)ContainsVersion(version string) (bool){
	has_elem, _ := this.find_elem(version)
	return has_elem;
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
	has_elem, idx := this.find_elem(version)
	if has_elem && this.GetLength() > 1 {
		this.Versions[idx] = ""
//resize the array of versions
		var r []string

		for _, str := range this.Versions{
			if str != ""{
				r = append(r,str)
			}
		}

		this.Versions = r
	}
}
