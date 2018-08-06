package kpc

func (this *KPC) SetPrefix(path string){
	this.Prefix = path
}

func (this *KPC) GetPrefix() (*string){
	return &(this.Prefix)
}

func (this *KPC) SetSourceDir(path string) {
	this.SrcDir = path
}

func (this *KPC) GetSourceDir() (*string){
	return &(this.SrcDir)
}

func (this *KPC) SetTypeDir(path string){
	this.TypeDir = path
}

func (this *KPC) GetTypeDir() (*string){
	return &(this.TypeDir)
}

func (this *KPC) SetIncludeDir(path string){
	this.IncludeDir = path
}

func (this *KPC) GetIncludeDir() (*string){
	return &(this.IncludeDir)
}

func (this *KPC) SetConstDir(path string) {
	this.ConstDir = path
}

func (this *KPC) GetConstDir() (*string) {
	return &(this.ConstDir)
}

func (this *KPC) SetFormDir(path string) {
	this.FormDir = path
}

func (this *KPC) GetFormDir() (*string) {
	return &(this.FormDir)
}

func (this *KPC) SetDictDir(path string) {
	this.DictDir = path
}

func (this *KPC) GetDictDir() (*string) {
	return &(this.DictDir)
}
