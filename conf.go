package conf

import (
	"log"
	"os"
)

var consoleLog = log.New(os.Stdout, "[conf] ", log.LstdFlags)
var DEBUG bool = false

type Config interface {
	Get(section, name string) string
	GetInt(section, name string) int
	Set(section, name, value string)
	GetSection(section string) []SectionConfig
	Reload() error
}

type SectionConfig interface {
	Get(name string) string
	GetInt(name string) int
	Set(name, value string)
}
