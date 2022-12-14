package main

import (
	"fmt"
	"math"
)

func main() {
	var oper, num1, num2 = getInput()
	var outputName string = "Output"
	var _ []bool
	var xylen int = int(math.Ceil(math.Log2(float64(max(num1, num2))))+float64(0.1)) + 1
	var x, y, output []bool = make([]bool, xylen), make([]bool, xylen), make([]bool, xylen+1)
	x = intToBl(num1, xylen)
	y = intToBl(num2, xylen)
	switch oper {
	case "a":
		output = add(x, y)
		outputName = "Sum"
	case "m":
		output = multiply(x, y)
		outputName = "Product"
	case "s":
		output = subtract(x, y)
		outputName = "Difference"
	case "d":
		output, _ = divide(x, y)
		outputName = "Quotient"
	case "mod":
		_, output = divide(x, y)
		outputName = "Remainder"
	case "e":
		output = pow(x, y)
		outputName = "Power"
	}
	fmt.Println("\n" + outputName + ": " + fmt.Sprint(blToBase2(output)) + "(" + fmt.Sprint(blToInt(output)) + ")")
}

func not(x bool) bool          { return !x }
func and(x bool, y bool) bool  { return x && y }
func or(x bool, y bool) bool   { return x || y }
func nand(x bool, y bool) bool { return not(and(x, y)) }
func nor(x bool, y bool) bool  { return not(or(x, y)) }
func xor(x bool, y bool) bool  { return or(and(x, not(y)), and(not(x), y)) }
func xnor(x bool, y bool) bool { return not(xor(x, y)) }

func add(x []bool, y []bool) []bool {
	var output []bool = make([]bool, len(x)+1)
	var nt, ft bool = false, false
	for i := len(x) - 1; i >= 0; i-- {
		ft = xor(xor(x[i], y[i]), nt)
		nt = or(or(and(x[i], nt), and(y[i], nt)), and(x[i], y[i]))

		output[i+1] = ft
		fmt.Println(blToBase2(output), ft, nt)
	}
	output[0] = nt
	return output
}
func multiply(x []bool, y []bool) []bool {
	var output, adj []bool = make([]bool, len(x)), []bool{}
	for i := 0; i < blToInt(y); i++ {
		if len(output) > len(x) {
			adj = make([]bool, len(output)-len(x))
			x = append(adj, x...)
		}
		output = add(output, x)
	}
	return output
}
func subtract(x []bool, y []bool) []bool {
	var output []bool = make([]bool, len(x))
	var nt, ft bool = false, false
	output[len(output)-1] = xor(x[len(x)-1], y[len(y)-1])
	nt = and(not(x[len(x)-1]), y[len(y)-1])
	for i := len(x) - 2; i >= 0; i-- {
		ft = xor(nt, xor(x[i], y[i]))
		nt = or(and(not(x[i]), y[i]), and(xnor(x[i], y[i]), nt))
		output[i] = ft
		fmt.Println(blToBase2(output), ft, nt)
	}
	return output
}
func divide(x []bool, y []bool) ([]bool, []bool) {
	var output []bool = make([]bool, len(x))
	for blToInt(x) >= blToInt(y) {
		x = subtract(x, y)
		output = intToBl(blToInt(output)+1, len(output))
	}

	return output, x
}
func pow(x []bool, y []bool) []bool {
	var output []bool = intToBl(1, len(x))
	for i := 0; i < blToInt(y); i++ {
		output = multiply(output, x)
	}
	return output
}

