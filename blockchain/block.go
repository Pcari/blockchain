package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

type Block struct {
	Index      uint64 // index
	Difficulty uint8  // num of bytes
	Data       string
	PrevHash   []byte
	Hash       []byte
	Proof      uint64
}

// Create new initial (Index 0) block.
func Initial(difficulty uint8) Block {
	b := Block{}
	b.Index = 0
	b.Difficulty = difficulty
	b.Data = ""
	b.PrevHash = make([]byte, 32)
	//b.Hash = b.CalcHash()
	return b

}

// Create new block to follow this block, with provided data.
func (prev_block Block) Next(data string) Block {

	b := Block{}
	b.Index = 1 + prev_block.Index
	b.Difficulty = prev_block.Difficulty
	b.Data = data
	b.PrevHash = prev_block.Hash
	//b.Hash = b.CalcHash()
	//b.Hash = hex.EncodeToString(b.Hash)
	return b

}

// Calculate the block's hash.
func (blk Block) CalcHash() []byte {

	prevHash := hex.EncodeToString(blk.PrevHash)
	index := fmt.Sprint(blk.Index)
	difficulty := fmt.Sprint(blk.Difficulty)
	data := blk.Data
	proof := fmt.Sprint(blk.Proof)

	s := []string{prevHash, index, difficulty, data, proof}
	hashInput := strings.Join(s, ":")

	h := sha256.New()
	h.Write([]byte(hashInput))
	return h.Sum(nil)

}

// Is this block's hash valid?
func (blk Block) ValidHash() bool {
	length := uint8(len(blk.Hash))
	sum := 0
	for _, b := range blk.Hash[(length - blk.Difficulty):] {
		sum += int(b)
	}
	return sum == 0
}

// Set the proof-of-work and calculate the block's "true" hash.
func (blk *Block) SetProof(proof uint64) {
	blk.Proof = proof
	blk.Hash = blk.CalcHash()
}
