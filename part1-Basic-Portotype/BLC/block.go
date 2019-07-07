package BLC

import (
	"bytes"
	"crypto/sha256"
	"strconv"
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
}

func (block *Block) SetHash() {
	// 1. 将时间戳转为字节数组
	//（1）将init64转字符串
	timeString := strconv.FormatInt(block.Timestamp, 2)
	// fmt.Println(timeString)
	//（2）将字符串转字节数组
	timestamp := []byte(timeString)
	// fmt.Println(timestamp)

	// 2. 将除了Hash以外的其他属性，以字节数组的形式全拼接起来
	// fmt.Println([][]byte{block.PrevBlockHash, block.Data, timestamp})
	headers := bytes.Join([][]byte{block.PrevBlockHash, block.Data, timestamp}, []byte{})
	// fmt.Println(headers)

	// 3. 将拼接起来的数据进行256 hash
	hash := sha256.Sum256(headers)

	// 4. 将hash赋值给Hash属性字段
	block.Hash = hash[:]
}


// 创建新的区块并返回(工厂方法)
func NewBlock(data string, prevBlockHash []byte) *Block  {
	//创建区块
	block := &Block{Timestamp: time.Now().Unix(), PrevBlockHash: prevBlockHash, Data: []byte(data), Hash: []byte{}}
	// 设置当前区块的Hash值
	block.SetHash()
	// 返回区块
	return block;
}

// 创建创世区块，并返回创世区块
func NewGenesisBlock() *Block {
	return NewBlock("Genenis Blcok", []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}