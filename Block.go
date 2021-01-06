package my_bitcoin

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Height int64
	PrevBlockHash []byte
	Data [] byte
	TimeStamp int64
	Hash []byte
}

//step2: 创建新的区块
func NewBlock(data string, provBlockHash [] byte, height int64) *Block {
	block := &Block{height, provBlockHash, []byte(data), time.Now().Unix(), nil}
	block.SetHash()
	return block
}
// step3: 设置区块的hash
func (block *Block) SetHash()  {
	//1. 将高度转化为字节数组
	heightBytes:= IntToHex(block.Height)
	timeString := strconv.FormatInt(block.TimeStamp, 2)
	timeBytes:= [] byte(timeString)
	blockBytes:= bytes.Join([][]byte{
		heightBytes,
		block.PrevBlockHash,
		block.Data,
		timeBytes}, []byte{})
	//生成hash值
	hash:= sha256.Sum256(blockBytes)
	block.Hash = hash[:]

}

//step4: 创建创世区块
func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, make([] byte, 32, 32),0)
}
