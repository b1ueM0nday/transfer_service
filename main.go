package main

import (
	"log"
	"math/big"
	"test_project/transfer"
)

func main() {
	var err error
	mtCli := transfer.New()
	err = mtCli.Deploy("http://localhost:22000", "43ce9f8b44fd0975882e0edba5062ee63cbd66db47ce5cf329609226ebd3f707")
	if err != nil {
		log.Panic(err)
	}
	err = mtCli.Deposit(big.NewInt(100))
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
	}
}
