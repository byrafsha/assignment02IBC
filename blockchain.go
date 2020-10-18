package assignment02IBC

import (
	"crypto/sha256"
	"fmt"
)

const miningReward = 100
const rootUser = "Satoshi"

type Block struct {
	Spender     map[string]int
	Receiver    map[string]int
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

func CalculateBalance(userName string, chainHead *Block) int {
	currBlock := chainHead
	bal := 0
	if chainHead == nil {
		return 0
	} else {
		for {
			for key, val := range currBlock.Receiver {
				if key == userName {
					bal += val
				}
			}
			for key, val := range currBlock.Spender {
				if key == userName {
					bal -= val
				}
			}

			if currBlock.PrevPointer == nil {
				return bal
			} else {
				currBlock = currBlock.PrevPointer
				fmt.Println(">>>>>")
			}
		}
	}
}

func CalculateHash(inputBlock *Block) string {
	//concatenate the entire blck in a string in order to calculate joint hash
	var spndr string
	var rcvr string
	for key, val := range inputBlock.Spender {
		spndr = fmt.Sprintf("%s=%v", key, val)
	}
	for key, val := range inputBlock.Receiver {
		rcvr = fmt.Sprintf("%s=%v", key, val)
	}
	entireBlock := fmt.Sprintf(spndr, rcvr, inputBlock.Receiver, inputBlock.PrevHash)
	//calculate hash
	ebHash := sha256.Sum256([]byte(entireBlock))
	inputBlock.CurrentHash = fmt.Sprintf("%x", ebHash)

	fmt.Println("Hash calculated and stored")
	return fmt.Sprintf("%x", ebHash)
}

func InsertBlock(spendingUser string, receivingUser string, miner string, amount int, chainHead *Block) *Block {
	//for verification of miner
	if miner != rootUser {
		return chainHead
	}

	var newBlock *Block = new(Block)
	newBlock.Spender = make(map[string]int)
	newBlock.Spender[spendingUser] = amount

	newBlock.Receiver = make(map[string]int)
	newBlock.Receiver[miner] = miningReward
	newBlock.Receiver[receivingUser] = amount

	//for verification of Balance
	if CalculateBalance(spendingUser, chainHead) < amount {
		return chainHead
	}

	if chainHead == nil {
		newBlock.PrevHash = ""
		newBlock.PrevPointer = nil
	} else {
		newBlock.PrevHash = chainHead.CurrentHash
		newBlock.PrevPointer = chainHead
	}

	newBlock.CurrentHash = CalculateHash(newBlock)

	//update chainhead
	chainHead = newBlock

	fmt.Println("\nNew block inserted with the following spenders and receivers: ")
	fmt.Println("Spender:", newBlock.Spender)
	fmt.Println("Receivers:", newBlock.Receiver, "\n")
	return newBlock
}

func ListBlocks(chainHead *Block) {

	currBlock := chainHead
	x := 0
	if chainHead != nil {
		for {
			fmt.Println("\nBlock:", x)
			fmt.Println(currBlock.Spender, currBlock.Receiver) //, currBlock.PrevPointer, currBlock.PrevHash, currBlock.CurrentHash)
			x += 1

			if currBlock.PrevPointer == nil {
				break
			} else {
				currBlock = currBlock.PrevPointer
				fmt.Println(">>>>>")
			}
		}
	}
}

func VerifyChain(chainHead *Block) {
	currBlock := chainHead
	if chainHead != nil {
		for {
			verHash := CalculateHash(currBlock.PrevPointer)
			if currBlock.PrevHash != verHash {
				fmt.Println("Uh-oh, blockchain has been compromised!")
				break
			} else {
				currBlock = currBlock.PrevPointer
			}
			println("Blockchain verified successfully!")
		}
	}
}
