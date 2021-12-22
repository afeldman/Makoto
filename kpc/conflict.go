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
	Name     string   `toml:"name"`
	Versions []string `toml:"versions"`
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
func InitConflict(name, version string) *Conflict {
	return &Conflict{
		Name:     name,
		Versions: []string{version},
	}
}

// add an version number
//
// another version can be added to the version set.
//
// @param {string} version the version to add to the list of conflicts
func (con *Conflict) AddElement(version string) {
	con.Versions = append(con.Versions, version)
}

// the number of versions for that project that builds into conflicts
//
// @return {int} the number of known versions for a conflict
func (con *Conflict) GetLength() int {
	return len(con.Versions)
}

// change a version, because the error is not happen anymore or the version string is wrong
//
// @param {string} old the old number
// @param {string} new the replaycement
//
// @return {bool} the version is changed
func (con *Conflict) ChangeVersion(old, new_str string) bool {
	for idx, val := range con.Versions {
		if val == old {
			con.Versions[idx] = new_str
			return true
		}
	}

	return false
}

// contains the version
// is the requested version into the set of versions
//
// @param {string} the version to check
//
// @return {bool} true if the version is in the list else false
func (con *Conflict) ContainsVersion(version string) bool {

	for _, val := range con.Versions {
		if val == version {
			return true
		}
	}

	return false
}

// get the name of the package that triggers the error
//
// @return {*string} the name of the project
func (con *Conflict) GetName() *string {
	return &(con.Name)
}

// set the name. Maybe the project name changes or the name is wrong
//
// @param {string} name the new name for the project
func (con *Conflict) SetName(name string) {
	con.Name = name
}

// reject a version from the list of conflicts.
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

// you did well and fixed an issue. then delete the conflict
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

// how many conflicts are known
func (kpc_obj *KPC) ConflictSize() int {
	return len(kpc_obj.Conflicts)
}

// you found a new conflict so add the conflict to the list
func (kpc_obj *KPC) AddConflict(con Conflict) {
	kpc_obj.Conflicts = append(kpc_obj.Conflicts, con)
}

// get conflicts
func (kpc_obj *KPC) GetConflicts() []Conflict {
	return kpc_obj.Conflicts
}

// get a specific conflict
func (kpc_obj *KPC) GetConflict(name string) *Conflict {
	for _, val := range kpc_obj.Conflicts {
		if val.Name == name {
			return &val
		}
	}
	return nil
}
