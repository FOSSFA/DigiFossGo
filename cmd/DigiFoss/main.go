package main

import (
	"flag"
	"github.com/amirhossein-ka/DigiFossGo/internal/application"
	"github.com/amirhossein-ka/DigiFossGo/internal/config"
	"log"
)

var cfgPath string
var cfg = &config.Config{}

const defaultConfig = "build/config/config.yaml"

func init() {
	parseCLI()
	if err := config.Parse(cfgPath, cfg); err != nil {
		log.Fatalln(err)
	} 
  config.ParseEnv(cfg)
	config.SetConfig(cfg)
}

func main() {
	if err := application.Run(cfg); err != nil {
		log.Fatalln(err)
	}

}

func parseCLI() {
	flag.StringVar(&cfgPath, "config", defaultConfig, "-config /path/to/cfg")
	flag.StringVar(&cfgPath, "c", defaultConfig, "-c /path/to/cfg (short hand)")
	flag.Parse()
}
