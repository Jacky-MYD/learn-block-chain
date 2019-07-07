package main

import (
	"fmt"
	"publicChain/part3-persistence-and-cli/BLC"
	"time"
)

func main()  {
	//block := BLC.NewGenesisBlock()
	//fmt.Println(block.Timestamp)
	//fmt.Printf("%x\n", block.PrevBlockHash)
	//fmt.Println(block.Data)
	//fmt.Printf("%x", block.Hash)


	blockchain := BLC.NewBlockChain()

	blockchain.AddBlock("Send 20 BTC to Jacky From aa")

	blockchain.AddBlock("Send 40 BTC to Jacky From bb")

	blockchain.AddBlock("Send 60 BTC to Jacky From cc")

	for _, blcok := range blockchain.Blocks {
		fmt.Printf("Data: %s \n", string(blcok.Data))
		fmt.Printf("PrevBlockHash: %x \n", string(blcok.PrevBlockHash))
		fmt.Printf("Timestamp: %s \n", time.Unix(blcok.Timestamp, 0).Format("2006-01-02 03:04:05 PM")) // 格式化时间是固定2006年的
		fmt.Printf("Hash: %x \n", blcok.Hash)
		fmt.Printf("Nonce: %d \n", blcok.Nonce)
		fmt.Println()
	}

}