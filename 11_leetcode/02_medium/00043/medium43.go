package main

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	res := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		n1 := int(num1[i] - '0')
		for j := len(num2) - 1; j >= 0; j-- {
			n2 := int(num2[j] - '0')

			sum := n1*n2 + res[i+j+1]

			res[i+j+1] = sum % 10 // 右の桁（余り）
			res[i+j] += sum / 10  // 該当の桁（1or0）

			fmt.Println("--------- 数値 ---------")
			fmt.Println("n1:", n1)
			fmt.Println("n2:", n2)
			fmt.Println("sum:", sum)
			fmt.Println("res[i+j+1]:", res[i+j+1])
			fmt.Println("i+j+1:", i+j+1)
			fmt.Println("res[i+j]:", res[i+j])
			fmt.Println("i+j:", i+j)
			fmt.Println()

		}
	}
	fmt.Println("res:", res)

	result := ""
	for i := 0; i < len(res); i++ {
		if i == 0 && res[i] == 0 {
			continue
		}
		result += strconv.Itoa(res[i])
	}

	return result
}

func main() {
	num1 := "123"
	num2 := "456"
	result := multiply(num1, num2)
	fmt.Println(result) // Output: 56088
}
