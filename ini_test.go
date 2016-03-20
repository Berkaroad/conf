package conf

import (
	"testing"
)

var config Config = LoadIniConfig("testconfig.ini")

func Test_Get(t *testing.T) {
	config.Get("command", "concurrent_num")
	t.Error(config.GetInt("command", "concurrent_num"))
}

func Test_Set(t *testing.T) {
	// config.Set("command", "concurrent_num", "2")
}

func Test_GetSection(t *testing.T) {
	sections := config.GetSection("command")
	if sections != nil {
		for _, sectionConfig := range sections {
			t.Error(sectionConfig.GetInt("concurrent_num"))
		}
	}
}

func Test_Reload(t *testing.T) {
	config.Reload()
}
