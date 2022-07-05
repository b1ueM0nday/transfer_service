package main

import (
	"context"
	"github.com/b1uem0nday/transfer_service/gg"
	"github.com/b1uem0nday/transfer_service/transfer"
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
	/*err = mtCli.Deposit(big.NewInt(100))
	if err != nil {
		log.Panic(err)
	}
	err = mtCli.Withdraw(big.NewInt(10))
	if err != nil {
		log.Panic(err)
	}
	err = mtCli.Transfer("0x8107cf2ca713cfde53e9ab5404bd79f429b5d176", big.NewInt(90))
	if err != nil {
		log.Panic(err)
	}*/

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
