package kpc

func (this *KPC)GetMainSourceFile() (*string){
	return &(this.Main)
}

func (this *KPC)SetMainSourceFile(name string){
	this.Main = name
}

func (this *KPC)AddConst(con string){
	this.Consts = append(this.Consts,con)
}

func (this *KPC)GetConsts() ([]string){
	return this.Consts
}

func (this *KPC)GetConst(con string) (*string){
	for _, val := range this.Consts {
		if val == con {
			return &val
		}
	}
	return nil
}

func (this *KPC)ConstsSize() (int){
	return len(this.Consts)
}

func (this *KPC)RejectConst(name string){
	var tmp []string
	for _, v := range this.Consts {
		if v == name {
			continue
		} else {
			tmp = append(tmp, v)
		}
	}
	this.Consts = tmp
}

func (this *KPC)AddInclude(inc string){
	this.Includes = append(this.Includes, inc)
}

func (this *KPC)GetInclude(inc string) (*string){
	for _, val := range this.Includes {
		if val == inc {
			return &val
		}
	}
	return nil
}

func (this *KPC)IncludesSize() (int){
	return len(this.Includes)
}

func (this *KPC)RejectInclude(name string){
	var tmp []string
	for _, v := range this.Includes {
		if v == name {
			continue
		} else {
			tmp = append(tmp, v)
		}
	}
	this.Includes = tmp
}

func (this *KPC)GetIncludes() ([]string){
	return this.Includes
}

func (this *KPC)AddType(type_ string){
	this.Types = append(this.Types,type_)
}

func (this *KPC)GetType(type_ string) (*string){
	for _, val := range this.Types {
		if val == type_ {
			return &val
		}
	}
	return nil
}

func (this *KPC)TypesSize() (int){
	return len(this.Types)
}

func (this *KPC)RejectType(name string){
	var tmp []string
	for _, v := range this.Types {
		if v == name {
			continue
		} else {
			tmp = append(tmp, v)
		}
	}
	this.Types = tmp
}

func (this *KPC)GetTypes() ([]string){
	return this.Types
}

func (this *KPC)AddForm(form string){
	this.Forms = append(this.Forms, form)
}

func (this *KPC)GetForm(form string) (*string){
	for _, val := range this.Forms {
		if val == form {
			return &val
		}
	}
	return nil
}

func (this *KPC)FormsSize() (int){
	return len(this.Forms)
}

func (this *KPC)RejectForm(name string){
	var tmp []string
	for _, v := range this.Forms {
		if v == name {
			continue
		} else {
			tmp = append(tmp, v)
		}
	}
	this.Forms = tmp

}

func (this *KPC)GetForms() ([]string){
	return this.Forms
}

func (this *KPC)GetDicts() ([]string){
	return this.Dicts
}

func (this *KPC)AddDict(dict string){
	this.Dicts = append(this.Dicts,dict)
}

func (this *KPC)GetDict(dict string) (*string){
	for _, val := range this.Dicts {
		if val == dict {
			return &val
		}
	}
	return nil
}

func (this *KPC)DictsSize() (int){
	return len(this.Dicts)
}

func (this *KPC)RejectDict(name string){
	var tmp []string
	for _, v := range this.Dicts {
		if v == name {
			continue
		} else {
			tmp = append(tmp, v)
		}
	}
	this.Dicts = tmp
}
