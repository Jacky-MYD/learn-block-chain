package BLC

type BlockChain struct {
	Blocks []*Block // 储存有序的区块
}

// 新增区块
func (blockChain *BlockChain) AddBlock(data string)  {
	// 1. 创建新的Block
	preBlock := blockChain.Blocks[len(blockChain.Blocks)- 1]
	newBlock := NewBlock(data, preBlock.Hash)

	// 2. 将区块添加到Blocks里面
	blockChain.Blocks = append(blockChain.Blocks, newBlock)
}

// 创建一个带有创世区块的区块链
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}