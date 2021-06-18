package main

import (
	"fmt"
	"math"
)

func main() {
	pva := 1200.0
	i := 0.02
	n := 6
	pmt := MontlyPayment(pva, i, n)
	fmt.Printf("loan amount: %0.2f\ninterest: %0.2f/month\nperiod(month): %d\npayment: %0.2f\n", pva, i, n, pmt)
	total := pmt * float64(n)
	fmt.Printf("total: %0.2f\n", total)
	fmt.Printf("total interest: %0.2f\n", total-pva)
}

func MontlyPayment(pva, i float64, n int) (pmt float64) {
	return pva / ((1 - (1 / math.Pow(1+i, float64(n)))) / i)
}
