package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Начальные значения как big.Int
	f0 := big.NewInt(1)
	f1 := big.NewInt(3)

	// Массив для хранения нечетных значений
	var A []*big.Int

	// Проверяем начальные значения на нечетность
	if f0.Bit(0) == 1 { // проверка на нечетность через младший бит
		A = append(A, new(big.Int).Set(f0))
	}
	if f1.Bit(0) == 1 {
		A = append(A, new(big.Int).Set(f1))
	}

	// Вычисляем значения функции пока не наберем 40 нечетных значений
	prev2 := new(big.Int).Set(f0)
	prev1 := new(big.Int).Set(f1)
	n := 2

	for len(A) <= 39 {
		// f(n) = 5*f(n-1) + f(n-2)
		current := new(big.Int)
		temp := new(big.Int).Mul(prev1, big.NewInt(5))
		current.Add(temp, prev2)

		// Проверяем на нечетность через младший бит
		if current.Bit(0) == 1 {
			A = append(A, new(big.Int).Set(current))
		}

		// Обновляем для следующей итерации
		prev2.Set(prev1)
		prev1.Set(current)
		n++
	}

	fmt.Println(A[39])
}
