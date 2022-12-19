package config

// Module represents a Go module in the configuration.
//
// Prefix represents the module's import prefix, without the hostname
// and the trailing slash.
// For example, the module `go.example.com/foo` would have the prefix `foo`.
//
// Dir is a URI to a page that lists the files inside a directory of a module.
// For a GitHub repository it would be https://github.com/USER/REPOSITORY/tree/MASTER_BRANCH{/dir}.
//
// More information can be found here: https://github.com/golang/gddo/wiki/Source-Code-Links
//
// File is a URI to a page containing the content of a file, with support for a specific line.
// For a GitHub repository it would be https://github.com/USER/REPOSITORY/blob/MASTER_BRANCH{/dir}/{file}#L{line}.
//
// More information can be found here: https://github.com/golang/gddo/wiki/Source-Code-Links
//
// Repo is the URI (preferably HTTPS) to the repository of the module.
// For a GitHub repository it would be https://github.com/USER/REPOSITORY.git
//
// VCS is the source control system used for the repository.
// The Go toolchain supports "git" (Git), "hg" (Mercurial),
// "svn" (Subversion) and "bzr" (Bazaar).
type Module struct {
	Prefix string `yaml:"prefix"`
	Dir    string `yaml:"dir"`
	File   string `yaml:"file"`
	Repo   string `yaml:"repo"`
	VCS    string `yaml:"vcs"`
}
