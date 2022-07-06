package main

import (
	"context"
	"github.com/b1uem0nday/transfer_service/internal/gg"
	"github.com/b1uem0nday/transfer_service/internal/transfer"
	"log"
)

func main() {
	var err error
	ctx := context.Background()
	mtCli := new(transfer.Client)
	err = mtCli.Deploy("http://localhost:22000", "43ce9f8b44fd0975882e0edba5062ee63cbd66db47ce5cf329609226ebd3f707")
	if err != nil {
		log.Panic(err)
	}

	gs := gg.New(ctx, mtCli)
	err = gs.Connect("3000")
	if err != nil {
		log.Panic(err)
	}
	select {
	case <-ctx.Done():
		return
	}
}
