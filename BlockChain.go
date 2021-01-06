package my_bitcoin

//step5 创建区链块
type BlockChain struct {
	Blocks []*Block
}

//step6: 创建区块链，带有创世区块
func CreateBlockChainWithGenesisBlock(data string) *BlockChain  {
	genesisBlock := CreateGenesisBlock(data)
	return &BlockChain{[]*Block{genesisBlock}}
}

//step7: 添加一个新的区块， 到区块链中
func (bc *BlockChain) AddBlockToBlockChain(data string, height int64, prevHash []byte) {
	newBlock := NewBlock(data, prevHash, height)
	bc.Blocks = append(bc.Blocks, newBlock)
}
