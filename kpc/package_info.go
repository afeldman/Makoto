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

func (this *KPC)AddRequirement(req *Requirement){
	this.Requirements.Add(req)
}

func (this *KPC)GetRequirement(name string) (*Requirement){
	var req *Requirement

	for i := range this.Requirements.Values() {
		req = (this.Requirements.Values()[i]).(*Requirement)
		if req.Name == name {
			return req
		}
	}
	return nil
}

func (this *KPC)RequirementSize() (int){
	return this.Requirements.Size()
}

func (this *KPC)RejectRequirement(name string){
	var req *Requirement

	for i := range this.Requirements.Values() {
		req = (this.Requirements.Values()[i]).(*Requirement)
		if req.Name == name {
			this.Requirements.Remove(req)
		}
	}
}

func (this *KPC)AddConflict(con *Conflict){
	this.Conflicts.Add(con)
}

func (this *KPC)GetConflicts() ([]Conflict){
	var req []Conflict
	 for _, val := range this.Conflicts.Values() {
		 req = append(req, val.(Conflict))
	 }

	return req
}

func (this *KPC)GetConflict(name string) (*Conflict){
	var con *Conflict

	for i := range this.Conflicts.Values() {
		con = (this.Conflicts.Values()[i]).(*Conflict)
		if con.Name == name {
			return con
		}
	}
	return nil
}

func (this *KPC)ConflictSize() (int){
	return this.Conflicts.Size()
}

func (this *KPC)RejectConflict(name string){
	var req *Conflict

	for i := range this.Conflicts.Values() {
		req = (this.Conflicts.Values()[i]).(*Conflict)
		if req.Name == name {
			this.Conflicts.Remove(req)
		}
	}
}

func (this *KPC)AddAuthor(aut *Author){
	this.Authors.Add(aut)
}

func (this *KPC)GetAuthor(name string) (*Author){
	var req *Author

	for i := range this.Authors.Values() {
		req = (this.Authors.Values()[i]).(*Author)
		if req.Name == name {
			return req
		}
	}
	return nil
}

func (this *KPC)AuthorsSize() (int){
	return this.Authors.Size()
}

func (this *KPC)RejectAuthor(name string){
	var req *Author

	for i := range this.Authors.Values() {
		req = (this.Authors.Values()[i]).(*Author)
		if req.Name == name {
			this.Authors.Remove(req)
		}
	}
}

func (this *KPC)AddRepo(repo *Repo) {
	this.Repository = repo
}

func (this *KPC)GetRepo() (*Repo){
	return this.Repository
}

func (this *KPC)GetIssue() (*string){
	return &(this.Issue)
}

func (this *KPC)SetIssue(issue string) {
	this.Issue = issue
}
