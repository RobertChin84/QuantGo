package mathutils

import "math"

type BaseDistribution struct {
	/*
	 * This base struct will be inherited to all distributions
	 */
	rgen *Randomgen // random number generator
	// density
	// cumaltive
}

type Normal struct {
	BaseDistribution
}

func (n *Normal) Density(z float64) float64 {

	res := math.Exp(-0.5*z*z) / (math.Sqrt(math.Pi * 2.0))
	return res
}

func (n *Normal) Cdf(z float64) float64 {
	a1 := 0.319381530
	a2 := -0.356563782
	a3 := 1.781477937
	a4 := -1.821255978
	a5 := 1.330274429
	Culmative := 0.0
	k := 1.0 / (1.0 + 0.2316419*z)
	if z >= 0.0 {
		Culmative = 1.0 - n.Density(z)*(a1*k+a2*math.Pow(k, 2)+a3*math.Pow(k, 3)+
			a4*math.Pow(k, 4)+a5*math.Pow(k, 5))
	} else {
		Culmative = 1.0 - n.Cdf(-z)
	}
	return Culmative
}
