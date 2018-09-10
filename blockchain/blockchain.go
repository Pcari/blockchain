package blockchain

import (


				"encoding/hex"
)

type Blockchain struct {
	Chain []Block
}

func (chain *Blockchain) Add(blk Block) {
	if !blk.ValidHash() {
		panic("adding block with invalid hash")
	}
	chain.Chain = append(chain.Chain, blk)
	//blockLength := len(chain.Chain)
	//fmt.Println("testing b lockchain length: %v", blockLength)
	//prevHash := chain.Chain[blockLength-1].Hash

	//fmt.Println("testing b lockchain length: %v", blockLength)
	/*
	//fmt.Println("HELLO???????????????")
	blockLength := len(chain.Chain)
	fmt.Println("Testing blockchain length: %v ", blockLength)
	if blockLength == 0 {
		newBlock := Initial(blk.Difficulty)
		//test with index 0 instead of append
		chain.Chain = append(chain.Chain, newBlock)
	} else {
		chain.Chain = append(chain.Chain, blk)
		prevHash := chain.Chain[blockLength-1].Hash
		blk.Generation = uint64(blockLength)
		blk.Difficulty = chain.Chain[blockLength-1].Difficulty
		blk.PrevHash = prevHash
		blk.Hash = blk.CalcHash()


		//Generation uint64 // index
		//Difficulty uint8  // num of bytes
		//Data       string
		//PrevHash   []byte
		//Hash       []byte
		//Proof      uint64
	}
*/


}

func (chain Blockchain) IsValid() bool {
	// TODO
	//sumOfDiff := uint8(0)
	//The initial block has previous hash all null bytes and is generation zero.
	if chain.Chain[0].Generation == 0 {
		//fmt.Println("length of prevHash: %v", len(chain.Chain[0].PrevHash))
		for i := 0; i < len(chain.Chain[0].PrevHash); i++ {

			if chain.Chain[0].PrevHash[i] != 0 {
				return false
			}

		}
		//return true
	}


	//	Each block has the same difficulty value.
	for i := 1; i < len(chain.Chain); i++ {
		if chain.Chain[i-1].Difficulty != chain.Chain[i].Difficulty {
			return false
		}
		//	Each block has a generation value that is one more than the previous block.
		if chain.Chain[i-1].Generation != (chain.Chain[i].Generation - 1) {
			return false
		}
		//	Each block's previous hash matches the previous block's hash.
		if hex.EncodeToString(chain.Chain[i-1].Hash) != hex.EncodeToString(chain.Chain[i].PrevHash) {
			return false
		}
		//	Each block's hash value actually matches its contents.
		if hex.EncodeToString(chain.Chain[i].CalcHash()) != hex.EncodeToString(chain.Chain[i].Hash) {
			return false
		}

		for j := uint8(0); j < chain.Chain[i].Difficulty; j++ {
			if chain.Chain[i].Hash[len(chain.Chain[i].Hash) - 1 - i] != 0 {
				return false
			}
		}
	}

	//	Each block's hash value ends in difficulty null bytes.
	return true
}
