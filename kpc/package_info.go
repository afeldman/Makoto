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

func (this *KPC)AddRequirement(req Requirement){
	this.Requirements = append(this.Requirements,req)
}

func (this *KPC)GetRequirement(name string) (*Requirement){
	for _, val := range this.Requirements {
		if val.Name == name {
			return &val
		}
	}
	return nil
}

func (this *KPC)RequirementSize() (int){
	return len(this.Requirements)
}

func (this *KPC)RejectRequirement(name string){
	var tmp []Requirement
	for _, v := range this.Requirements {
		if v.Name == name {
			continue
		} else {
			tmp = append(tmp, v)
		}
	}
	this.Requirements = tmp
}

func (this *KPC)AddConflict(con Conflict){
	this.Conflicts = append(this.Conflicts, con)
}

func (this *KPC)GetConflicts() ([]Conflict){
	return this.Conflicts
}

func (this *KPC)GetConflict(name string) (*Conflict){
	for _, val := range this.Conflicts {
		if val.Name == name {
			return &val
		}
	}
	return nil
}

func (this *KPC)ConflictSize() (int){
	return len(this.Conflicts)
}

func (this *KPC)RejectConflict(name string){
	var tmp []Conflict
	for _, v := range this.Conflicts {
		if v.Name == name {
			continue
		} else {
			tmp = append(tmp, v)
		}
	}
	this.Conflicts = tmp
}

func (this *KPC)AddAuthor(aut Author){
	this.Authors = append(this.Authors, aut)
}

func (this *KPC)GetAuthor(name string) (*Author){
	for _, val := range this.Authors {
		if val.Name == name {
			return &val
		}
	}
	return nil
}

func (this *KPC)AuthorsSize() (int){
	return len(this.Authors)
}

func (this *KPC)RejectAuthor(name string){
	var tmp []Author
	for _, v := range this.Authors {
		if v.Name == name {
			continue
		} else {
			tmp = append(tmp, v)
		}
	}
	this.Authors = tmp
}

func (this *KPC)AddRepo(repo Repo) {
	this.Repository = repo
}

func (this *KPC)GetRepo() (*Repo){
	return &(this.Repository)
}

func (this *KPC)GetIssue() (*string){
	return &(this.Issue)
}

func (this *KPC)SetIssue(issue string) {
	this.Issue = issue
}

func (this *KPC)GetKeywords() []string{
	return this.Keywords
}

func (this *KPC)ContainKeyword(key string) (bool,string){
	for _, val := range this.Keywords{
		if val == key {
			return true, val
		}
	}
	return false, ""
}

func (this *KPC)SetKeyword(key string) {
	this.Keywords = append(this.Keywords, key)
}
