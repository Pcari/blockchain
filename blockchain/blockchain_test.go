package blockchain

import (
	"testing"
	"time"
	"fmt"
	"encoding/hex"
)

func TestBlocks(t *testing.T) {
	start := time.Now()

	numWorker := uint64(2)
	b00 := Block{}
	b00 = Initial(2)
	b00.Mine(numWorker)

	fmt.Println("TESTING BLOCKCHAIN")
	fmt.Println(b00)
	testBlockChain1 := Blockchain{}
	testBlockChain1.Chain = make([]Block, 1)

	//where INDEX PROBLEM IS OCCURING
	testBlockChain1.Chain[0] = b00

	fmt.Println(len(testBlockChain1.Chain))

	b01 := b00.Next("secondblock")
	b01.Mine(numWorker)
	fmt.Println(b01)

	testBlockChain1.Add(b01)
	fmt.Println(testBlockChain1)
	fmt.Println(len(testBlockChain1.Chain))
	fmt.Println(testBlockChain1.Chain[0].Hash)
	fmt.Println(testBlockChain1.IsValid())

	//b00.Mine(2)
	//b01 := b00.Next("message")
	//b01.Mine(2)
	//fmt.Println(b00.Index)
	//fmt.Println(b00.Difficulty)
	//fmt.Println(b00.Data)
	//fmt.Println(b00.Hash)
	//fmt.Println(hex.EncodeToString(b00.Hash))
	//fmt.Println(b00.PrevHash)
	//fmt.Println(b00.Proof)
	//fmt.Println("b01:")
	//fmt.Println(b01.Index)
	//fmt.Println(b01.Difficulty)
	//fmt.Println(b01.Data)
	//fmt.Println(b01.Hash)
	//b01.CalcHash()
	//fmt.Println(b01.PrevHash)
	//fmt.Println(b01.Proof)

	end := time.Now()
	fmt.Println("Time took: ", end.Sub(start))
}

func TestBlockChain(t *testing.T) {







	start := time.Now()
	numWorker := uint64(100)
	difficulty := uint8(3)
	b0 := Initial(difficulty)
	b0.Mine(numWorker)
	fmt.Println(b0.Proof, hex.EncodeToString(b0.Hash))
	b1 := b0.Next("this is an interesting message")
	b1.Mine(numWorker)
	fmt.Println(b1.Proof, hex.EncodeToString(b1.Hash))
	b2 := b1.Next("this is not interesting")
	b2.Mine(numWorker)
	fmt.Println(b2.Proof, hex.EncodeToString(b2.Hash))
	end := time.Now()
	fmt.Println("Time took: ", end.Sub(start))
}
