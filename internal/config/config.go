package config

import (
	"flag"
	"log"

	"github.com/notnull-co/cfg"
)

var (
	Cfg config
)

type config struct {
	DB struct {
		ConnectionString string `cfg:"connection_string"`
	} `cfg:"db"`
}

func ParseFromFlags() {
	var configDir string

	flag.StringVar(&configDir, "config-dir", "../internal/config/", "Configuration file directory")
	flag.Parse()

	parse(configDir)
}

func parse(dirs ...string) {
	if err := cfg.Load(&Cfg,
		cfg.Dirs(dirs...),
		cfg.UseEnv("app"),
	); err != nil {
		log.Panic(err)
	}
}

func Get() config {
	return Cfg
}
