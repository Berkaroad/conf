/*
Package conf is to provide a config, such as ini, json, xml etc.
*/
package conf

import (
	"log"
	"os"
)

var consoleLog = log.New(os.Stdout, "[conf] ", log.LstdFlags)

// DEBUG is switcher for debug
var DEBUG = false

// Config is an interface for config
type Config interface {
	// Get value
	Get(section, name string) string
	// GetInt is to get interger value
	GetInt(section, name string) int
	// Set value
	Set(section, name, value string)
	// GetSection is to get section
	GetSection(section string) []SectionConfig
	// Reload config
	Reload() error
}

// SectionConfig is section part for config
type SectionConfig interface {
	// Get value
	Get(name string) string
	//GetInt is to get interger value
	GetInt(name string) int
	// Set value
	Set(name, value string)
}
