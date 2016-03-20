/*
import (
	"github.com/berkaroad/conf"
)

func main(){
	config := conf.LoadInitConfig("~/config1.ini")
	config.Get("command", "concurrent_num")
	config.Set("command", "concurrent_num", "3")
	if err := config.Reload(); err == nil {

	}
}
*/
package conf

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type iniConfig struct {
	filepath         string
	loaded           bool
	sectionConfigDic map[string][]*iniSectionConfig
	conflist         []map[string]map[string]string
}

type iniSectionConfig struct {
	sectionName string
	keyvals     map[string]string
}

func (self *iniSectionConfig) Get(name string) string {
	if DEBUG {
		consoleLog.Printf("[info] Get config:%s.%s.\n", self.sectionName, name)
	}
	return self.keyvals[name]
}

func (self *iniSectionConfig) GetInt(name string) int {
	str := self.Get(name)
	if val, err := strconv.Atoi(str); err == nil {
		return val
	} else {
		return 0
	}
}

func (self *iniSectionConfig) Set(name, value string) {
	self.keyvals[name] = value
	if DEBUG {
		consoleLog.Printf("[info] Set config:%s.%s=%s.\n", self.sectionName, name, value)
	}
}

func LoadIniConfig(filepath string) Config {
	c := new(iniConfig)
	c.sectionConfigDic = make(map[string][]*iniSectionConfig)
	c.filepath = filepath
	if err := c.Reload(); err == nil {
		if DEBUG {
			consoleLog.Printf("[info] Load file \"%s\" success.\n", c.filepath)
		}
	} else {
		consoleLog.Panicf("[error] Load file \"%s\" error:%s!\n", c.filepath, err.Error())
	}
	return c
}

func (self *iniConfig) Get(section, name string) string {
	if self.loaded {
		sectionConfigs := self.GetSection(section)
		if sectionConfigs != nil && len(sectionConfigs) > 0 {
			return sectionConfigs[0].Get(name)
		}
	}
	return ""
}

func (self *iniConfig) GetInt(section, name string) int {
	str := self.Get(section, name)
	if val, err := strconv.Atoi(str); err == nil {
		return val
	} else {
		return 0
	}
}

func (self *iniConfig) Set(section, name, value string) {
	if self.loaded {
		sectionConfigs := self.GetSection(section)
		if sectionConfigs != nil && len(sectionConfigs) > 0 {
			sectionConfigs[0].Set(name, value)
		}
	}
}

func (self *iniConfig) GetSection(section string) []SectionConfig {
	if self.loaded {
		if self.sectionConfigDic[section] != nil && len(self.sectionConfigDic[section]) > 0 {
			sectionConfigs := make([]SectionConfig, len(self.sectionConfigDic[section]))
			for i, sectionConfig := range self.sectionConfigDic[section] {
				sectionConfigs[i] = sectionConfig
			}
			return sectionConfigs
		}
	}
	return nil
}

func (self *iniConfig) Reload() error {
	self.sectionConfigDic = make(map[string][]*iniSectionConfig)
	file, err := os.OpenFile(self.filepath, os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	var section string
	var sectionConfig *iniSectionConfig
	buf := bufio.NewReader(file)
	for {
		l, err := buf.ReadString('\n')
		line := strings.TrimSpace(l)
		if err != nil {
			if err != io.EOF {
				return err
			}
			if len(line) == 0 {
				break
			}
		}
		switch {
		case len(line) == 0 || line[0] == '#':
		case line[0] == '[' && len(line) > 2 && line[len(line)-1] == ']':
			section = strings.TrimSpace(line[1 : len(line)-1])
			if self.sectionConfigDic[section] == nil {
				self.sectionConfigDic[section] = []*iniSectionConfig{new(iniSectionConfig)}
				sectionConfig = self.sectionConfigDic[section][0]
			} else {
				self.sectionConfigDic[section] = append(self.sectionConfigDic[section], new(iniSectionConfig))
				sectionConfig = self.sectionConfigDic[section][len(self.sectionConfigDic[section])-1]
			}
			sectionConfig.sectionName = section
			sectionConfig.keyvals = make(map[string]string)
		default:
			if sectionConfig == nil {
				return errors.New("Cann't find \"[section]\" in single line")
			}
			i := strings.IndexAny(line, "=")
			if i < 0 {
				return errors.New("Cann't find \"=\" in single line")
			} else if i == 0 {
				return errors.New("Cann't find name before \"=\" in single line")
			}
			key := strings.TrimSpace(line[0:i])
			val := strings.TrimSpace(line[i+1 : len(line)])
			sectionConfig.keyvals[key] = val
		}
	}
	self.loaded = true
	return nil
}
