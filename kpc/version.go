// Copyright Anton Feldmann
//
// KPC is a Karel package management structure.
// This structure is used to inform the tools about the structure of the project
package kpc

// get the package version
func (kpc_obj *KPC) GetVersion() *string {
	return &(kpc_obj.Version)
}

// set the current package version
func (kpc_obj *KPC) SetVersion(version string) {
	kpc_obj.Version = version
}
