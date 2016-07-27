package main

import (
	"fmt"

	"github.com/RobertChin84/QuantGo/algorithms"
	"github.com/RobertChin84/QuantGo/mathutils"
)

func main() {
	bm := new(algorithms.Blackmodel)
	/*
		  S= $210.59, K= $205 t = 4 days r = 0.2175% s = 14.04%
			  price = 37
			  S0, K, sigma, r, T
	*/
	bm.Init(210.59, 205, 0.1404, 0.002175, 0.01096)
	bm.Computeprice()

	norm := new(mathutils.Normal)

	fmt.Println(norm.Cdf(2))
}
