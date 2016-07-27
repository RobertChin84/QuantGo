package algorithms

import (
	"errors"
	"fmt"
	"math"

	"github.com/RobertChin84/QuantGo/mathutils"
)

/*TODO
* modify to be general i.e. black-76 and B-S model
* different underlying
 */
type Blackmodel struct {
	/*
	 * This is the base closed form solution for the Black-76 mode;
	 * V0 = A*N(d1) -KN(d2)
	 */
	d1, d2    float64
	sigma, r  float64 // volatility and risk free rate
	T         float64 // time to maturity
	S, K      float64 // current price of underlying strike on underlying
	Call, Put float64
	norm      *mathutils.Normal
}

func (b *Blackmodel) Init(S0, K, sigma, r, T float64) error {
	b.S = S0
	b.K = K
	b.sigma = sigma
	b.r = r
	b.T = T
	if b.S < 0 || b.K < 0 {
		return errors.New("Blackmodel: make sure S0 and K are > 0")
	}
	b.computeD()
	return nil
}
func (b *Blackmodel) computeD() {
	b.d1 = (math.Log(b.S/b.K) + (b.r+0.5*math.Pow(b.sigma, 2))*b.T) / (b.sigma * math.Pow(b.T, .5))
	b.d2 = b.d1 - b.sigma*math.Pow(b.T, .5)
	fmt.Println(b.d1, b.d2)
}

func (b *Blackmodel) Computeprice() {
	b.Call = ((b.S)*b.norm.Cdf(b.d1) - (b.K)*b.norm.Cdf(b.d2)) * math.Exp(-b.r*b.T)
	b.Put = (-(b.S)*b.norm.Cdf(-b.d1) + (b.K)*b.norm.Cdf(-b.d2)) * math.Exp(-b.r*b.T)
}
