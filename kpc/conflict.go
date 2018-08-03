package kpc

type Conflict struct {
	Name     string   `yaml:"name"`
	Versions []string `yaml:"version"`
}

func Conflict_Init(name, version string) (*Conflict){
	tmp := []string{version}

	return &Conflict{
		Name: name,
		Versions: tmp,
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
}
