package common

import (
	"crypto/ecdsa"
	"fmt"
	"strings"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func AddressFromPrivKeyHex(privateKeyHex string) (ethCommon.Address, error) {
	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")
	priv, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return ethCommon.Address{}, err
	}
	pub, ok := priv.Public().(*ecdsa.PublicKey)
	if !ok {
		return ethCommon.Address{}, fmt.Errorf("public key type assertion failed")
	}
	return crypto.PubkeyToAddress(*pub), nil
}
