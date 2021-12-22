// Copyright Anton Feldmann
//
// KPC is a Karel package management structure.
// This structure is used to inform the tools about the structure of the project
package kpc

// get project name
func (kpc_obj *KPC) GetName() *string {
	return &(kpc_obj.Name)
}

// get the project name
func (kpc_obj *KPC) SetName(name string) {
	kpc_obj.Name = name
}
