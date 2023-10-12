package kpc

// Copyright Anton Feldmann
//
// KPC is a Karel package management structure.
// This structure is used to inform the tools about the structure of the project

// get the desctiption
func (kpc_obj *KPC) GetDescription() *string {
	return &(kpc_obj.Description)
}

// set description
func (kpc_obj *KPC) SetDescription(desc string) {
	kpc_obj.Description = desc
}
