package kpc

func (kpc_obj *KPC) SetPrefix(path string) {
	kpc_obj.Prefix = path
}

func (kpc_obj *KPC) GetPrefix() *string {
	return &(kpc_obj.Prefix)
}

func (kpc_obj *KPC) SetSourceDir(path string) {
	kpc_obj.SrcDir = path
}

func (kpc_obj *KPC) GetSourceDir() *string {
	return &(kpc_obj.SrcDir)
}

func (kpc_obj *KPC) SetTypeDir(path string) {
	kpc_obj.TypeDir = path
}

func (kpc_obj *KPC) GetTypeDir() *string {
	return &(kpc_obj.TypeDir)
}

func (kpc_obj *KPC) SetIncludeDir(path string) {
	kpc_obj.IncludeDir = path
}

func (kpc_obj *KPC) GetIncludeDir() *string {
	return &(kpc_obj.IncludeDir)
}

func (kpc_obj *KPC) SetConstDir(path string) {
	kpc_obj.ConstDir = path
}

func (kpc_obj *KPC) GetConstDir() *string {
	return &(kpc_obj.ConstDir)
}

func (kpc_obj *KPC) SetFormDir(path string) {
	kpc_obj.FormDir = path
}

func (kpc_obj *KPC) GetFormDir() *string {
	return &(kpc_obj.FormDir)
}

func (kpc_obj *KPC) SetDictDir(path string) {
	kpc_obj.DictDir = path
}

func (kpc_obj *KPC) GetDictDir() *string {
	return &(kpc_obj.DictDir)
}
