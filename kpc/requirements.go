// Copyright Anton Feldmann
//
// packages can depend on different pages. This structure will list all information
package kpc

// packages can contain requirements.
// this methods readout the information of additional KPC packages.
// if the packages are not available, then the program is not buildable
type Requirement struct {
	Name    string     `toml:"name"`
	Version string     `toml:"version,omitempty"`
	Source  Repository `toml:"source,omitempty"`
}

// build a Requirement structure
// to speedup the workflow.
// the name of the requirement is mandatory, The default Version is '0.1.0'
//
// @param {string} the name of the requirements
//
// @return {*Requirement} The Structure ony containing the base information
func InitRequirement(name string) *Requirement {
	return &Requirement{
		Name:    name,
		Version: "0.1.0"}
}

// set a new version number
// each project starts with version '0.1.0'
//
// this function is pagt of the Requirement structure
//
// @param {string} version. the version can set by each user
func (req *Requirement) SetVersion(version string) {
	req.Version = version
}

// set a new name.
// each project has a name. Maybe if the name was written wrong or
// changed druing time of process. This method change the name
//
// this function is pagt of the Requirement structure
//
// @param {string} name. the new name for the requirement
func (req *Requirement) SetName(name string) {
	req.Name = name
}

// get version of requirement
// get the requirement version
//
// this function is pagt of the Requirement structure
//
// @return {*string} the requirement version
func (req *Requirement) GetVersion() *string {
	return &(req.Version)
}

// get name of requirement
// get the requirement name
//
// this function is pagt of the Requirement structure
//
// @return {*string} the requirement name
func (req *Requirement) GetName() *string {
	return &(req.Name)
}
