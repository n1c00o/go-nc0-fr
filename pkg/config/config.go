package config

// Config represents the configuration for the generator.
//
// Hostname is the hostname of the URL the site will be deployed to.
//
// Modules is a list of Go modules to generate.
type Config struct {
	Hostname string   `yaml:"hostname"`
	Modules  []Module `yaml:"modules"`
}
