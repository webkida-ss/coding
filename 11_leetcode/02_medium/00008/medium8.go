package main

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	sign := 1
	start := 0
	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	result := 0
	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		digit := int(s[i] - '0')
		if result > (math.MaxInt32-digit)/10 {
			if sign == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
		result = result*10 + digit
	}

	return sign * result
}

func main() {
	fmt.Println(myAtoi("42"))              // 42
	fmt.Println(myAtoi("   -42"))          // -42
	fmt.Println(myAtoi("4193 with words")) // 4193
	fmt.Println(myAtoi("words and 987"))   // 0
	fmt.Println(myAtoi("-91283472332"))    // -2147483648
}
