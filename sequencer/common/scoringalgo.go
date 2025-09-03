package common

import (
	"fmt"
	"math/big"
)

const (
	alpha   = 0.5
	maxIter = 100
	req_bal = 9
	cutoff  = 6
	// Scale factor for precision (1,000,000) - preserves 6 decimal places
	scaleFactor = 1000000
)

func CalculateScore(vouches [][]int, balances []*big.Int, scores []*big.Int) []*big.Int {
	n := len(balances)
	
	vouch_matrix := make([][]*big.Int, n)
	for i := 0; i < n; i++ {
		vouch_matrix[i] = make([]*big.Int, n)
		for j := 0; j < n; j++ {
			vouch_matrix[i][j] = big.NewInt(int64(vouches[i][j]))
		}
	}
	
	old_scores := make([]*big.Int, n)
	for i := 0; i < n; i++ {
		old_scores[i] = new(big.Int).Set(scores[i])
	}
	
	active_nodes := 0
	reqBalBig := big.NewInt(req_bal)
	for _, val := range balances {
		if val.Cmp(reqBalBig) >= 0 {
			active_nodes++
		}
	}
	
	new_scores := make([]*big.Int, n)
	for i := range new_scores {
		new_scores[i] = big.NewInt(0)
	}
	
	if active_nodes < cutoff {
		if active_nodes > 0 {
			numerator := big.NewInt(scaleFactor)
			denominator := big.NewInt(int64(active_nodes))
			sharePerNode := new(big.Int).Div(numerator, denominator)
			
			for i, val := range balances {
				if val.Cmp(reqBalBig) >= 0 {
					new_scores[i].Set(sharePerNode)
				}
			}
		}
	} else {
		A := computeA(vouch_matrix, balances)
		W := computeW(A)
		P := computeP(W, alpha, maxIter)
		Q := computeQ(P)
		J := computeJ(Q)
		s := computeS(J, old_scores)
		new_scores = computeY(J, s, W)
	}

	fmt.Println("New Scores:")
	for i, val := range new_scores {
		actualValue := new(big.Float).SetInt(val)
		actualValue.Quo(actualValue, big.NewFloat(float64(scaleFactor)))
		fmt.Printf("Node %d: %s (scaled: %s)\n", i, actualValue.Text('f', 6), val.String())
	}
	
	return new_scores
}

func computeA(vouchMatrix [][]*big.Int, balances []*big.Int) [][]*big.Int {
	n := len(balances)
	A := make([][]*big.Int, n)
	for i := range A {
		A[i] = make([]*big.Int, n)
		for j := range A[i] {
			A[i][j] = big.NewInt(0)
		}
	}

	reqBalBig := big.NewInt(req_bal)
	one := big.NewInt(1)
	
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if vouchMatrix[i][j].Cmp(one) == 0 &&
				vouchMatrix[j][i].Cmp(one) == 0 &&
				balances[i].Cmp(reqBalBig) >= 0 &&
				balances[j].Cmp(reqBalBig) >= 0 {
				A[i][j] = big.NewInt(1)
			}
		}
	}
	return A
}

func computeW(A [][]*big.Int) [][]*big.Int {
	n := len(A)
	W := make([][]*big.Int, n)
	for i := range W {
		W[i] = make([]*big.Int, n)
		for j := range W[i] {
			W[i][j] = big.NewInt(0)
		}
	}

	for i := 0; i < n; i++ {
		d := big.NewInt(0)
		for j := 0; j < n; j++ {
			d.Add(d, A[i][j])
		}
		
		for j := 0; j < n; j++ {
			delta := big.NewInt(0)
			if i == j {
				delta = big.NewInt(1)
			}
			
			if d.Sign() != 0 {
				term := new(big.Int).Mul(A[i][j], big.NewInt(scaleFactor))
				term.Div(term, d)
				
				deltaScaled := new(big.Int).Mul(delta, big.NewInt(scaleFactor))
				sum := new(big.Int).Add(deltaScaled, term)
				
				W[i][j].Div(sum, big.NewInt(2))
			} else {
				deltaScaled := new(big.Int).Mul(delta, big.NewInt(scaleFactor))
				W[i][j].Div(deltaScaled, big.NewInt(2))
			}
		}
	}
	return W
}

