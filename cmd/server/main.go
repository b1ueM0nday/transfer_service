package main

import (
	"context"
	"github.com/b1uem0nday/transfer_service/internal/contracts"
	"github.com/b1uem0nday/transfer_service/internal/gg"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const configPath = "./config/config.yaml"

type config struct {
	Port           string `yaml:"port"`
	AddrPath       string `yaml:"address_path"`
	DeployAddress  string `yaml:"deploy_address"`
	PrivateKeyPath string `yaml:"pk_path"`
}

var defConfig = config{
	Port:           "3000",
	AddrPath:       "./config",
	PrivateKeyPath: "./config",
	DeployAddress:  "http://localhost:22000",
}

func main() {
	cfg := loadConfig()
	ctx := context.Background()
	cWorker := new(contracts.Worker)

	err := cWorker.Prepare(cfg.AddrPath, cfg.DeployAddress, cfg.PrivateKeyPath)
	if err != nil {
		log.Fatal(err)
	}

	gs := gg.New(ctx, cWorker)
	err = gs.Connect(cfg.Port)
	if err != nil {
		log.Fatal(err)
	}
	gs.Run()
}

func loadConfig() *config {
	var cfg config
	b, err := ioutil.ReadFile(configPath)
	if err != nil || b == nil {
		log.Println("run using default config")
		cfg = defConfig
	} else {
		err = yaml.Unmarshal(b, &cfg)
		if err != nil {
			log.Println("run using default config")
			cfg = defConfig
		}
	}
	return &cfg
}
