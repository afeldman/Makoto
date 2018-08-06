package kpc

func (this *KPC)GetMainSourceFile() (*string){
	return &(this.Main)
}

func (this *KPC)SetMainSourceFile(name string){
	this.Main = name
}

func (this *KPC)AddConst(con string){
	this.Consts.Add(con)
}

func (this *KPC)GetConsts() ([]string){
	var req []string
	 for _, val := range this.Consts.Values() {
		 req = append(req, val.(string))
	 }

	return req
}

func (this *KPC)GetConst(con string) (*string){
	var req string

	for i := range this.Consts.Values() {
		req = (this.Consts.Values()[i]).(string)
		if req == con {
			return &req
		}
	}
	return nil
}

func (this *KPC)ConstsSize() (int){
	return this.Consts.Size()
}

func (this *KPC)RejectConst(name string){
	this.Consts.Remove(name)
}

func (this *KPC)AddInclude(inc string){
	this.Includes.Add(inc)
}

func (this *KPC)GetInclude(inc string) (*string){
	var req string

	for i := range this.Includes.Values() {
		req = (this.Includes.Values()[i]).(string)
		if req == inc {
			return &req
		}
	}
	return nil
}

func (this *KPC)IncludesSize() (int){
	return this.Includes.Size()
}

func (this *KPC)RejectInclude(name string){
	this.Includes.Remove(name)
}

func (this *KPC)GetIncludes() ([]string){
	var req []string
	 for _, val := range this.Includes.Values() {
		 req = append(req, val.(string))
	 }

	return req
}

func (this *KPC)AddTye(type_ string){
	this.Types.Add(type_)
}

func (this *KPC)GetType(type_ string) (*string){
	var req string

	for i := range this.Types.Values() {
		req = (this.Types.Values()[i]).(string)
		if req == type_ {
			return &req
		}
	}
	return nil
}

func (this *KPC)TypesSize() (int){
	return this.Types.Size()
}

func (this *KPC)RejectTypes(name string){
	this.Types.Remove(name)
}

func (this *KPC)GetTypes() ([]string){
	var req []string
	 for _, val := range this.Types.Values() {
		 req = append(req, val.(string))
	 }

	return req
}

func (this *KPC)AddForm(form string){
	this.Forms.Add(form)
}

func (this *KPC)GetForm(form string) (*string){
	var req string

	for i := range this.Forms.Values() {
		req = (this.Forms.Values()[i]).(string)
		if req == form {
			return &req
		}
	}
	return nil
}

func (this *KPC)FormsSize() (int){
	return this.Forms.Size()
}

func (this *KPC)RejectForm(name string){
	this.Forms.Remove(name)
}

func (this *KPC)GetForms() ([]string){
	var req []string
	 for _, val := range this.Forms.Values() {
		 req = append(req, val.(string))
	 }

	return req
}

func (this *KPC)GetDicts() ([]string){
	var req []string
	 for _, val := range this.Dicts.Values() {
		 req = append(req, val.(string))
	 }

	return req
}

func (this *KPC)AddDict(dict string){
	this.Dicts.Add(dict)
}

func (this *KPC)GetDict(dict string) (*string){
	var req string

	for i := range this.Dicts.Values() {
		req = (this.Dicts.Values()[i]).(string)
		if req == dict {
			return &req
		}
	}
	return nil
}

func (this *KPC)DictsSize() (int){
	return this.Dicts.Size()
}

func (this *KPC)RejectDict(name string){
	this.Dicts.Remove(name)
}
