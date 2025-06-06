package common

import "math/big"

// RollupForgeBatchArgs are the arguments to the ForgeBatch function in the Rollup Smart Contract
type RollupForgeBatchArgs struct {
	NewAccountRoot *big.Int
	NewVouchRoot   *big.Int
	NewScoreRoot   *big.Int
	ProofA         [2]*big.Int
	ProofB         [2][2]*big.Int
	ProofC         [2]*big.Int
}
