package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func convertToBytesFromHash(hash common.Hash) []byte {
	new := make([]byte, 32)
	copy(new, hash[:])
	return new
}

func changeBytesToHash(bytes []byte) *common.Hash {
	var hash common.Hash
	copy(hash[:], bytes)
	return &hash
}

func changeByteContent(bytes []byte) []byte {
	//0x656765 - 3 bytes: which is my name "eges"
	bytes[0] = 0x65 //E
	bytes[1] = 0x67 //G
	bytes[2] = 0x65 //E

	return bytes
}

func changeTxHash() *common.Hash {
	tx_hash := common.HexToHash("0xdf4e14abd904447b9aec870b40e9fd620df39c7a9e740d40a3096de2d644e23e")
	bytes := convertToBytesFromHash(tx_hash)
	changed := changeByteContent(bytes)
	hash := changeBytesToHash(changed)
	fmt.Println(hash)
	return hash
}

func main() {
	changeTxHash()
	fmt.Println(testRlpHash())
	fmt.Println(testChangedRlpHash())
}
