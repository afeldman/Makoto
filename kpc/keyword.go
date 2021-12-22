// Copyright Anton Feldmann
//
// KPC is a Karel package management structure.
// This structure is used to inform the tools about the structure of the project
package kpc

// get all keqwords
func (kpc_obj *KPC) GetKeywords() []string {
	return kpc_obj.Keywords
}

// does the kexword exist
func (kpc_obj *KPC) ContainKeyword(key string) (bool, string) {
	for _, val := range kpc_obj.Keywords {
		if val == key {
			return true, val
		}
	}
	return false, ""
}

// set a new keyword
func (kpc_obj *KPC) SetKeyword(key string) {
	kpc_obj.Keywords = append(kpc_obj.Keywords, key)
}
