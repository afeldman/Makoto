package kpc

func (this *KPC)GetName() (*string){
	return &(this.Name)
}

func (this *KPC)SetName(name string){
	this.Name = name
}

func (this *KPC)SetDescription(desc string){
	this.Description = desc
}

func (this *KPC)GetDescription() (*string){
	return &(this.Description)
}

func (this *KPC)SetVersion(version string){
	this.Version = version
}

func (this *KPC)GetVersion() (*string){
	return &(this.Version)
}

func (this *KPC)GetHomepage() (*string){
	return &(this.Homepage)
}

func (this *KPC)SetHomepage(url string){
	this.Homepage = url
}

func (this *KPC)SetRequirement(req Requirement){
	this.Requirements.Add(req)
}

func (this *KPC)GetRequirement(name string) (*Requirement){
	var req Requirement

	for i := range this.Requirements.Values() {
		req = (this.Requirements.Values()[i]).(Requirement)
		if req.Name == name {
			return &(req)
		}
	}
	return nil
}
