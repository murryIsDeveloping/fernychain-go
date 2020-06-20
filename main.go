package main

import (
	"github.com/murryIsDeveloping/fernychain-go/modules/wallet"
)

func main() {
	// fmt.Print("...Starting")
	// bc := blockchain.Genisis()
	// bc.MineBlock("Hello")
	// bc.MineBlock("There")
	// bc.MineBlock("You")
	// bc.MineBlock("lovely")
	// bc.MineBlock("thing")

	// nc := blockchain.Genisis()
	// nc.MineBlock("Hello")
	// nc.MineBlock("There")

	// nc.ReplaceChain(bc)

	// fmt.Print("PRINTING NEW BLOCK !!!!! \n\n")
	// nc.PrintBlocks()
	wallet.GenerateWallet()

}
