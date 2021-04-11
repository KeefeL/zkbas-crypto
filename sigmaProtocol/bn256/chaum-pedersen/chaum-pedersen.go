package chaum_pedersen

import (
	"Zecrey-crypto/ecc/zbn256"
	"Zecrey-crypto/ffmath"
	"github.com/consensys/gurvy/bn256"
	"math/big"
)

// prove v = g^{\beta} \and w = u^{\beta}
func Prove(beta *big.Int, g, u, v, w *bn256.G1Affine) (z *big.Int, Vt, Wt *bn256.G1Affine) {
	// betat \gets_R Z_p
	betat := zbn256.RandomValue()
	// Vt = g^{betat}
	Vt = zbn256.G1ScalarMult(g, betat)
	// Wt = u^{betat}
	Wt = zbn256.G1ScalarMult(u, betat)
	// c = H(Vt,Wt,v,w)
	c := HashChaumPedersen(Vt, Wt, v, w)
	// z = betat + beta * c
	z = ffmath.AddMod(betat, ffmath.MultiplyMod(c, beta, Order), Order)
	return z, Vt, Wt
}

func Verify(z *big.Int, g, u, Vt, Wt, v, w *bn256.G1Affine) bool {
	// c = H(Vt,Wt,v,w)
	c := HashChaumPedersen(Vt, Wt, v, w)
	// check if g^z = Vt * v^c
	l1 := zbn256.G1ScalarMult(g, z)
	r1 := zbn256.G1Add(Vt, zbn256.G1ScalarMult(v, c))
	// check if u^z = Wt * w^c
	l2 := zbn256.G1ScalarMult(u, z)
	r2 := zbn256.G1Add(Wt, zbn256.G1ScalarMult(w, c))
	return l1.Equal(r1) && l2.Equal(r2)
}