func computeP(W [][]*big.Int, alpha float64, maxIter int) [][]*big.Int {
	n := len(W)
	P := make([][]*big.Int, n)
	
	alphaScaled := big.NewInt(int64(alpha * float64(scaleFactor)))
	oneMinusAlphaScaled := big.NewInt(int64((1-alpha) * float64(scaleFactor)))
	
	for i := 0; i < n; i++ {
		P[i] = make([]*big.Int, n)
		for j := range P[i] {
			P[i][j] = big.NewInt(0)
		}
		
		s := make([]*big.Int, n)
		for j := range s {
			s[j] = big.NewInt(0)
		}
		s[i] = big.NewInt(scaleFactor) // 1.0 scaled
		
		power := make([]*big.Int, n)
		for j := range power {
			power[j] = new(big.Int).Set(s[j])
		}
		
		scale := new(big.Int).Set(alphaScaled)
		
		for j := range P[i] {
			P[i][j].Mul(power[j], scale)
			P[i][j].Div(P[i][j], big.NewInt(scaleFactor))
		}
		
		for t := 1; t < maxIter; t++ {
			power = matVecMul(power, W)
			
			scale.Mul(scale, oneMinusAlphaScaled)
			scale.Div(scale, big.NewInt(scaleFactor))
			
			for j := range P[i] {
				scaledPower := new(big.Int).Mul(power[j], scale)
				scaledPower.Div(scaledPower, big.NewInt(scaleFactor))
				P[i][j].Add(P[i][j], scaledPower)
			}
		}
	}
	return P
}

func computeQ(P [][]*big.Int) [][]int {
	n := len(P)
	Q := make([][]int, n)
	for i := 0; i < n; i++ {
		idx := make([]int, n)
		for j := 0; j < n; j++ {
			idx[j] = j
		}
		for j := 0; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if P[i][idx[j]].Cmp(P[i][idx[k]]) < 0 {
					idx[j], idx[k] = idx[k], idx[j]
				}
			}
		}
		Q[i] = idx
	}
	return Q
}

func computeJ(Q [][]int) [][][]int {
	n := len(Q)
	J := make([][][]int, n)
	for i := 0; i < n; i++ {
		J[i] = make([][]int, n)
		for k := 0; k < n; k++ {
			J[i][k] = make([]int, n)
			J[i][k][i] = 1
			for m := 0; m < k; m++ {
				J[i][k][Q[i][m]] = 1
			}
		}
	}
	return J
}

func computeS(J [][][]int, old_scores []*big.Int) [][]int {
	n := len(J)
	s := make([][]int, n)
	for i := 0; i < n; i++ {
		s[i] = make([]int, n)
		for k := 0; k < n; k++ {
			sumJ := big.NewInt(0)
			sumC := big.NewInt(0)
			
			for j := 0; j < n; j++ {
				if J[i][k][j] == 1 {
					sumJ.Add(sumJ, old_scores[j])
				} else {
					sumC.Add(sumC, old_scores[j])
				}
			}
			if sumJ.Cmp(sumC) <= 0 {
				s[i][k] = 1
			}
		}
	}
	return s
}

func computeY(J [][][]int, s [][]int, W [][]*big.Int) []*big.Int {
	n := len(J)
	y := make([]*big.Int, n)	
	infValue := new(big.Int).Mul(big.NewInt(scaleFactor), big.NewInt(1000000000)) // Very large number
	for i := 0; i < n; i++ {
		y[i] = new(big.Int).Set(infValue)
		
		for k := 0; k < n; k++ {
			if s[i][k] == 1 {
				jVec := J[i][k]
				jComp := make([]*big.Int, n)
				for m := 0; m < n; m++ {
					jComp[m] = big.NewInt(scaleFactor - int64(jVec[m])*scaleFactor)
				}
				jVecBig := make([]*big.Int, n)
				for m := 0; m < n; m++ {
					jVecBig[m] = big.NewInt(int64(jVec[m]) * scaleFactor)
				}
				
				temp := matVecMul(jComp, W)
				num := dotVec(temp, jVecBig)
				count := big.NewInt(int64(sum(jVec)))
				
				if count.Sign() > 0 {
					val := new(big.Int).Div(num, count)
					if val.Cmp(y[i]) < 0 {
						y[i].Set(val)
					}
				}
			}
		}		
		if y[i].Cmp(infValue) >= 0 {
			y[i].SetInt64(0)
		}
	}
	return y
}

func matVecMul(v []*big.Int, M [][]*big.Int) []*big.Int {
	res := make([]*big.Int, len(M))
	for i := 0; i < len(M); i++ {
		res[i] = big.NewInt(0)
		for j := 0; j < len(v); j++ {
			term := new(big.Int).Mul(v[j], M[j][i])
			term.Div(term, big.NewInt(scaleFactor)) // Account for scaling
			res[i].Add(res[i], term)
		}
	}
	return res
}

func dotVec(a, b []*big.Int) *big.Int {
	sum := big.NewInt(0)
	for i := range a {
		term := new(big.Int).Mul(a[i], b[i])
		term.Div(term, big.NewInt(scaleFactor)) // Account for scaling
		sum.Add(sum, term)
	}
	return sum
}

func sum(v []int) int {
	s := 0
	for _, x := range v {
		s += x
	}
	return s
}
