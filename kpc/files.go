// Copyright Anton Feldmann
//
// KPC is a Karel package managemant file.
//
//This module will handle the file.
//
package kpc

//karel has  amain program file. this method will
//have return the handle
//
//@return {string} name of the file with
func (kpc_obj *KPC) GetMainSourceFile() *string {
	return &(kpc_obj.Main)
}

//set the name of the main program file
//
//@param {strig} file name
func (kpc_obj *KPC) SetMainSourceFile(name string) {
	kpc_obj.Main = name
}

//add a new file with constant parameters
//
//@param {string} path to constant parameter file
func (kpc_obj *KPC) AddConst(con string) {
	kpc_obj.Consts = append(kpc_obj.Consts, con)
}

//get the list of constant parameter files
//
//@return {string array} list of strings
func (kpc_obj *KPC) GetConsts() []string {
	return kpc_obj.Consts
}

//get on constant file path utilizing specific name
//
//@param {string} constant file name
//@return {string*} path to file
func (kpc_obj *KPC) GetConst(con string) *string {
	for _, val := range kpc_obj.Consts {
		if val == con {
			return &val
		}
	}
	return nil
}

func (kpc_obj *KPC) ConstsSize() int {
	return len(kpc_obj.Consts)
}

func (kpc_obj *KPC) RejectConst(name string) {
	var tmp []string
	for _, v := range kpc_obj.Consts {
		if v == name {
			continue
		} else {
			tmp = append(tmp, v)
		}
	}
	kpc_obj.Consts = tmp
}

func (kpc_obj *KPC) AddInclude(inc string) {
	kpc_obj.Includes = append(kpc_obj.Includes, inc)
}

func (kpc_obj *KPC) GetInclude(inc string) *string {
	for _, val := range kpc_obj.Includes {
		if val == inc {
			return &val
		}
	}
	return nil
}

func (kpc_obj *KPC) IncludesSize() int {
	return len(kpc_obj.Includes)
}

func (kpc_obj *KPC) RejectInclude(name string) {
	var tmp []string
	for _, v := range kpc_obj.Includes {
		if v == name {
			continue
		} else {
			tmp = append(tmp, v)
		}
	}
	kpc_obj.Includes = tmp
}

func (kpc_obj *KPC) GetIncludes() []string {
	return kpc_obj.Includes
}

func (kpc_obj *KPC) AddType(type_ string) {
	kpc_obj.Types = append(kpc_obj.Types, type_)
}

func (kpc_obj *KPC) GetType(type_ string) *string {
	for _, val := range kpc_obj.Types {
		if val == type_ {
			return &val
		}
	}
	return nil
}

func (kpc_obj *KPC) TypesSize() int {
	return len(kpc_obj.Types)
}

func (kpc_obj *KPC) RejectType(name string) {
	var tmp []string
	for _, v := range kpc_obj.Types {
		if v == name {
			continue
		} else {
			tmp = append(tmp, v)
		}
	}
	kpc_obj.Types = tmp
}

func (kpc_obj *KPC) GetTypes() []string {
	return kpc_obj.Types
}

func (kpc_obj *KPC) AddForm(form string) {
	kpc_obj.Forms = append(kpc_obj.Forms, form)
}

func (kpc_obj *KPC) GetForm(form string) *string {
	for _, val := range kpc_obj.Forms {
		if val == form {
			return &val
		}
	}
	return nil
}

func (kpc_obj *KPC) FormsSize() int {
	return len(kpc_obj.Forms)
}

func (kpc_obj *KPC) RejectForm(name string) {
	var tmp []string
	for _, v := range kpc_obj.Forms {
		if v == name {
			continue
		} else {
			tmp = append(tmp, v)
		}
	}
	kpc_obj.Forms = tmp

}

func (kpc_obj *KPC) GetForms() []string {
	return kpc_obj.Forms
}

func (kpc_obj *KPC) GetDicts() []string {
	return kpc_obj.Dicts
}

func (kpc_obj *KPC) AddDict(dict string) {
	kpc_obj.Dicts = append(kpc_obj.Dicts, dict)
}

func (kpc_obj *KPC) GetDict(dict string) *string {
	for _, val := range kpc_obj.Dicts {
		if val == dict {
			return &val
		}
	}
	return nil
}

func (kpc_obj *KPC) DictsSize() int {
	return len(kpc_obj.Dicts)
}

func (kpc_obj *KPC) RejectDict(name string) {
	var tmp []string
	for _, v := range kpc_obj.Dicts {
		if v == name {
			continue
		} else {
			tmp = append(tmp, v)
		}
	}
	kpc_obj.Dicts = tmp
}
