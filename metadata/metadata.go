package metadata

import "fmt"

type Metadata struct {
	ProductFiles []ProductFile `yaml:"product_files"`
	Release      Release       `yaml:"release"`
}

type Release struct {
	ReleaseType     string `yaml:"release_type"`
	EulaSlug        string `yaml:"eula_slug"`
	ProductVersion  string `yaml:"product_version"`
	ReleaseDate     string `yaml:"release_date"`
	Description     string `yaml:"description"`
	ReleaseNotesURL string `yaml:"release_notes_url"`
	Availability    string `yaml:"availability"`
	UserGroupIDs    []int  `yaml:"user_group_ids"`
}

type ProductFile struct {
	File        string `yaml:"file"`
	Description string `yaml:"description"`
	UploadAs    string `yaml:"upload_as"`
}

func (m Metadata) Validate() error {
	for _, productFile := range m.ProductFiles {
		if productFile.File == "" {
			return fmt.Errorf("empty value for file")
		}
	}
	return nil
}
