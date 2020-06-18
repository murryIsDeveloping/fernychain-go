package main

import (
	"fmt"

	"github.com/murryIsDeveloping/fernychain-go/modules/blockchain"
)

func main() {
	fmt.Print("...Starting")
	bc := blockchain.Genisis()
	bc.MineBlock("Hello")
	bc.MineBlock("There")
	bc.MineBlock("You")
	bc.MineBlock("lovely")
	bc.MineBlock("thing")
}
