package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	testAddr = common.HexToAddress("b94f5374fce5edbc8e2a8697c15331677e6ebf0b")

	//Legacy transaction without chainID
	rightvrsTx, _ = types.NewTransaction(
		3,
		testAddr,
		big.NewInt(10),
		2000,
		big.NewInt(1),
		common.FromHex("5544"),
	).WithSignature(
		types.HomesteadSigner{},
		common.Hex2Bytes("98ff921201554726367d2be8c804a7ff89ccf285ebc57dff8ae4c44b9c19ac4a8887321be575c8095f789dd4c743dfe42c1820f9231f98a962b210e3ac2452a301"),
	)
)

func convertToNewFormat(hex string) []byte {
	return common.FromHex(hex)
}

func main() {
	v, r, s := rightvrsTx.RawSignatureValues()
	fmt.Printf("%+v\n", rightvrsTx.Value())
	fmt.Printf("%+v\n", v)
	fmt.Printf("%+v\n", r)
	fmt.Printf("%+v\n", s)

	fmt.Printf("%+v\n", rightvrsTx.Hash())
	a := convertToNewFormat("4da580fd2e4c04f328d9f947ecf356411eb8e4a3a5c745f383b3ccd79c36a8d4")
	fmt.Println(a)
}
