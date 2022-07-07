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
	Grpc     gg.Config        `yaml:"grpc"`
	Contract contracts.Config `yaml:"node"`
}

func main() {
	cfg := loadConfig()
	ctx := context.Background()
	cWorker := contracts.NewContract(ctx)

	err := cWorker.Prepare(&cfg.Contract)
	if err != nil {
		log.Fatal(err)
	}

	gs := gg.New(ctx, cWorker)
	err = gs.Connect(cfg.Grpc.Port)
	if err != nil {
		log.Fatal(err)
	}
	gs.Run()
	<-ctx.Done()
}

func loadConfig() *config {

	var cfg config
	b, err := ioutil.ReadFile(configPath)
	if err != nil || b == nil {
		log.Println("run using default config")
		cfg = *defaultConfig()
	} else {
		err = yaml.Unmarshal(b, &cfg)
		if err != nil {
			log.Println("run using default config")
			cfg = *defaultConfig()
		}
	}
	return &cfg
}

func defaultConfig() *config {
	return &config{
		Grpc:     gg.DefaultConfig,
		Contract: contracts.DefaultConfig,
	}
}
