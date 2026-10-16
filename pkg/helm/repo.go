package helm

import (
	"fmt"

	helmpush "github.com/chartmuseum/helm-push/pkg/helm"
	v3cli "helm.sh/helm/v3/pkg/cli"
	v3repo "helm.sh/helm/v3/pkg/repo"
	v4cli "helm.sh/helm/v4/pkg/cli"
	v4repo "helm.sh/helm/v4/pkg/repo/v1"
)

// RepoEntry holds the config fields for a named helm repo.
type RepoEntry struct {
	URL      string
	Username string
	Password string
	CAFile   string
	CertFile string
	KeyFile  string
}

// GetRepoEntryByName looks up a named repo in the local helm repo config.
func GetRepoEntryByName(name string) (*RepoEntry, error) {
	if helmpush.HelmMajorVersionCurrent() == helmpush.HelmMajorVersion3 {
		f, err := v3repo.LoadFile(v3cli.New().RepositoryConfig)
		if err != nil {
			return nil, err
		}
		for _, e := range f.Repositories {
			if e.Name == name {
				return &RepoEntry{e.URL, e.Username, e.Password, e.CAFile, e.CertFile, e.KeyFile}, nil
			}
		}
	} else {
		f, err := v4repo.LoadFile(v4cli.New().RepositoryConfig)
		if err != nil {
			return nil, err
		}
		for _, e := range f.Repositories {
			if e.Name == name {
				return &RepoEntry{e.URL, e.Username, e.Password, e.CAFile, e.CertFile, e.KeyFile}, nil
			}
		}
	}
	return nil, fmt.Errorf("no repo named %q found", name)
}

// GetRepoEntryByURL finds the first repo entry whose URL matches the given URL.
// Returns nil (no error) if no matching entry exists.
func GetRepoEntryByURL(repoURL string) (*RepoEntry, error) {
	if helmpush.HelmMajorVersionCurrent() == helmpush.HelmMajorVersion3 {
		f, err := v3repo.LoadFile(v3cli.New().RepositoryConfig)
		if err != nil {
			return nil, err
		}
		for _, e := range f.Repositories {
			if e.URL == repoURL {
				return &RepoEntry{e.URL, e.Username, e.Password, e.CAFile, e.CertFile, e.KeyFile}, nil
			}
		}
	} else {
		f, err := v4repo.LoadFile(v4cli.New().RepositoryConfig)
		if err != nil {
			return nil, err
		}
		for _, e := range f.Repositories {
			if e.URL == repoURL {
				return &RepoEntry{e.URL, e.Username, e.Password, e.CAFile, e.CertFile, e.KeyFile}, nil
			}
		}
	}
	return nil, nil
}
