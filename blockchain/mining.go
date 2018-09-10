package blockchain

import (
	"a3/work_queue"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

type miningWorker struct {
	Block Block
	Start uint64
	End   uint64
	//Taken bool
}

func (w miningWorker) Run() interface{} {
	miningResult := MiningResult{}
	blk := w.Block
	start := w.Start
	end := w.End

	prevHash := hex.EncodeToString(blk.PrevHash)
	index := fmt.Sprint(blk.index)
	difficulty := blk.Difficulty
	data := blk.Data
	//fmt.Println("Running")
	for p := start; p < end; p++ {
		proof := fmt.Sprint(p)
		s := []string{prevHash, index, fmt.Sprint(difficulty), data, proof}
		hashInput := strings.Join(s, ":")
		h := sha256.New()
		h.Write([]byte(hashInput))
		hashValue := h.Sum(nil)
		length := len(hashValue)
		sum := 0
		//fmt.Println(hashValue)
		for _, b := range hashValue[(length - int(difficulty)):] {
			sum += int(b)
		}
		if sum == 0 {
			//fmt.Println("Sum is zero")
			miningResult.Found = true
			miningResult.Proof = p
			return miningResult
		}

	}

	return miningResult
}

type MiningResult struct {
	Proof uint64 // proof-of-work value, if found.
	Found bool   // true if valid proof-of-work was found.
}

// Mine the range of proof values, by breaking up into chunks and checking
// "workers" chunks concurrently in a work queue. Should return shortly after a result
// is found.
func (blk Block) MineRange(start uint64, end uint64, workers uint64, chunks uint64) MiningResult {

	length := end - start
	numChunks := uint64(length / chunks)
	if numChunks == 0 {
		numChunks = 1
		chunks = end
	}
	queue := work_queue.Create(uint(workers), uint(numChunks))
	mr := MiningResult{}

	found := false

	for i := uint64(0); i < numChunks; i++ {
		begin := i * chunks
		finish := begin + chunks
		//fmt.Println(begin, " ", finish)
		mw := miningWorker{}
		mw.Block = blk
		mw.Start = begin
		mw.End = finish
		queue.Enqueue(mw)
	}
	results := queue.Results

	//fmt.Println("End of Queueing")
	for i := uint64(0); i < numChunks && !found; {
		//fmt.Println(len(results))
		result := <-results
		//fmt.Println(result)
		if result != nil {
			i++
		}
		mr = result.(MiningResult)

		if mr.Found {
			found = true
			queue.Shutdown()
			return mr
		}
	}
	return mr
}

// Call .MineRange with some reasonable values that will probably find a result.
// Good enough for testing at least. Updates the block's .Proof and .Hash if successful.
func (blk *Block) Mine(workers uint64) bool {
	//fmt.Println("Mine")
	reasonableRangeEnd := uint64(4 * 1 << (8 * blk.Difficulty)) // 4 * 2^(bits that must be zero)
	mr := blk.MineRange(0, reasonableRangeEnd, workers, 4321)
	if mr.Found {
		//fmt.Println("Setting Proof")
		blk.SetProof(mr.Proof)
	}
	return mr.Found
}
