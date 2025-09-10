package kpc

// Repository beschreibt die Quellcode-Herkunft eines Pakets.
// Standard ist ein Git-Repository, zusätzlich mit optionalem Tag, Commit (rev) oder Branch.
type Repository struct {
	URL    string `toml:"url"`              // Git URL
	Tag    string `toml:"tag,omitempty"`    // Tag Name
	Rev    string `toml:"rev,omitempty"`    // Commit Hash
	Branch string `toml:"branch,omitempty"` // Optional Branch
}

// InitRepository liefert ein Repository mit Defaultwerten zurück.
func InitRepository() *Repository {
	return &Repository{
		URL:    "",
		Tag:    "",
		Rev:    "",
		Branch: "",
	}
}

/*********************** Getter & Setter ***********************/

// URL
func (repo *Repository) SetURL(address string) { repo.URL = address }
func (repo *Repository) GetURL() *string       { return &repo.URL }

// Tag
func (repo *Repository) SetTag(tag string) { repo.Tag = tag }
func (repo *Repository) GetTag() *string   { return &repo.Tag }

// Rev
func (repo *Repository) SetRev(rev string) { repo.Rev = rev }
func (repo *Repository) GetRev() *string   { return &repo.Rev }

// Branch
func (repo *Repository) SetBranch(branch string) { repo.Branch = branch }
func (repo *Repository) GetBranch() *string      { return &repo.Branch }

// Zugriff über KPC
func (kpc_obj *KPC) GetRepo() *Repository    { return &kpc_obj.Source }
func (kpc_obj *KPC) AddRepo(repo Repository) { kpc_obj.Source = repo }
