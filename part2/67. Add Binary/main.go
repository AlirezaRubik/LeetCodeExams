package main

import (
	"math/big"
)
func addBinary(a string, b string) string {
	num1, _ := new(big.Int).SetString(a, 2)
	num2, _ := new(big.Int).SetString(b, 2)
	num1.Add(num1, num2)
	return num1.Text(2)
}