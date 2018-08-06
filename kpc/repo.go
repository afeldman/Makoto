package kpc

type Repo struct {
	Type string `json:"repo"`
	URL  string `json:"repo"`
}

func Repo_Init() (*Repo){
	return &Repo{Type: "", URL: ""}
}

func (this *Repo)SetType(repo_type string) {
	this.Type = repo_type
}

func (this *Repo)GetType() (*string) {
	return &(this.Type)
}

func (this *Repo)SetURL(address string) {
	this.URL = address
}

func (this *Repo)GetURL() (*string) {
	return &(this.URL)
}
