package code

import "strings"

var thousand = []string{"", "Thousand", "Million", "Billion"} //
var ten = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
var digit = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
	"Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen", "Twenty"}

func numberToWords(num int) (ans string) {
	if num == 0 {
		return "Zero"
	}
	var i int
	for num > 0 {
		res := num % 1000
		if res != 0 {
			ans = calc(res) + thousand[i] + " " + ans
		}
		num = num / 1000
		i++
	}
	return strings.TrimRight(ans, " ")
}

// 处理小于1000的数
func calc(num int) string {
	if num == 0 {
		return ""
	}
	if num <= 20 {
		return digit[num] + " "
	}
	if num < 100 {
		return ten[num/10] + " " + calc(num%10)
	}
	return digit[num/100] + " Hundred " + calc(num%100)
}
