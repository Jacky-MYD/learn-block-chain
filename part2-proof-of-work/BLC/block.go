package BLC

import (
	"fmt"
	"time"
)

type Block struct {

	// 时间戳 创建区块时的时间
	Timestamp int64
	// 上一个区块的哈希 父Hash
	PrevBlockHash []byte
	// data 交易数据
	Data []byte
	// Hash 当前区块的Hash
	Hash []byte
	// Nonce 随机数(挖矿)
	Nonce int
}



// 创建新的区块并返回(工厂方法)
func NewBlock(data string, prevBlockHash []byte) *Block  {
	//创建区块
	block := &Block{Timestamp: time.Now().Unix(), PrevBlockHash: prevBlockHash, Data: []byte(data), Hash: []byte{}, Nonce: 0}

	// 将block作为参数，创建一个pow工作证明对象
	pow := NewProofOfWork(block)

	// Run()执行一次工作量证明
	nonce, hash := pow.Run()

	// 设置区块Hash
	block.Hash = hash[:]

	// 设置Nonce
	block.Nonce = nonce

	isValid := pow.Validate()

	fmt.Println("\n当前区块的有效性", isValid)

	// 返回区块
	return block;
}

// 创建创世区块，并返回创世区块
func NewGenesisBlock() *Block {
	return NewBlock("Genenis Blcok", []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}