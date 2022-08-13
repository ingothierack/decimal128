package decimal128_test

import (
	"fmt"

	"github.com/woodsbury/decimal128"
)

func ExampleNew() {
	fmt.Println(decimal128.New(3, -2))
	fmt.Println(decimal128.New(3, 0))
	fmt.Println(decimal128.New(3, 2))
	// Output:
	// 0.03
	// 3
	// 300
}

func ExampleDecimal_Add() {
	x := decimal128.New(3, 0)
	y := decimal128.New(2, -1)
	fmt.Printf("%g + %g = %g\n", x, y, x.Add(y))
	// Output:
	// 3 + 0.2 = 3.2
}

func ExampleDecimal_Cmp() {
	x := decimal128.New(1, 0)
	y := decimal128.New(2, 0)
	r := x.Cmp(y)
	fmt.Printf("%g < %g = %t\n", x, y, r.Less())
	fmt.Printf("%g == %g = %t\n", x, y, r.Equal())
	fmt.Printf("%g > %g = %t\n", x, y, r.Greater())
	// Output:
	// 1 < 2 = true
	// 1 == 2 = false
	// 1 > 2 = false
}

func ExampleDecimal_Mul() {
	x := decimal128.New(3, 0)
	y := decimal128.New(2, -1)
	fmt.Printf("%g * %g = %g\n", x, y, x.Mul(y))
	// Output:
	// 3 * 0.2 = 0.6
}

func ExampleDecimal_Quo() {
	x := decimal128.New(152, -1)
	y := decimal128.New(5, 0)
	fmt.Printf("%g / %g = %g\n", x, y, x.Quo(y))
	// Output:
	// 15.2 / 5 = 3.04
}

func ExampleDecimal_QuoRem() {
	x := decimal128.New(152, -1)
	y := decimal128.New(5, 0)
	q, r := x.QuoRem(y)
	fmt.Printf("%g / %g = (%g, %g)\n", x, y, q, r)
	// Output:
	// 15.2 / 5 = (3, 0.2)
}

func ExampleDecimal_Round() {
	x := decimal128.New(123456, -3)
	fmt.Println("unrounded:", x)
	fmt.Println("+2 places:", x.Round(2, decimal128.DefaultRoundingMode))
	fmt.Println(" 0 places:", x.Round(0, decimal128.DefaultRoundingMode))
	fmt.Println("-2 places:", x.Round(-2, decimal128.DefaultRoundingMode))
	// Output:
	// unrounded: 123.456
	// +2 places: 123.46
	//  0 places: 123
	// -2 places: 100
}

func ExampleDecimal_Sub() {
	x := decimal128.New(3, 0)
	y := decimal128.New(2, -1)
	fmt.Printf("%g - %g = %g\n", x, y, x.Sub(y))
	// Output:
	// 3 - 0.2 = 2.8
}
