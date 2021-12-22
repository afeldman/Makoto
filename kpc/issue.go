// Copyright Anton Feldmann
//
// KPC is a Karel package management structure.
// This structure is used to inform the tools about the structure of the project
package kpc

// the the issue reference page
func (kpc_obj *KPC) GetIssue() *string {
	return &(kpc_obj.Issue)
}

// set the issue page
func (kpc_obj *KPC) SetIssue(issue string) {
	kpc_obj.Issue = issue
}
