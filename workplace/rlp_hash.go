package main

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

var hasherPool = sync.Pool{
	New: func() interface{} { return sha3.NewLegacyKeccak256() },
}

func rlpHash(x interface{}) (h common.Hash) {
	sha := hasherPool.Get().(crypto.KeccakState)
	defer hasherPool.Put(sha)
	sha.Reset()
	rlp.Encode(sha, x)
	sha.Read(h[:])

	bytes := convertToBytesFromHash(h)
	changed := changeByteContent(bytes)

	return changeBytesToHash(changed)
}

func testRlpHash() common.Hash {
	list := [3]string{"Ege", "Ethereum", "Bitcoin"}
	hash := rlpHash(list)

	return hash
}

/*
func testChangedRlpHash() *common.Hash {
	list := [3]string{"Ege", "Ethereum", "Bitcoin"}
	hash := rlpHash(list)

	bytes := convertToBytesFromHash(hash)
	changed := changeByteContent(bytes)

	return changeBytesToHash(changed)
}
*/
