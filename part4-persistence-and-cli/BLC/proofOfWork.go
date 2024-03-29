package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var (
	// 定义Nonce最大值
	maxNonce = math.MaxInt64
)
const targetBits = 20

type ProofOfWork struct {
	block *Block // 当前需要验证的区块
	target *big.Int // 大数据存储(挖矿难度)
}

// 数据拼接，返回字节数组
func (pow *ProofOfWork) prepareData(nonce int) []byte  {
	data := bytes.Join(
		[][]byte {
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(int64(pow.block.Timestamp)),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

// ProofOfWork对象的方法
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	fmt.Println()
	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		// Cmp 比较大小
		// -1 if hashInt < pow.target
		// 0 if hashInt == pow.target
		// 1 if hashInt > pow.target
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	return nonce, hash[:]
}

// 工厂方法
func NewProofOfWork(block *Block) *ProofOfWork  {
	target := big.NewInt(1)
	target.Lsh(target, uint(256 - targetBits)) // 左移targetBits位
	fmt.Println(target)
	pow := &ProofOfWork{block:block, target:target}
	
	return pow
}

// 验证当前的工作量证明有效性
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}

