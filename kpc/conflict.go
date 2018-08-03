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

func (this *Conflict)AddElement(version string){

	if !this.ContainsVersion(version) {
		this.Versions = append(this.Versions, version)
	}

}

func (this *Conflict)GetLength() (int) {
	return len(this.Versions)
}

func (this *Conflict)find_elem(version string) (bool, int){
	for i := range this.Versions {
		if this.Versions[i] == version {
			return true, i
		}
	}

	return false, -1
}

func (this *Conflict)ChangeVersion(old, new string) (bool){
	has_elem, idx := this.find_elem(old)
	if  has_elem{
		this.Versions[idx] = new
		return true;
	}
	return false
}


func (this *Conflict)ContainsVersion(version string) (bool){
	has_elem, _ := this.find_elem(version)
	return has_elem;
}

func (this *Conflict)GetName()(*string){
	return &(this.Name)
}

func (this *Conflict)SetName(name string){
	this.Name = name
}

func (this *Conflict)RejectVersion(version string){
	has_elem, idx := this.find_elem(version)
	if has_elem && this.GetLength() > 1{
		this.Versions[idx] = ""
	}

	var r []string

	for _, str := range this.Versions{
		if str != ""{
			r = append(r,str)
		}
	}

	this.Versions = r
}
