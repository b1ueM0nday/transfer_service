package main

import (
	"context"
	"github.com/b1uem0nday/transfer_service/internal/client"
	"github.com/b1uem0nday/transfer_service/internal/gg"
	"github.com/b1uem0nday/transfer_service/internal/repository"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const configPath = "./config/config.yaml"

type config struct {
	Grpc     gg.Config         `yaml:"grpc"`
	Contract client.Config     `yaml:"node"`
	Base     repository.Config `yaml:"base"`
}

func main() {
	cfg := loadConfig()
	ctx := context.Background()
	db := repository.NewRepository(ctx, &cfg.Base)
	err := db.Connect(cfg.Base.Address, cfg.Base.Login, cfg.Base.Password, cfg.Base.Port)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	client := client.NewClient(db, ctx)

	err = client.Prepare(&cfg.Contract)
	if err != nil {
		log.Fatal(err)
	}

	gs := gg.New(ctx, client)
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
		Contract: client.DefaultConfig,
		Base:     repository.DefaultConfig,
	}
}
