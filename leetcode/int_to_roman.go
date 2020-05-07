package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/integer-to-roman/

func main() {
	result := intToRoman(2994)
	fmt.Println(result)
}

func intToRoman(num int) string {
	if num < 1 || num > 3999 {
		return ""
	}
	result := ""
	for num > 0 {
		b := 1
		a := num
		for a >= 10 {
			a = a / 10
			b = b * 10
		}
		result += transform(a*b, a)
		num = num - a*b
		fmt.Println(result)
	}
	return result
}

func transform(num, a int) string {
	result := ""
	switch {
	case num >= 1000 && num <= 3000:
		for i := 1; i <= a; i++ {
			result += "M"
		}
	case num >= 100 && num <= 900:
		flag := true
		for i := 1; i <= a; i++ {
			if a == 4 {
				return "CD"
			} else if a == 5 {
				return "D"
			} else if a == 9 {
				return "CM"
			}
			if a > 5 && flag {
				result += "D"
				i = 5
				flag = false
				continue
			}
			result += "C"
		}
	case num == 400:
		result = "CD"
	case num == 500:
		result = "D"
	case num == 900:
		result = "CM"
	case num >= 10 && num <= 90:
		flag := true
		for i := 1; i <= a; i++ {
			if a == 4 {
				return "XL"
			} else if a == 5 {
				return "L"
			} else if a == 9 {
				return "XC"
			}
			if a > 5 && flag {
				result += "L"
				i = 5
				flag = false
				continue
			}
			result += "X"
		}
	case num == 40:
		result = "XL"
	case num == 50:
		result = "L"
	case num == 90:
		result = "XC"

	case num >= 1 && num <= 9:
		flag := true
		for i := 1; i <= a; i++ {
			if a == 4 {
				return "IV"
			} else if a == 5 {
				return "V"
			} else if a == 9 {
				return "IX"
			}
			if a > 5 && flag {
				result += "V"
				i = 5
				flag = false
				continue
			}
			result += "I"
		}
	}
	return result
}
