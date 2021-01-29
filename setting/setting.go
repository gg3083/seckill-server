package setting

import (
	"github.com/go-ini/ini"
	"log"
)

type Sql struct {
	Goods string
	Userinfo string
}

var SqlScript = &Sql{}

var cfg *ini.File

// InitSetting initialize the configuration instance
func InitSetting() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.InitSetting, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("sql", SqlScript)
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
