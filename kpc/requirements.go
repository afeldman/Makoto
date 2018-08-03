// Copyright Anton Feldmann
//
// packages can depend on different pages. This structure will list all information
//
package kpc

// packages can contain requirements.
// this methods readout the information of additional KPC packages.
// if the packages are not available, then the program is not buildable
type Requirement struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

// build a Requirement structure
// to speedup the workflow.
// the name of the requirement is mandatory, The default Version is '0.1.0'
//
// @param {string} the name of the requirements
//
// @return {*Requirement} The Structure ony containing the base information
func Requirement_Init(name string) (*Requirement) {
	return &Requirement{
		Name: name,
		Version: "0.1.0"}
}

// set a new version number
// each project starts with version '0.1.0'
//
// this function is pagt of the Requirement structure
//
// @param {string} version. the version can set by each user
func (this *Requirement) SetVersion(version string) {
	this.Version = version
}

// set a new name.
// each project has a name. Maybe if the name was written wrong or
// changed druing time of process. This method change the name
//
// this function is pagt of the Requirement structure
//
// @param {string} name. the new name for the requirement
func (this *Requirement) SetName(name string){
	this.Name = name
}

// get version of requirement
// get the requirement version
//
// this function is pagt of the Requirement structure
//
// @return {*string} the requirement version
func (this *Requirement) GetVersion() (*string) {
	return &(this.Version)
}

// get name of requirement
// get the requirement name
//
// this function is pagt of the Requirement structure
//
// @return {*string} the requirement name
func (this *Requirement) GetName() (*string){
	return &(this.Name)
}
