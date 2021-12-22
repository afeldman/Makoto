// Copyright Anton Feldmann
//
// KPC is a Karel package management structure.
// This structure is used to inform the tools about the structure of the project
package kpc

// get all rewuirements
func (kpc_obj *KPC) GetRequirement(name string) *Requirement {
	for _, val := range kpc_obj.Requirements {
		if val.Name == name {
			return &val
		}
	}
	return nil
}

// add a requirement
func (kpc_obj *KPC) AddRequirement(req Requirement) {
	kpc_obj.Requirements = append(kpc_obj.Requirements, req)
}

// how many requirements are stored
func (kpc_obj *KPC) RequirementSize() int {
	return len(kpc_obj.Requirements)
}

// delete a requirement
func (kpc_obj *KPC) RejectRequirement(name string) {
	var tmp []Requirement
	for _, v := range kpc_obj.Requirements {
		if v.Name == name {
			continue
		} else {
			tmp = append(tmp, v)
		}
	}
	kpc_obj.Requirements = tmp
}
