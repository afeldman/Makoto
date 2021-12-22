// Copyright Anton Feldmann
//
// KPC is a Karel package management structure.
// This structure is used to inform the tools about the structure of the project
package kpc

// get the referenced homepage
func (kpc_obj *KPC) GetHomepage() *string {
	return &(kpc_obj.Homepage)
}

// set the homepage
func (kpc_obj *KPC) SetHomepage(url string) {
	kpc_obj.Homepage = url
}
