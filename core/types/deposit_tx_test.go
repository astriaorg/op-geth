package types

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestEncodeDepositTx(t *testing.T) {
	to := common.HexToAddress("0xa0Ee7A142d267C1f36714E4a8F75612F20a79720")
	depositTx := &DepositTx{
		SourceHash:          common.HexToHash("0x7113be8bbb6ff4bb99fae05639cf76cdecf5a1afbc033b9a01d8bb16b00b9a80"),
		From:                to,
		To:                  &to,
		Mint:                big.NewInt(10000000000000000),
		Value:               big.NewInt(10000000000000000),
		Gas:                 21000,
		IsSystemTransaction: false,
		Data:                []byte("nootwashere"),
	}
	bytes := &bytes.Buffer{}
	depositTx.encode(bytes)
	t.Log(hex.EncodeToString(bytes.Bytes()))
}