func intToBl(num int, l int) []bool {
	var strt int = int(math.Pow(float64(2), float64(l-1)) + float64(0.1))
	var lis []bool = make([]bool, l)
	for i := 0; i < l; i++ {
		if num >= strt {
			lis[i] = true
			num -= strt
		}
		strt /= 2
	}
	return lis
}
func blToInt(list []bool) int {
	var output, val int = 0, 1
	for i := len(list) - 1; i >= 0; i-- {
		if list[i] {
			output += val
		}
		val *= 2
	}
	return output
}
func blToBase2(list []bool) []int {
	var output []int = make([]int, len(list))
	for i, ival := range list {
		if ival {
			output[i] = 1
		} else {
			output[i] = 0
		}
	}

	for _, ival := range output {
		if ival == 1 {
			break
		}
		output = output[1:]
	}
	return output
}
func max(x int, y int) int {
	if y > x {
		return y
	}
	return x
}
func base2ToBl(list []int) []bool {
	var output []bool = make([]bool, len(list))
	for i, ival := range list {
		output[i] = (ival == 1)
	}
	return output
}
func validateBase2(input int) ([]bool, bool) {
	var output []bool =  make([]bool, int(math.Log10(float64(input))));
	var c int = len(output);
	for input >= 1{
		c--;
		if input % 10 != 0 && input % 10 != 0{
			return []bool(nil), false;
		}
		output[c] = ((input % 10) == 1);
		input -= input % 10;
		input /= 10;
	}
	return output, true;
}
func getInput() (string, int, int) {
	// operation codes: addition = a, subtraction = s, multiplication = m, division = d, modulus = mod, exponents = e
	var oper string
	for true {
		fmt.Println("vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv")
		fmt.Println("|    Operators    |Format |Constraints|")
		fmt.Println("|-----------------|-------|-----------|")
		fmt.Println("|    Addition     | a + b | a, b > 0  |")
		fmt.Println("|-----------------|-------|-----------|")
		fmt.Println("|   Subtraction   | a - b |a >= b > 0 |")
		fmt.Println("|-----------------|-------|-----------|")
		fmt.Println("| Multiplication  | a * b | a, b > 0  |")
		fmt.Println("|-----------------|-------|-----------|")
		fmt.Println("|    Division     | a / b |a >= b > 0 |")
		fmt.Println("|-----------------|-------|-----------|")
		fmt.Println("|     Modulo      |a mod b|a >= b > 0 |")
		fmt.Println("|-----------------|-------|-----------|")
		fmt.Println("|    Exponent     | a ^ b | a,b > 0   |")
		fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
		fmt.Println("Please enter an operator: ")
		fmt.Scan(&oper)
		if oper == "Multiplication" || oper == "Subtraction" || oper == "Addition" || oper == "Division" || oper == "Modulo" || oper == "Exponent" {
			switch oper {
			case "Multiplication":
				oper = "m"
			case "Subtraction":
				oper = "s"
			case "Addition":
				oper = "a"
			case "Division":
				oper = "d"
			case "Modulo":
				oper = "mod"
			case "Exponent":
				oper = "e"
			}
			break
		}
	}
	var selection string
	for true {
		fmt.Println("Would you like to enter your numbers in (Binary or Decimal)")
		fmt.Scan(&selection)
		if selection == "Decimal" || selection == "Binary" {
			break
		}
	}

	var num11, num22 int
	if selection == "Decimal" {
		var num1, num2 int
		for true {
			fmt.Println("Enter 2 numbers following the corresponding constraints in the format:\na b")
			fmt.Scan(&num1, &num2)
			if (oper == "s" || oper == "d" || oper == "mod") && num1 < num2 {
				continue
			}
			if num1 == 0 || num2 == 0 {
				continue
			}
			if num1%1 == 0 && num1 > 0 && num2%1 == 0 && num2 > 0 {
				break
			}
		}
		num11 = num1
		num22 = num2
	}
	if selection == "Binary" {
		var num1, num2 int;
		var b,d bool;
		var a,c []bool;
		for true {
			fmt.Println("Enter your 2 numbers, there should be no leading 0's and each number should be seperated by a space, but digits should not")
			fmt.Scan(&num1, &num2)
			a, b = validateBase2(num1);
			c, d = validateBase2(num2);
			if b && d{
				break;
			}
		}
		num11 = blToInt(a);
		num22 = blToInt(c);
	}

	return oper, num11, num22
}
