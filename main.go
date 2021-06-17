package main

import (
	"fmt"
	"math"
)

func main() {
	pva := 1200.0
	i := 2.0
	n := 6
	pmt := MontlyPayment(pva, i, n)
	fmt.Println(pmt)
}

// calculate to calculate monthly payment, pva (Present Value of Annuities) is loan amount, i is interest (percent/month), 6 = period (month), pmt is payment
func MontlyPayment(pva, i float64, n int) (pmt float64) {
	x1 := math.Pow(1+i, float64(n))
	fmt.Println(x1)
	x2 := (1 / x1)
	fmt.Println(x2)
	pmt = pva / ((1 - x2) / i)
	return
}
