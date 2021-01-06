package main

import "fmt"

func main()  {
	blockChain := CreateBlockChainWithGenesisBlock("Genesis Block..")
	blockChain.AddBlockToBlockChain("Send 100RMB To Wangergou", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	blockChain.AddBlockToBlockChain("Send 500RMB To licai", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	blockChain.AddBlockToBlockChain("Send 200RMB To lijuan", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)

	fmt.Println(blockChain)
}
